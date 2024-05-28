package drivers

import (
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongod *mongo.Database

func ConnectMongo() {
	log.Println("Process: Connecting to MongoDb")
	client, err := mongo.Connect(config.CTX, options.Client().ApplyURI(config.MONGO_CONNECTION))
	if err != nil {
		log.Panicln(err)
	}

	err = client.Ping(config.CTX, nil)
	if err != nil {
		log.Panicln(err)
	}
	Mongod = client.Database(config.MONGO_DBNAME)
	log.Println("Success: Connected to MongoDb")
}

func DisonnectMongo() {
	err := Mongod.Client().Disconnect(config.CTX)
	if err != nil {
		log.Println("Failed to get DB to close it, most probably it is already closed")
	}
}
