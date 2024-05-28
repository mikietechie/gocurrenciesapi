/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/robfig/cron/v3"
)

func main() {
	initialize.Init()
	defer initialize.Tear()
	log.Println("Cron jobs being Created")
	c := cron.New()
	c.AddFunc(fmt.Sprintf("@every %dm", config.RATES_LIFETIME), func() {
		log.Println("Cron: FetchExchangeRates")
		go services.FetchExchangeRates()
	})
	c.Start()
	time.Sleep(time.Minute)
	log.Println("Cron jobs Created")
	defer log.Println("Exited Cron")
	defer c.Stop()
	select {}
}
