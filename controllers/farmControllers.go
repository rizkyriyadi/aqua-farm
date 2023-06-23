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

// GET
// GET ALL
func GetAll(db *gorm.DB) ([]models.Farm, error) {
	var farms []models.Farm
	err := db.Model(&models.Farm{}).Preload("Pond").Find(&farms).Error
	return farms, err
}

// Get Farm with Ponds
func Aqua(c *gin.Context) {
	farms, err := GetAll(initializers.DB)
	if err != nil {
		if len(farms) == 0 {
			c.JSON(404, gin.H{
				"Error": "There is no data in Table",
			})
		}
		// Tangani kesalahan
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Farms": farms})
}

// GET All by ID
func GetID(db *gorm.DB, id string) ([]models.Farm, error) {
	var farms []models.Farm
	result := db.Model(&models.Farm{}).Preload("Pond").Where("id = ?", id).Find(&farms).Error

	return farms, result
}

// Get Farms by ID
func FarmIndexID(c *gin.Context) {
	id := c.Param("id")
	farms, err := GetID(initializers.DB, id)

	// Validasi
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

	}
	// Return
	c.JSON(http.StatusOK, gin.H{"Farms": farms})
}

// PUT

// Update Farm

func FarmUpdate(c *gin.Context) {
	// Get id by URL
	id := c.Param("id")

	// Get data from body
	var body struct {
		FarmName string `json:"farm_name"`
	}

	c.Bind(&body)

	// Find data were updating
	var farm models.Farm
	initializers.DB.First(&farm, id)

	// Update
	initializers.DB.Model(&farm).Updates(models.Farm{
		FarmName: body.FarmName})
	// Respon
	c.JSON(200, gin.H{
		"Updated": farm,
	})
}

// DELETE
func FarmDelete(c *gin.Context) {
	// Get id by URL
	id := c.Param("id")

	// Find data were Delete
	var farm models.Farm
	initializers.DB.Where("id = ?", id).Delete(&farm)

	c.JSON(200, gin.H{
		"Deleted": farm,
	})
}
