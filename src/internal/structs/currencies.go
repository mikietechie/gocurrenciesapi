/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

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
