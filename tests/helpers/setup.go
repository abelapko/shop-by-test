package helpers

import (
	"currency-service/config"
	"currency-service/database"
	"currency-service/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestEnvironment() (*gin.Engine, *gorm.DB, error) {
	cfg := config.LoadConfig()
	cfg.Check()

	db, err := gorm.Open(mysql.Open(cfg.TestDatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	router := gin.Default()
	database.DB = db
	routes.SetupRoutes(router)

	return router, db, nil
}
