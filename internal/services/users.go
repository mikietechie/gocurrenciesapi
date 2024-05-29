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
	"fmt"

	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

func UpdateUserPassword(user *models.User, body structs.UpdateUserPasswordPayload) error {
	fmt.Println("body.CurrentPassword\t:", body.CurrentPassword)
	fmt.Println("body.NewPassword\t:", body.NewPassword)
	fmt.Println("body.NewPasswordConfirmation\t:", body.NewPasswordConfirmation)
	fmt.Println("body.CheckPassword\t:", user.CheckPassword(body.CurrentPassword))
	if !user.CheckPassword(body.CurrentPassword) {
		return errors.New("passwords don't match")
	}
	// log.Println("user.CheckPassword(body.CurrentPassword)")
	if body.NewPassword != body.NewPasswordConfirmation {
		return errors.New("passwords confirmation failed")
	}
	user.NewPassword = body.NewPassword
	err := models.Db.Save(&user).Error
	if err != nil {
		return errors.New("failed to save to database")
	}
	return nil
}

func UpdateUser(user *models.User, body *models.User) error {
	body.NewPassword = user.Password
	err := models.Db.Model(&user).Updates(body).Error
	if err != nil {
		return err
	}
	return nil
}

func DeactivateUser(user *models.User) error {
	user.Active = false
	return models.Db.Save(&user).Error
}

func RegisterUser(body *models.User) (models.User, error) {
	body.Role = "client"
	err := models.Db.Create(&body).Error
	return *body, err
}
