/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package initialize

import (
	"log"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/cache"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func Init() {
	cache.ConnectRedis()
	models.ConnectDb()
	models.ConnectMongo()
}

func Tear() {
	log.Println("Tearing down, will sleep for 30 seconds to allow go routines to finish")
	time.Sleep(time.Second * 30)
	cache.DisonnectRedis()
	models.DisonnectDb()
	models.DisonnectMongo()
}
