// handlers/home.go

package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"os"
)

type RouteInfo struct {
	Route       string `json:"route"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type HomeResponse struct {
	Version      string      `json:"version"`
	Description  string      `json:"description"`
	Environment  string      `json:"environment"`
	Timestamp    string      `json:"timestamp"`
	Documentation string     `json:"documentation"`
	RateLimit    RateInfo   `json:"rate_limit"`
	Routes       []RouteInfo `json:"routes"`
}

type RateInfo struct {
	Limit     int    `json:"limit"`
	Window    string `json:"window"`
	Type      string `json:"type"`
}

func GetHome(c *gin.Context) {
	// api version header
	c.Header("X-API-Version", "0.1.4")

	response := HomeResponse{
		Version:     "0.1.4",
		Description: "Welcome to the Go Translation API. Below are the available endpoints.",
		Environment: getEnvironment(),
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Documentation: "https://github.com/wadedesign/flagapi#readme",
		RateLimit: RateInfo{
			Limit:  100,
			Window: "1 minute",
			Type:   "sliding window",
		},
		Routes: []RouteInfo{
			{Route: "/flags", Method: "GET", Description: "Get flag to language mappings"},
			{Route: "/flags/search", Method: "GET", Description: "Search flags and languages (use ?q=query)"},
			{Route: "/heartbeat", Method: "GET", Description: "Get the heartbeat status"},
			{Route: "/stats", Method: "GET", Description: "Get route call statistics"},
		},
	}

	c.JSON(http.StatusOK, response)
}

func getEnvironment() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		return "development"
	}
	return env
}
