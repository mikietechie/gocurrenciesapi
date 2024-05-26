package config

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

var _ = godotenv.Overload()

var CTX = context.Background()

var SYS_NAME = GetEnvOrDef("SYS_NAME", "Go Currencies API")
var SECRET_KEY = GetEnvOrDef("SECRET_KEY", "$UPER_$EXRE8!")

var ENV = GetEnvOrDef("ENV", "PROD")
var DEV = ENV == "DEV"

var SERVER_ADDRESS = GetEnvOrDef("SERVER_ADDRESS", "0.0.0.0:8000")

// postgres://pg:pass@localhost:5432/crud
var dns = fmt.Sprintf(
	"postgres://%s:%s@%s:%s/%s",
	GetEnvOrDef("POSTGRES_USER", "postgres"),
	GetEnvOrDef("POSTGRES_PASSWORD", "localhost"),
	GetEnvOrDef("DB_HOST", "localhost"),
	GetEnvOrDef("DB_PORT", "5432"),
	GetEnvOrDef("POSTGRES_DB", "gocurrenciesapidb"),
)
var DATABASE_CONNECTION = GetEnvOrDef(
	"DATABASE_CONNECTION",
	dns,
)
var REDIS_CONNECTION = GetEnvOrDef("REDIS_CONNECTION", "localhost:6379")

var BEACON_KEY = GetEnvOrDef("BEACON_KEY", "")
var BEACON_URL = GetEnvOrDef("BEACON_URL", "https://api.currencybeacon.com/v1")
var BEACON_BASE_CURRENCY = GetEnvOrDef("BEACON_BASE_CURRENCY", "USD")

var JWT_TOKEN_LIFETIME = time.Hour * 24
