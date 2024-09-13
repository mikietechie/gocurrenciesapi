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
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
)

func AssertIsAdmin(c *gin.Context) error {
	if GetUserFromC(c).Role != "admin" {
		return errors.New("only admins can access this section")
	}
	return nil
}

func WithAssertion(fn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := fn(c)
		if err != nil {
			responses.JSON403(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
