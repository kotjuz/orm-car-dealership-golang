package main

import (
	"example.com/ormproject/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.SeedData()

	server := gin.Default()

	server.Run(":8080")
}
