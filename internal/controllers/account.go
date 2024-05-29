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
	"github.com/mikietechie/gocurrenciesapi/internal/services"
)

// Account			godoc
// @Summary      	Account
// @Description  	Gets Account Using Key
// @Tags         	Account
// @Produce      	json
// @Success      	200   {object}  models.Client
// @Router       	/api/v1/account/client/using-key [get]
// @Security 		ApiKeyAuth
func GetClientUsingKey(c *gin.Context) {
	client := middleware.GetClientFromC(c)
	responses.JSON200(c, client)
}

// Account			godoc
// @Summary      	Account
// @Description  	Gets Account Using TOken
// @Tags         	Account
// @Produce      	json
// @Success      	200   {object}  models.Client
// @Router       	/api/v1/account/client/using-token [get]
// @Security 		Bearer
func GetClientUsingToken(c *gin.Context) {
	client, err := middleware.GetUserFromC(c).GetClient()
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, client)
}

// Register Client  godoc
// @Summary      	Register Client
// @Description  	Register New Client
// @Tags         	Account
// @Produce      	json
// @Param        	payload  body      models.Client  true  "models.Client JSON"
// @Success      	200   {object}  models.Client
// @Failure      	400   {object}  responses.Error400Response
// @Router       	/api/v1/account/register [post]
// @Security     	Bearer
func RegisterClient(c *gin.Context) {
	var client *models.Client
	err := c.BindJSON(&client)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	err = services.CreateClientForUser(client, *middleware.GetUserFromC(c))
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, client)
}

func AccountRouter(r gin.RouterGroup) {
	r.GET("/client/using-key", middleware.WithClient(), GetClientUsingKey)
	r.GET("/client/using-token", middleware.WithAuth(), middleware.WithUser(), GetClientUsingToken)
	r.POST("/register", middleware.WithAuth(), middleware.WithUser(), RegisterClient)
}
