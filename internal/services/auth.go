package services

import (
	"errors"

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
