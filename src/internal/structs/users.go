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

type UpdateUserPasswordPayload struct {
	CurrentPassword         string `json:"current_password" binding:"required"`
	NewPassword             string `json:"new_password" binding:"required"`
	NewPasswordConfirmation string `json:"new_password_confirmation" binding:"required"`
}
