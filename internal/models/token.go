package models

import "gorm.io/gorm"

type BlackListedToken struct {
    gorm.Model
    Token string `json:"token"`
}
