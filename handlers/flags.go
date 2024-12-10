// handlers/flags.go

package handlers

import (
	"go-translation-api/data"
	"go-translation-api/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetFlags(c *gin.Context) {
	db.IncrementRouteCall("GetFlags")

	c.JSON(http.StatusOK, data.FlagToLanguage)
}

func SearchFlags(c *gin.Context) {
	db.IncrementRouteCall("SearchFlags")
	
	query := strings.ToLower(c.Query("q"))
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	results := make(map[string]interface{})
	for flag, language := range data.FlagToLanguage {
		switch v := language.(type) {
		case string:
			if strings.Contains(strings.ToLower(v), query) {
				results[flag] = v
			}
		case []string:
			for _, lang := range v {
				if strings.Contains(strings.ToLower(lang), query) {
					results[flag] = v
					break
				}
			}
		}
	}

	c.JSON(http.StatusOK, results)
}
