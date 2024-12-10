// handlers/home.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteInfo struct {
	Route       string `json:"route"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type HomeResponse struct {
	Version     string      `json:"version"`
	Description string      `json:"description"`
	Routes      []RouteInfo `json:"routes"`
}

func GetHome(c *gin.Context) {
	response := HomeResponse{
		Version:     "0.1.3",
		Description: "Welcome to the Go Translation API. Below are the available endpoints.",
		Routes: []RouteInfo{
			{Route: "/flags", Method: "GET", Description: "Get flag to language mappings"},
			{Route: "/flags/search", Method: "GET", Description: "Search flags and languages (use ?q=query)"},
			{Route: "/heartbeat", Method: "GET", Description: "Get the heartbeat status"},
			{Route: "/stats", Method: "GET", Description: "Get route call statistics"},
		},
	}

	c.JSON(http.StatusOK, response)
}
