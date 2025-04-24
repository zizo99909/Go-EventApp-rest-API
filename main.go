package main

import (
	"os"

	"example.com/rest-api/caching"
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	caching.ConnectRedis()
	db.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":" + port)

}
