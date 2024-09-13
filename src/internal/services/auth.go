/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

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
