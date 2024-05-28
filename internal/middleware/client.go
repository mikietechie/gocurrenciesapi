/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
)

func WithClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Query("apikey")
		var client *models.Client
		log.Println("WithClient Middleware:\tapiKey", apiKey)
		err := models.Db.Model(client).First(&client, "api_key = ?", apiKey).Error
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Set("client", client)
		if !client.HasReads() {
			c.Next()
		}
		go services.AddClientReadsUsed(client, 1)
	}
}
