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
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index        godoc
// @Summary      Index
// @Description  Gets Index
// @Tags         Auth
// @Produce      json
// @Success      200   {object}  map[string]any
// @Router       /api/v1 [get]
func Index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

func IndexRouter(r gin.RouterGroup) {
	r.GET("", Index)
}
