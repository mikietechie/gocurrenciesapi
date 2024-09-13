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

type BeaconRates = map[string]float64

type BeaconResponse struct {
	Meta struct {
		Code       int    `json:"code"`
		Disclaimer string `json:"disclaimer"`
	} `json:"meta"`
	Data struct {
		Date  time.Time   `json:"date"`
		Base  string      `json:"base"`
		Rates BeaconRates `json:"rates"`
	} `json:"response"`
}
