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
	router.PUT("/farms/:id", controllers.FarmUpdate)
	router.DELETE("/farms/:id", controllers.FarmDelete)

	// POND
	router.POST("/createpond", controllers.PondCreate)
	router.GET("/ponds", controllers.PondsIndex)
	router.GET("/ponds/:id", controllers.PondsIndexID)
	router.PUT("/ponds/:id", controllers.PondUpdate)
	router.DELETE("/ponds/:id", controllers.PondDelete)

	// RUN RUN RUN RUN RUN
	router.Run()

}
