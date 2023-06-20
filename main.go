package main

import (
	"test-api/controllers"
	"test-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectToDB()
}

func main() {
	// Router
	router := gin.Default()
	// POST || Temporary
	router.POST("/createfarm", controllers.FarmCreate)
	router.POST("/createpond", controllers.PondCreate)

	router.Run()

}
