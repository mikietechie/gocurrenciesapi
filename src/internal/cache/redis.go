/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package cache

import (
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.REDIS_CONNECTION,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := RDB.Ping(config.CTX).Err(); err != nil {
		log.Fatalln("Failed to connect to redis\n", err)
	}
}

func DisonnectRedis() {
	err := RDB.Close()
	if err != nil {
		log.Println("Failed to close Redis, most probably it is already closed")
	} else {
		log.Println("Closed Redis")
	}
}
