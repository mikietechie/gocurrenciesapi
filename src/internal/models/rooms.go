/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Rooms system database models
*/

package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Uid  string `gorm:"size:255;not null;unique" json:"uid"`
	Name string `gorm:"size:255" json:"name"`
}

type RoomModel struct {
	gorm.Model
	Uid  string `gorm:"size:255;not null;unique" json:"uid"`
	Name string `gorm:"size:255" json:"name"`
}

type RoomUserModel struct {
	gorm.Model
	UserID int `gorm:"not null" json:"userID"`
	// User    UserModel `json:"-"`
	RoomID int `gorm:"not null" json:"roomID"`
	// Room    RoomModel `json:"-"`
	Expires time.Time `gorm:"not null" json:"expires"`
}
