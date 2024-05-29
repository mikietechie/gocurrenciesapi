/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package models

import (
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"size:255;not null;unique" json:"email"`
	Password    string `gorm:"size:255;not null" json:"-"`
	NewPassword string `gorm:"size:255" json:"password"`
	Role        string `gorm:"size:255;not null" json:"role"`
	Active      bool   `gorm:"default:true" json:"active"`
}

func (user User) CheckPassword(password string) bool {
	log.Println("UPssword\t:", user.Password)
	log.Println("Password\t:", utils.Hash256(password))
	return user.Password == utils.Hash256(password)
}

func (user *User) SetPassword() {
	log.Println("Step: Setting Password")
	if user.NewPassword != "" && user.NewPassword != user.Password {
		user.Password = utils.Hash256(user.NewPassword)
		user.NewPassword = ""
	}
}

func (user *User) SetRole() {
	if user.Role != "admin" {
		user.Role = "client"
	}
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	user.SetRole()
	user.SetPassword()
	return
}
