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

func (u User) CheckPassword(password string) bool {
	return u.Password == utils.Hash256(password)
}

func (user *User) SetPassword() {
	log.Println("Setting Password")
	if user.NewPassword != "" && user.NewPassword != user.Password {
		user.Password = utils.Hash256(user.NewPassword)
		user.NewPassword = ""
	}
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	user.SetPassword()
	return
}
