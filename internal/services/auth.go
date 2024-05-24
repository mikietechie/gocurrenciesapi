package services

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

func GetUserFromCreds(creds structs.LoginPayload) (models.User, error) {
	var user *models.User
	models.Db.Model(&user).First(&user, "Email = ?", creds.Email)
	if user.Active && user.CheckPassword(creds.Password) {
		return *user, nil
	}
	return *user, errors.New("no active user found with the given credentials")
}

func BlackListToken(token jwt.Token) {
	obj := models.BlackListedToken{Token: token.Raw}
	models.Db.Model(&obj).FirstOrCreate(&obj)
	log.Println("Blacklisted Token")
}
