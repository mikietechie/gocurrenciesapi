/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/middleware"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
)

// Users Clients Report godoc
// @Summary      		Users Clients Report
// @Description  		Users with Clients Report
// @Tags         		Reports
// @Produce      		json
// @Success      		200   {array}  structs.UsersClientsReportRow
// @Failure      		400   {object}  responses.Error400Response
// @Router       		/api/v1/reports/users-clients [get]
// @Security     		Bearer
func GetUsersClientsReport(c *gin.Context) {
	data, err := services.GetUsersClientsReportData()
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, data)
}

func ReportsRouter(r gin.RouterGroup) {
	r.Use(
		middleware.WithAuth(),
		middleware.WithUser(),
		middleware.WithAssertion(middleware.AssertIsAdmin),
	)
	r.GET("/users-clients", GetUsersClientsReport)
}
