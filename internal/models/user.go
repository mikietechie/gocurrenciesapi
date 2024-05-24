package models

import (
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	// password string `gorm:"size:255;not null" json:"-"`
	Role   string `gorm:"size:255;not null" json:"role"`
	Active bool   `gorm:"default:true" json:"active"`
}

type ReadUser struct {
	User
	Password string `gorm:"size:255;not null" json:"-"`
}

func (u User) CheckPassword(password string) bool {
	return u.Password == utils.Hash256(password)
}

func (u *User) GetReadUser() ReadUser {
	var ru ReadUser
	Db.Model(&u).Scan(&ru)
	return ru
}
