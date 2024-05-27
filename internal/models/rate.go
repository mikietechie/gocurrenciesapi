package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
A MongoDB Data Model for storing fetched Rates at a given timestamp.
*/
type Rate struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Currency  string             `json:"currency" Usage:"required"`
	Value     float64            `json:"value"`
	Timestamp time.Time          `json:"timestamp"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
