package structs

import "time"

type UsersClientsReportRow struct {
	Name           string    `json:"name"`
	ReadsUsed      int       `json:"reads_used"`
	ReadsAvailable int       `json:"reads_available"`
	ClientID       int       `json:"client_id"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	Email          string    `json:"email"`
	Active         bool      `json:"active"`
}
