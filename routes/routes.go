package routes

import (
	"currency-service/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/rates", handler.GetRates)
	router.GET("/rates/:date", handler.GetRateByDate)
}
