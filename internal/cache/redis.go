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
