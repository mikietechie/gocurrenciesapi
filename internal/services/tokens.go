package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/cache"
	"github.com/mikietechie/gocurrenciesapi/internal/config"
)

const BLACK_TOKEN_PREFIX = "BLACK_TOKEN-"

func BlackListToken(token jwt.Token) {
	expiry, err := token.Claims.GetExpirationTime()
	if err != nil {
		log.Println("Failed to black list token, because of expiry time err\n", token)
		log.Println(err)
		return
	}
	expiryTime := time.Now().Unix() - expiry.Time.Unix()
	_, err = cache.RDB.Set(config.CTX, BLACK_TOKEN_PREFIX+token.Raw, true, time.Duration(expiryTime)).Result()
	if err != nil {
		log.Println("Failed to black list token, in cache")
		log.Println(err)
	}
}

func CheckBlackToken(token jwt.Token) error {
	blackTokenKey := BLACK_TOKEN_PREFIX + token.Raw
	result, err := cache.RDB.Get(config.CTX, blackTokenKey).Result()
	if err != nil {
		return nil
	}
	if result != "" {
		err = errors.New("token is black listed")
		log.Println("Error ", err.Error())
		return err
	}
	return nil
}
