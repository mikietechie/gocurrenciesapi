package utils

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
)

func Hash256(str string) string {
	harsher := sha256.New()
	harsher.Write([]byte(str))
	hashed := harsher.Sum([]byte(config.SECRET_KEY))
	return base64.URLEncoding.EncodeToString(hashed)
}
