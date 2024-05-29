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
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
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

func ReplenishClientsReads() error {
	err := models.Db.Raw(`
		UPDATE
			clients
			SET
				reads_available=reads_available+?
			WHERE
				deleted_at IS NULL;
	`, config.PERIODIC_READS).Error
	if err != nil {
		log.Println("Error: ", err.Error())
	} else {
		log.Println("Success: ReplenishClientsReads")
	}
	return err
}

func CreateClientForUser(body *models.Client, user models.User) error {
	body.UserID = int(user.ID)
	body.ReadsAvailable = int(config.INITITIAL_READS)
	err := models.Db.Create(&body).Error
	return err
}
