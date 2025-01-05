package routes

import (
	"github.com/gin-gonic/gin"
	"currency-service/handler"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/rates", handler.GetRates)
	router.GET("/rates/:date", handler.GetRateByDate)
}
