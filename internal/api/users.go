package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/middleware"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
)

// List User     godoc
// @Summary      List User
// @Description  List New User
// @Tags         User
// @Produce      json
// @Success      200	{array}  models.User
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/users [get]
// @Security     Bearer
func GetUsers(c *gin.Context) {
	var items []models.User
	err := models.Db.Model(models.User{}).Find(&items).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, items)
}

// Expand User   godoc
// @Summary      Expand User
// @Description  Expand User
// @Tags         User
// @Produce      json
// @Param        id	path	int	true	"User ID"
// @Success      200	{object}	models.User
// @Failure      400	{object}	responses.Error400Response
// @Router       /api/v1/users/{id} [get]
// @Security     Bearer
func GetUser(c *gin.Context) {
	var item models.User
	err := models.Db.Model(models.User{}).First(&item, c.Param("id")).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, item)
}

// Delete User   godoc
// @Summary      Delete User
// @Description  Delete User
// @Tags         User
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200   {object}  bool
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/users/{id} [delete]
// @Security     Bearer
func DeleteUser(c *gin.Context) {
	var item models.User
	err := models.Db.First(&item, c.Param("id")).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	models.Db.Delete(&item)
	responses.JSON200(c, true)
}

// Create User   godoc
// @Summary      Create User
// @Description  Create New User
// @Tags         User
// @Produce      json
// @Param        payload  body      models.User  true  "models.User JSON"
// @Success      200   {object}  models.User
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/users [post]
// @Security     Bearer
func CreateUser(c *gin.Context) {
	var item *models.User
	err := c.BindJSON(&item)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	item.Password = utils.Hash256(item.Password)
	err = models.Db.Create(&item).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, item)
}

// Update User   godoc
// @Summary      Update User
// @Description  Update User
// @Tags         User
// @Produce      json
// @Param        payload  body      models.User  true  "models.User JSON"
// @Param        id path int true "User ID"
// @Success      200   {object}  models.User
// @Failure      400   {object}  responses.Error400Response
// @Router       /api/v1/users/{id} [put]
// @Security     Bearer
func UpdateUser(c *gin.Context) {
	var body models.User
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	var item models.User
	err = models.Db.First(&item, c.Param("id")).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	err = services.UpdateUser(&item, &body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, item)
}

func UsersRouter(r gin.RouterGroup) {
	r.Use(
		middleware.WithAuth(),
		middleware.WithUser(),
		middleware.WithAssertion(middleware.AssertIsAdmin),
	)
	r.GET("/", GetUsers)
	r.GET("/:id", GetUser)
	r.DELETE("/:id", DeleteUser)
	r.PUT("/:id", UpdateUser)
	r.POST("/", CreateUser)
}
