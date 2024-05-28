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
	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
)

func WithUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token jwt.Token = c.Value("token").(jwt.Token)
		var user *models.User
		userID, err := token.Claims.GetSubject()
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		err = models.Db.Model(user).First(&user, userID).Error
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
