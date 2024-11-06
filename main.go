package main

import (
	database "example.com/questions/db"
	"example.com/questions/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	database.ConnectDatabase()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
