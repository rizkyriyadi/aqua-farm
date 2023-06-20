package controllers

import (
	"errors"
	"net/http"
	"test-api/initializers"
	"test-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PondCreate(c *gin.Context) {
	// Get Body
	var body struct {
		PondName string `json:"pond_name"`
		FarmID   uint   `json:"farm_id"`
	}
	c.Bind(&body)

	// Create || Cek duplikat
	pond := models.Pond{PondName: body.PondName, FarmID: body.FarmID}
	result := initializers.DB.Where("pond_name = ?", body.PondName).First(&pond)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			initializers.DB.Create(&pond)
			c.JSON(200, gin.H{
				"Pond has been successfully created, Detailed": pond,
			})

		}
	} else {
		c.JSON((http.StatusBadRequest), gin.H{
			"message ": http.StatusText(http.StatusBadRequest),
			"reason":   "dupilacte entry / one pond can only be registered to one farm(" + body.PondName + ")",
		})
	}

}
