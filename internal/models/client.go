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
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name           string `gorm:"size:255;not null" json:"name"`
	APIKey         string `gorm:"size:255;not null;unique" json:"api_key"`
	UserID         int    `gorm:"not null;unique" json:"user_id"`
	User           User   `json:"-"`
	ReadsUsed      int    `gorm:"default:0;not null" json:"reads_used"`
	ReadsAvailable int    `gorm:"default:0;not null" json:"reads_available"`
	Domains        string `gorm:"size:1024;default:'*'" json:"domains"`
}

func (client *Client) SetAPIKey() {
	client.APIKey = uuid.NewString()
}

func (client Client) GetReadsLeft() int {
	return client.ReadsAvailable - client.ReadsUsed
}

func (client Client) HasReads() bool {
	return client.GetReadsLeft() > 0
}

func (client *Client) BeforeCreate(tx *gorm.DB) (err error) {
	client.SetAPIKey()
	return
}
