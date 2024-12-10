package main

import (
	"go-translation-api/db"
	"go-translation-api/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("dummy, where is that env file?")
	}

	db.InitDB()
	defer db.Conn.Close()

	router := gin.Default()
	router.GET("/", handlers.GetHome)
	router.GET("/flags", handlers.GetFlags)
	router.GET("/flags/search", handlers.SearchFlags)
	router.GET("/heartbeat", handlers.GetHeartBeat)
	router.GET("/stats", handlers.GetRouteStats)

	serverAddress := "http://localhost:65000"
	log.Printf("Server is running on %s\n", serverAddress)
	log.Fatal(router.Run(":65000"))
}
