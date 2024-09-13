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

type UpdateClientReadsBody struct {
	Client int `json:"client" binding:"required"`
	Reads  int `json:"reads" binding:"required"`
}
