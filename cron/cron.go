package cron

import (
	"currency-service/services"
	"fmt"

	"github.com/robfig/cron/v3"
)

func StartCron() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() { services.FetchAndStoreRates("") })
	if err != nil {
		fmt.Println("Error while adding cron job:", err)
	}
	c.Start()
}

func StopCron() {
	c := cron.New()
	c.Stop()
}
