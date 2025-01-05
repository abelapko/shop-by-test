package main

import (
	"currency-service/config"
	"currency-service/database"

)

func main() {
	cfg := config.LoadConfig()
	cfg.Check()

	database.Connect()
}
