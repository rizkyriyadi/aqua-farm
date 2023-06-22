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

	// FARM
	router.POST("/createfarm", controllers.FarmCreate)
	router.GET("/farms", controllers.Aqua)
	router.GET("/farms/:id", controllers.FarmIndexID)

	// POND
	router.POST("/createpond", controllers.PondCreate)
	router.GET("/ponds", controllers.PondsIndex)
	router.GET("/ponds/:id", controllers.PondsIndexID)
	// RUN RUN RUN RUN RUN
	router.Run()

}
