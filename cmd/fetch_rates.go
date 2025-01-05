package main

import (
	"currency-service/database"
	"currency-service/services"
	"flag"
	"fmt"
	"log"
)

func main() {
	url := flag.String("url", "https://api.nbrb.by/exrates/rates?periodicity=0", "API URL to fetch exchange rates")
	flag.Parse()

	database.Connect()

	err := fetchAndStoreRates(*url)
	if err != nil {
		log.Fatalf("Error fetching and storing exchange rates: %v", err)
	}
}

func fetchAndStoreRates(apiURL string) error {
	err := services.FetchAndStoreRates(apiURL)
	if err != nil {
		return fmt.Errorf("failed to fetch and store rates: %v", err)
	}
	return nil
}
