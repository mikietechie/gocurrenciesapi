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
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func AddClientReadsAvailable(client *models.Client, reads int) error {
	client.ReadsAvailable += reads
	err := models.Db.Save(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func AddClientReadsUsed(client *models.Client, reads int) error {
	client.ReadsUsed += reads
	err := models.Db.Save(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func ChangeClientAPIKey(client *models.Client, reads int) error {
	client.SetAPIKey()
	err := models.Db.Save(&client).Error
	if err != nil {
		return err
	}
	return nil
}
