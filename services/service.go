package services

import (
	"currency-service/database"
	"currency-service/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func parseDate(dateStr string) (time.Time, error) {
	layoutWithoutTimezone := "2006-01-02T15:04:05"
	layoutWithTimezone := "2006-01-02T15:04:05Z07:00"

	parsedDate, err := time.Parse(layoutWithTimezone, dateStr)
	if err == nil {
		return parsedDate, nil
	}

	parsedDate, err = time.Parse(layoutWithoutTimezone, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date '%s': %v", dateStr, err)
	}

	return parsedDate, nil
}

func FetchAndStoreRates(url string) error {
	if url == "" {
		url = "https://api.nbrb.by/exrates/rates?periodicity=0"
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching rates: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	var apiRates []models.APIRate
	if err := json.Unmarshal(body, &apiRates); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, apiRate := range apiRates {
		parsedDate, err := parseDate(apiRate.Date)
		if err != nil {
			return fmt.Errorf("error parsing date for %s: %v", apiRate.Code, err)
		}

		rate := models.Rate{
			Code:    apiRate.Code,
			Name:    apiRate.Name,
			Rate:    apiRate.Rate,
			Date:    parsedDate,
			Nominal: apiRate.Nominal,
		}

		if err := database.DB.Create(&rate).Error; err != nil {
			return fmt.Errorf("error saving rate for %s: %v", rate.Code, err)
		}

		fmt.Printf("Saved rate for %s: %s, %f\n", rate.Code, rate.Name, rate.Rate)
	}

	fmt.Println("Exchange rates fetched and stored successfully")
	return nil
}
