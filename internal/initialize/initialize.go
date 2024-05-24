package initialize

import (
	"github.com/mikietechie/gocurrenciesapi/internal/cache"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func Init() {
	cache.ConnectRedis()
	models.ConnectDb()
}

func Tear() {
	cache.DisonnectRedis()
	models.DisonnectDb()
}
