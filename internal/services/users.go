package services

import (
	"errors"
	"fmt"

	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
)

func UpdateUserPassword(user *models.User, body structs.UpdateUserPasswordPayload) error {
	if !user.CheckPassword(body.CurrentPassword) {
		return errors.New("passwords don't match")
	}
	// log.Println("user.CheckPassword(body.CurrentPassword)")
	if body.NewPassword != body.NewPasswordConfirmation {
		return errors.New("passwords confirmation failed")
	}
	user.Password = utils.Hash256(body.NewPassword)
	err := models.Db.Save(&user).Error
	if err != nil {
		return errors.New("failed to save to database")
	}
	return nil
}

func UpdateUser(user *models.User, body *models.User) error {
	body.Password = user.Password
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

func CreateUser(body *models.User) (models.User, error) {
	fmt.Println("Password is here \t:\t", body.Password)
	body.Password = utils.Hash256(body.Password)
	err := models.Db.Create(&body).Error
	return *body, err
}