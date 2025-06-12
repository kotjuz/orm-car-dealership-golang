package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/car/add")
	server.DELETE("/car/sell", sellCar)
}
