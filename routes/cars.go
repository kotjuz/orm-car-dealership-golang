package routes

import (
	"net/http"

	"example.com/ormproject/database"
	"example.com/ormproject/models"
	"github.com/gin-gonic/gin"
)

func sellCar(context *gin.Context) {
	vin := context.Param("vin")

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	// 	return
	// }

	var count int64
	result := database.DB.Model(&models.Samochod{}).Where("VIN = ?", vin).Count(&count)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There is no such car in database"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": count})

}
