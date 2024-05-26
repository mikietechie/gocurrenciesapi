package main

import (
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func main() {
	var user *models.User
	err := models.Db.First(&user, "role = ?", "admin").Error
	if err != nil {
		log.Fatalln(err)
	}
	if user.ID != 0 {
		log.Println("Admin user already exists.")
		return
	}
	err = models.Db.Model(&user).Create(models.User{
		Role:        "admin",
		Email:       config.GetEnvOrDef("ADMIN_EMAIL", "su@mail.com"),
		NewPassword: config.GetEnvOrDef("ADMIN_PASSWORD", "password"),
		Active:      true,
	}).Error
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully created new admin user.")
}
