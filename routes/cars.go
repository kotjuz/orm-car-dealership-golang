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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	if count == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There is no car with that VIN"})
		return
	}

	result = database.DB.Delete(&models.Samochod{}, "VIN = ?", vin)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete car"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Succesfully deleted car"})

}

func addNewCar(context *gin.Context) {
	var newCar models.Samochod
	err := context.ShouldBindJSON(&newCar)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	var count int64

	result := database.DB.Model(&models.Samochod{}).Where("VIN = ?", newCar.VIN).Count(&count)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	if count > 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There is already car with that VIN number"})
		return
	}

	result = database.DB.Model(&models.PosSilnik{}).Where("typ_sil_id = ? AND model_id = ?", newCar.TypSil, newCar.ModelID).Count(&count)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	if count == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "This type of engine does not appear in this model"})
		return
	}

	result = database.DB.Model(&models.Dealer{}).Where("Nazwa = ?", newCar.DealerNazwa).Count(&count)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	if count == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Wrong dealer name"})
		return
	}

	result = database.DB.Create(&newCar)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add new car"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Succesfully added new car"})

}
