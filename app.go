package main

import (
	"example.com/gin/db"
	"example.com/gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.SetupRoutes(server)

	server.Run(":8080")
}
