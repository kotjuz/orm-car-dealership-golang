package routes

import (
	"net/http"

	"example.com/ormproject/database"
	"example.com/ormproject/models"
	"github.com/gin-gonic/gin"
)

func sellCar(context *gin.Context) {
	var vin string
	err := context.ShouldBindJSON(&vin)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	var car models.Samochod

	result := database.DB.Model(&models.Samochod{}).First(&car, "VIN = ?", vin)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There is no such car in database"})
		return
	}

	database.DB.sp

}
