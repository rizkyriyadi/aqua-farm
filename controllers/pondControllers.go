package controllers

import (
	"errors"
	"net/http"
	"test-api/initializers"
	"test-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// POST
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

// GET
// Get All Ponds
func PondsIndex(c *gin.Context) {
	var ponds []models.Pond

	result := initializers.DB.Find(&ponds)

	if result.Error != nil {
		if len(ponds) == 0 {
			c.JSON(404, gin.H{
				"Error": "There is no data in Ponds table",
			})
		}
		panic("Error While Getting data from Ponds")
	}
	// Return
	c.JSON(200, gin.H{
		"Ponds": ponds,
	})
}
func PondsIndexID(c *gin.Context) {
	var ponds models.Pond
	id := c.Param("id")
	result := initializers.DB.First(&ponds, id)

	if result.Error != nil {
		if id != gorm.ErrRecordNotFound.Error() {
			c.JSON(404, gin.H{
				"Error": "There is no id = " + id + " in Ponds Table",
			})
		}
		panic("Error While Getting data from Ponds")
	}
	// Return
	c.JSON(200, gin.H{
		"Ponds": ponds,
	})
}

// PUT
func PondUpdate(c *gin.Context) {
	// Get id by URL
	id := c.Param("id")

	// Get data from body
	var body struct {
		PondName string `json:"pond_name"`
		FarmID   int    `json:"farm_id"`
	}

	c.Bind(&body)

	// Find data were updating
	var pond models.Pond
	initializers.DB.First(&pond, id)

	// Update
	initializers.DB.Model(&pond).Updates(models.Pond{
		PondName: body.PondName, FarmID: uint(body.FarmID)})
	// Respon
	c.JSON(200, gin.H{
		"Updated": pond,
	})
}

// DELETE
func PondDelete(c *gin.Context) {
	// Get id by URL
	id := c.Param("id")

	// Find data were delete
	var pond models.Pond
	initializers.DB.First(&pond, id)

	// Delete
	initializers.DB.Where("id = ?", id).Delete(&pond)

	// Respon
	c.JSON(200, gin.H{
		"Updated": pond,
	})
}
