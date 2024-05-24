package api

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
// @Router       /api/v1/ [get]
func Index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

func IndexRouter(r gin.RouterGroup) {
	r.GET("", Index)
}
