package main

import (
	"example.com/ormproject/database"
	"example.com/ormproject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.SeedData()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
