package main

import (
	"log"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/robfig/cron/v3"
)

func main() {
	initialize.Init()
	defer initialize.Tear()
	log.Println("Cron jobs being Created")
	c := cron.New()
	c.AddFunc("@every 15m", func() {
		log.Println("Cron: FetchExchangeRates")
		services.FetchExchangeRates()
	})
	c.Start()
	time.Sleep(time.Minute)
	log.Println("Cron jobs Created")
	defer log.Println("Cron jobs Created")
	defer c.Stop()
	select {}
}
