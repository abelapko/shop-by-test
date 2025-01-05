package main

import (
	"currency-service/config"
	"currency-service/database"
	"currency-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	cfg.Check()

	database.Connect()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":" + cfg.Port)
}
