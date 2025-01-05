package main

import (
	"currency-service/config"
	"currency-service/cron"
	"currency-service/database"
	"currency-service/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/pressly/goose/v3"
)

func main() {
	cfg := config.LoadConfig()
	cfg.Check()

	database.Connect()

	go cron.StartCron()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":" + cfg.Port)
}
