// handlers/flags.go

package handlers

import (
	"go-translation-api/data"
	"go-translation-api/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

var cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
	Name:        "FlagsService",
	MaxRequests: 5,
	Interval:    0, // no interval 
	Timeout:     0, // no timeout for now
})

func GetFlags(c *gin.Context) {
	db.IncrementRouteCall("GetFlags")

	result, err := cb.Execute(func() (interface{}, error) {
		return data.FlagToLanguage, nil
	})

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func SearchFlags(c *gin.Context) {
	db.IncrementRouteCall("SearchFlags")

	query := strings.ToLower(c.Query("q"))
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	results := make(map[string]interface{})

	_, err := cb.Execute(func() (interface{}, error) {
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
		return results, nil
	})

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
		return
	}

	c.JSON(http.StatusOK, results)
}
