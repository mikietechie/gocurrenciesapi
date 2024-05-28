/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON200(c *gin.Context, data interface{}) {
	c.IndentedJSON(http.StatusOK, data)
}

func JSON404(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{})
}

func JSON403(c *gin.Context) {
	c.IndentedJSON(http.StatusForbidden, gin.H{})
}

func JSON400(c *gin.Context, err interface{}) {
	c.IndentedJSON(http.StatusBadRequest, Error400Response{Error: err})
}
