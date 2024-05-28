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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
)

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		token := strings.Replace(bearerToken, "Bearer ", "", 1)
		verifiedToken, err := utils.VerifyToken(token)
		if err == nil {
			err = services.CheckBlackToken(*verifiedToken)
		}
		if err != nil {
			log.Println("WithAuth ", err)
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Set("token", *verifiedToken)
		c.Next()
	}
}
