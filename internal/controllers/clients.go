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
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
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
func GetClients(c *gin.Context) {
	var items []models.Client
	err := models.Db.Model(models.Client{}).Preload("User").Find(&items).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, items)
}

// Expand Client   godoc
// @Summary      Expand Client
// @Description  Expand Client
// @Tags         Client
// @Produce      json
// @Param        id path int true "Client ID"
// @Success      200   {object}  models.Client
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/clients/{id} [get]
// @Security     Bearer
func GetClient(c *gin.Context) {
	var item models.Client
	err := models.Db.Model(models.Client{}).First(&item, c.Param("id")).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, item)
}

// Delete Client   godoc
// @Summary      Delete Client
// @Description  Delete Client
// @Tags         Client
// @Produce      json
// @Param        id path int true "Client ID"
// @Success      200   {object}  bool
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/clients/{id} [delete]
// @Security     Bearer
func DeleteClient(c *gin.Context) {
	var item models.Client
	err := models.Db.First(&item, c.Param("id")).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	models.Db.Delete(&item)
	responses.JSON200(c, true)
}

// Create Client   godoc
// @Summary      Create Client
// @Description  Create New Client
// @Tags         Client
// @Produce      json
// @Param        payload  body      models.Client  true  "models.Client JSON"
// @Success      200   {object}  models.Client
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/clients [post]
// @Security     Bearer
func CreateClient(c *gin.Context) {
	var item *models.Client
	err := c.BindJSON(&item)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	err = models.Db.Create(&item).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, item)
}

// Update Client   godoc
// @Summary      Update Client
// @Description  Update Client
// @Tags         Client
// @Produce      json
// @Param        payload  body      models.Client  true  "models.Client JSON"
// @Param        id path int true "Client ID"
// @Success      200   {object}  models.Client
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/clients/{id} [put]
// @Security     Bearer
func UpdateClient(c *gin.Context) {
	var body models.Client
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	var item models.Client
	err = models.Db.First(&item, c.Param("id")).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	err = models.Db.Model(&item).Updates(body).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, item)
}

// Client        godoc
// @Summary      Client
// @Description  Gets Credentials and Returns Auth Token
// @Tags         Auth
// @Produce      json
// @Param        payload  body      structs.UpdateClientReadsBody  true  "structs.UpdateClientReadsBody JSON"
// @Success      200   {object}  models.Client
// @Router       /api/v1/auth/password [patch]
// @Security     Bearer
func AddClientReadsAvailable(c *gin.Context) {
	var body structs.UpdateClientReadsBody
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	var client models.Client
	err = models.Db.First(&client, body.Client).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	err = services.AddClientReadsAvailable(&client, body.Reads)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	models.Db.Model(&client).Scan(&client)
	responses.JSON200(c, client)
}

func ClientsRouter(r gin.RouterGroup) {
	r.Use(
		middleware.WithAuth(),
		middleware.WithUser(),
		middleware.WithAssertion(middleware.AssertIsAdmin),
	)
	r.GET("/", GetClients)
	r.GET("/:id", GetClient)
	r.DELETE("/:id", DeleteClient)
	r.PUT("/:id", UpdateClient)
	r.PATCH("/", AddClientReadsAvailable)
	r.POST("/", CreateClient)
}
