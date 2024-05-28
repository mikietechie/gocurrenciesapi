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
	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func GetUserFromC(c *gin.Context) *models.User {
	user := c.Value("user").(*models.User)
	return user
}
