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
