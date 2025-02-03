package routes

import (
	"api_tracker/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/create-client", controllers.CreateClient)
	router.POST("/login", controllers.GetUser)
	router.POST("/track", controllers.TrackRequest)
}
