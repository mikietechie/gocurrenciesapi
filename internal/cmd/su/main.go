/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package main

import (
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func main() {
	initialize.Init()
	defer initialize.Tear()
	var user models.User
	err := models.Db.First(&user, "role = ?", "admin").Error
	if err == nil {
		log.Println("Admin user already exists.")
		return
	}
	user = models.User{
		Role:        "admin",
		Email:       config.GetEnvOrDef("ADMIN_EMAIL", "su@mail.com"),
		NewPassword: config.GetEnvOrDef("ADMIN_PASSWORD", "password"),
		Active:      true,
	}
	err = models.Db.Model(&user).Create(&user).Error
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully created new admin user.")
}
