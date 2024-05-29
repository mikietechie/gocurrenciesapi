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
	"fmt"

	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

// This is a file for writing Go Code Trials, since go has no REPL
// go run internal/cmd/repl/main.go
// All code added to be removed after execution
func main() {
	initialize.Init()
	defer initialize.Tear()
	var user models.User
	models.Db.First(&user, "email = ?", "su@mail.com")
	fmt.Println(user.Email)
	user.NewPassword = "password"
	err := models.Db.Save(&user).Error
	if err != nil {
		fmt.Println("failed to set password")
	}
}
