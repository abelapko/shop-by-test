package main

import (
	"currency-service/config"

)

func main() {
	cfg := config.LoadConfig()
	cfg.Check()

}
