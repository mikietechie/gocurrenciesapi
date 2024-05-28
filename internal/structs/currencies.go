package structs

import "time"

type RateAtDateBody struct {
	Timestamp time.Time `form:"timestamp" binding:"required"`
	Currency  string    `form:"currency" binding:"required"`
}

type RatesInPeriodBody struct {
	Start      time.Time `form:"start" binding:"required"`
	End        time.Time `form:"end" binding:"required"`
	Currencies []string  `form:"currencies" binding:"required"`
}
