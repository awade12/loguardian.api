// handlers/heartbeat.go

package handlers

import (
	"net/http"
	"time"
	"go-translation-api/db"
	"github.com/gin-gonic/gin"
)

var startTime time.Time

type StatusResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
	Message   string `json:"message"`
}

func init() {
	startTime = time.Now()
}

func GetHeartBeat(c *gin.Context) {
	db.IncrementRouteCall("GetHeartBeat")
	uptime := time.Since(startTime)

	response := StatusResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    uptime.String(),
		Message:   "Service is running smoothly.",
	}

	c.JSON(http.StatusOK, response)
}
