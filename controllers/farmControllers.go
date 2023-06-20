package controllers

import (
	"errors"
	"net/http"
	"test-api/initializers"
	"test-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FarmCreate(c *gin.Context) {
	// Get Body
	var body struct {
		FarmName string `json:"farm_name"`
	}
	c.Bind(&body)

	// Create || Cek duplikat
	farm := models.Farm{FarmName: body.FarmName}
	result := initializers.DB.Where("farm_name = ?", body.FarmName).First(&farm)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			initializers.DB.Create(&farm)
			c.JSON(200, gin.H{
				"Farm has been successfully created, Detailed": farm,
			})

		}
	} else {
		c.JSON(409, gin.H{
			"message ": http.StatusText(http.StatusBadRequest),
			"reason":   "dupilacte entry (" + body.FarmName + ")",
		})
	}
	// Return

}
