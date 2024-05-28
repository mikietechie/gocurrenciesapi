package initialize

import (
	"log"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/cache"
	"github.com/mikietechie/gocurrenciesapi/internal/drivers"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func Init() {
	cache.ConnectRedis()
	models.ConnectDb()
	drivers.ConnectMongo()
}

func Tear() {
	log.Println("Tearing down, will sleep for 30 seconds to allow go routines to finish")
	time.Sleep(time.Second * 30)
	cache.DisonnectRedis()
	models.DisonnectDb()
	drivers.DisonnectMongo()
}
