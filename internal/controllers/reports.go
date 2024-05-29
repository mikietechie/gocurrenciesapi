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
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
)

// List Client   godoc
// @Summary      List Client
// @Description  List New Client
// @Tags         Client
// @Produce      json
// @Success      200   {array}  models.Client
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/clients [get]
// @Security     Bearer
func GetUsersClientsReport(c *gin.Context) {
	var items []models.Client
	err := models.Db.Model(models.Client{}).Preload("User").Find(&items).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, items)
}

func ReportsRouter(r gin.RouterGroup) {
	r.Use(
		middleware.WithAuth(),
		middleware.WithUser(),
		middleware.WithAssertion(middleware.AssertIsAdmin),
	)
	r.GET("/users-clients", GetUsersClientsReport)
}
