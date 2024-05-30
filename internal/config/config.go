/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package config

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var _ = godotenv.Overload()
var CTX = context.Background()

/* System */
var SYS_NAME = GetEnvOrDef("SYS_NAME", "Go Currencies API")
var SECRET_KEY = GetEnvOrDef("SECRET_KEY", "$UPER_$EXRE8!")
var ENV = GetEnvOrDef("ENV", "PROD")
var DEV = ENV == "DEV"
var SERVER_ADDRESS = GetEnvOrDef("SERVER_ADDRESS", "0.0.0.0:8000")

/* POSTGRES MONGO REDIS */
var DATABASE_CONNECTION = GetEnvOrDef(
	"POSTGRES_CONNECTION",
	"postgres://pg:pass@localhost:5432/db",
)
var MONGO_CONNECTION = GetEnvOrDef(
	"MONGO_CONNECTION",
	"mongodb://localhost:27017/?authSource=admin",
)
var MONGO_DBNAME = GetEnvOrDef("MONGO_DBNAME", "db")
var REDIS_CONNECTION = GetEnvOrDef("REDIS_CONNECTION", "localhost:6379")

/*BEACON DATA SOURCE*/
var BEACON_KEY = GetEnvOrDef("BEACON_KEY", "")
var BEACON_URL = GetEnvOrDef("BEACON_URL", "https://api.currencybeacon.com/v1")
var BEACON_BASE_CURRENCY = GetEnvOrDef("BEACON_BASE_CURRENCY", "USD")
var RATES_LIFETIME, _ = strconv.ParseInt(
	GetEnvOrDef("RATES_LIFETIME", "60"),
	10,
	64,
) // Int: Rates cache lifetime
var envJWTTokenLifetime, _ = strconv.ParseInt(
	GetEnvOrDef("JWT_TOKEN_LIFETIME", "3600"),
	10,
	64,
)
var JWT_TOKEN_LIFETIME = time.Minute * time.Duration(envJWTTokenLifetime)

// Client Reads Assignment
var INITITIAL_READS, _ = strconv.ParseInt(
	GetEnvOrDef("INITITIAL_READS", "10"),
	10,
	64,
)
var PERIODIC_READS, _ = strconv.ParseInt(
	GetEnvOrDef("PERIODIC_READS", "10"),
	10,
	64,
)
var REPLANISH_PERIOD, _ = strconv.ParseInt(
	GetEnvOrDef("REPLANISH_PERIOD", "30"),
	10,
	64,
)

func init() {
	if RATES_LIFETIME <= 0 {
		log.Fatalln("Environment 'RATES_LIFETIME' should be a positive integer")
	}
	if JWT_TOKEN_LIFETIME <= 0 {
		log.Fatalln("Environment 'JWT_TOKEN_LIFETIME' should be a positive integer")
	}
	if PERIODIC_READS <= 0 {
		log.Fatalln("Environment 'PERIODIC_READS' should be a positive integer")
	}
}
