package structs

import "time"

type BeaconResponse struct {
	Meta struct {
		Code       int    `json:"code"`
		Disclaimer string `json:"disclaimer"`
	} `json:"meta"`
	Data struct {
		Date  time.Time          `json:"date"`
		Base  string             `json:"base"`
		Rates map[string]float64 `json:"rates"`
	} `json:"response"`
}
