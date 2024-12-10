// handlers/heartbeat.go

package handlers

import (
	"net/http"
	"runtime"
	"time"
	"go-translation-api/db"
	"github.com/gin-gonic/gin"
	"context"
)

var (
	startTime = time.Now()
	version   = "0.1.4"
)

type StatusResponse struct {
	Status    string      `json:"status"`
	Timestamp string      `json:"timestamp"`
	Uptime    string      `json:"uptime"`
	Message   string      `json:"message"`
	Version   string      `json:"version"`
	System    SystemInfo  `json:"system"`
	Database  DbStatus    `json:"database"`
	Metrics   MetricsInfo `json:"metrics"`
}

type SystemInfo struct {
	GoVersion    string `json:"go_version"`
	NumGoroutine int    `json:"num_goroutines"`
	NumCPU       int    `json:"num_cpu"`
}

type DbStatus struct {
	Connected bool   `json:"connected"`
	Message   string `json:"message"`
}

type MetricsInfo struct {
	TotalRequests int64 `json:"total_requests"`
}

func GetHeartBeat(c *gin.Context) {
	db.IncrementRouteCall("GetHeartBeat")
	uptime := time.Since(startTime)

	dbStatus := DbStatus{
		Connected: true,
		Message:   "Connected",
	}
	if err := db.Conn.Ping(context.Background()); err != nil {
		dbStatus.Connected = false
		dbStatus.Message = "Disconnected: " + err.Error()
	}

	totalRequests := db.GetTotalRequests()

	response := StatusResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    uptime.String(),
		Message:   "Service is running smoothly.",
		Version:   version,
		System: SystemInfo{
			GoVersion:    runtime.Version(),
			NumGoroutine: runtime.NumGoroutine(),
			NumCPU:       runtime.NumCPU(),
		},
		Database: dbStatus,
		Metrics: MetricsInfo{
			TotalRequests: totalRequests,
		},
	}

	c.JSON(http.StatusOK, response)
}
