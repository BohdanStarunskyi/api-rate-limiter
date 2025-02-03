package main

import (
	"api_tracker/config"
	cronManager "api_tracker/cron_manager"
	"api_tracker/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cm := cronManager.New()
	cm.Start()

	config.LoadConfig()

	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
