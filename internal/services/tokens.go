package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func BlackListToken(token jwt.Token) {
	expiry, err := token.Claims.GetExpirationTime()
	if err != nil {
		log.Println("Failed to black list token, because of expiry time err\n", token)
		log.Println(err)
		return
	}
	obj := models.BlackToken{Token: token.Raw, ExpiresAt: expiry.Time}
	models.Db.Model(&obj).FirstOrCreate(&obj)
	log.Println("Blacklisted Token")
}

func ClearExpiredBlackToken() {
	now := time.Now()
	result := models.Db.Model(&models.BlackToken{}).Where("expires_at > ?", now).Delete(nil)
	if result.Error != nil {
		log.Println("Failed to clear black tokens at ", now)
		log.Println(result.Error)
		return
	}
	log.Printf("Success, cleared %d black tokens at %s \n", result.RowsAffected, time.Now())
}

func CheckBlackToken(token jwt.Token) error {

	var blackToken models.BlackToken
	err := models.Db.Model(&blackToken).First(&blackToken, "token = ?", token.Raw)
	log.Println("Black token \t", blackToken)
	if err != nil {
		return nil
		// log.Println("Failed to find black token\n", token)
		// log.Println(err)
		// return err.Error
	}
	return errors.New("token is black listed")
}
