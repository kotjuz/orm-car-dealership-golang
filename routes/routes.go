package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/car/add", addNewCar)
	server.DELETE("/car/sell/:vin", sellCar)
}
