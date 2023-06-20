package main

import (
	"test-api/initializers"
	"test-api/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	// initializers.DB.AutoMigrate(&models.User{}, &models.CreditCard{})
	initializers.DB.AutoMigrate(&models.Farm{}, &models.Pond{})
}
