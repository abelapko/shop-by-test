package handler

import (
	"currency-service/database"
	"currency-service/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRates(c *gin.Context) {
	var rates []models.Rate
	if err := database.DB.Find(&rates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch rates"})
		return
	}
	c.JSON(http.StatusOK, rates)
}

func GetRateByDate(c *gin.Context) {
	dateStr := c.Param("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	var rates []models.Rate
	formattedDate := date.Format("2006-01-02")

	fmt.Println("Formatted date for query:", formattedDate)

	if err := database.DB.Where("DATE(date) = ?", formattedDate).Find(&rates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch rates"})
		return
	}

	fmt.Println("Rates found:", len(rates))

	c.JSON(http.StatusOK, rates)
}
