package structs

import "time"

type RateAtDateBody struct {
	Timestamp time.Time `json:"timestamp" binding:"required"`
	Currency  string    `json:"currency" binding:"required"`
}

type RatesInPeriodBody struct {
	Start      time.Time `json:"start" binding:"required"`
	End        time.Time `json:"end" binding:"required"`
	Currencies []string  `json:"currencies" binding:"required"`
}
