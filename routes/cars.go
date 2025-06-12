package routes

import (
	"net/http"

	"example.com/ormproject/database"
	"example.com/ormproject/models"
	"github.com/gin-gonic/gin"
)

func sellCar(context *gin.Context) {
	vin := context.Param("vin")
	var count int64

	result := database.DB.Model(&models.Samochod{}).Where("VIN = ?", vin).Count(&count)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There is no such car in database"})
		return
	}

	result = database.DB.Delete(&models.Samochod{}, "VIN = ?", vin)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete car"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Succesfully deleted car"})

}
