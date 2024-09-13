/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package services

import (
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

func GetUsersClientsReportData() ([]structs.UsersClientsReportRow, error) {
	var data []structs.UsersClientsReportRow
	query := `
		SELECT
			t2.email,
			t2.id AS user_id,
			t2.active,
			t1.id AS client_id,
			t1.name,
			t1.reads_available,
			t1.reads_used,
			t1.created_at
		FROM public.clients t1
		JOIN public.users t2
		ON
			t1.user_id = t2.id
	`
	err := models.Db.Raw(query).Find(&data).Error
	return data, err
}
