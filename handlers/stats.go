// handlers/stats.go

package handlers

import (
	"context"
	"go-translation-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteStats struct {
	RouteName string `json:"route_name"`
	CallCount int    `json:"call_count"`
}

func GetRouteStats(c *gin.Context) {
	rows, err := db.Conn.Query(context.Background(), "SELECT route_name, call_count FROM route_calls")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch route statistics"})
		return
	}
	defer rows.Close()

	var stats []RouteStats
	for rows.Next() {
		var stat RouteStats
		err := rows.Scan(&stat.RouteName, &stat.CallCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning route statistics"})
			return
		}
		stats = append(stats, stat)
	}

	c.JSON(http.StatusOK, stats)
}

func IncrementRouteCall(routeName string) error {
	_, err := db.Conn.Exec(context.Background(), `
		INSERT INTO route_calls (route_name, call_count) 
		VALUES ($1, 1) 
		ON CONFLICT (route_name) 
		DO UPDATE SET call_count = route_calls.call_count + 1`, routeName)
	return err
}
