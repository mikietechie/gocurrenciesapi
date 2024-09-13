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
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
A MongoDB Data Model for storing fetched Rates at a given timestamp.
*/
type Rate struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Currency  string             `json:"currency" Usage:"required"`
	Value     float64            `json:"value"`
	Timestamp time.Time          `json:"timestamp"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
