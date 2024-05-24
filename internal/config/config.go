package config

import (
	"context"

	"github.com/joho/godotenv"
)

var _ = godotenv.Overload()

var CTX = context.Background()

var SYS_NAME = GetEnvOrDef("SYS_NAME", "Go Currencies API")
var SECRET_KEY = GetEnvOrDef("SECRET_KEY", "$UPER_$EXRE8!")

var SERVER_ADDRESS = GetEnvOrDef("SERVER_ADDRESS", "0.0.0.0:8000")
var REDIS_CONNECTION = GetEnvOrDef("REDIS_CONNECTION", "localhost:6379")

var BEACON_KEY = GetEnvOrDef("BEACON_KEY", "")
var BEACON_URL = GetEnvOrDef("BEACON_URL", "https://api.currencybeacon.com/v1")
var BEACON_BASE_CURRENCY = GetEnvOrDef("BEACON_BASE_CURRENCY", "USD")
