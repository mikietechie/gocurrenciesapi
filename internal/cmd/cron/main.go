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
	log.Println("Status: Cron jobs being Created")
	c := cron.New()

	/*
		After every time period we want to fetch fresh data from our
		Data Source cache it locally and store it in the Database
	*/
	c.AddFunc(fmt.Sprintf("@every %dm", config.RATES_LIFETIME), func() {
		log.Println("Cron: FetchExchangeRates")
		go services.FetchExchangeRates()
	})

	/*
		After a certain time period we give free reads to our clients
	*/
	c.AddFunc(fmt.Sprintf("@every %dd", config.REPLANISH_PERIOD), func() {
		log.Println("Cron: ReplenishClientsReads")
		go services.ReplenishClientsReads()
	})
	c.Start()

	time.Sleep(time.Minute)
	log.Println("Success: Cron jobs Created")
	defer log.Println("Status: Exited Cron")
	defer c.Stop()
	select {}
}
