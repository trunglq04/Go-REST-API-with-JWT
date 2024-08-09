package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	// setup handler for incoming GET request
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
