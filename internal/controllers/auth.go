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
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mikietechie/gocurrenciesapi/internal/middleware"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
	"github.com/mikietechie/gocurrenciesapi/internal/utils"
)

func GetAuthUser(c *gin.Context) {
	user := middleware.GetUserFromC(c)
	responses.JSON200(c, user)
}

type UserToken struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func returnAuthUser(c *gin.Context, user models.User) {
	token, err := utils.GerateToken(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	data := UserToken{User: user, Token: token}
	responses.JSON200(c, data)
}

// Logout        godoc
// @Summary      Logout
// @Description  Logout user and black list token
// @Tags         Auth
// @Success      200   {object}  bool
// @Router       /api/v1/auth/logout [get]
// @Security     Bearer
func Logout(c *gin.Context) {
	go services.BlackListToken(c.Value("token").(jwt.Token))
	responses.JSON200(c, true)
}

// Deactivate    godoc
// @Summary      Deactivate
// @Description  Deactivate user
// @Tags         Auth
// @Success      200   {object}  bool
// @Router       /api/v1/auth/deactivate [put]
// @Security     Bearer
func Deactivate(c *gin.Context) {
	go services.DeactivateUser(middleware.GetUserFromC(c))
	responses.JSON200(c, true)
}

// Register      godoc
// @Summary      Register
// @Description  Register and Get User with Token
// @Tags         Auth
// @Produce      json
// @Param        payload  body      models.User  true  "models.User JSON"
// @Success      200   {object}  UserToken
// @Router       /api/v1/auth/register [post]
func Register(c *gin.Context) {
	var body *models.User
	var user models.User
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	user, err = services.RegisterUser(body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	returnAuthUser(c, user)
}

// Login         godoc
// @Summary      Login
// @Description  Gets Credentials and Returns Auth Token
// @Tags         Auth
// @Produce      json
// @Param        payload  body      structs.LoginPayload  true  "structs.LoginPayload JSON"
// @Success      200   {object}  UserToken
// @Router       /api/v1/auth/login [post]
func Login(c *gin.Context) {
	var body *structs.LoginPayload
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	user, err := services.GetUserFromCreds(*body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	returnAuthUser(c, user)
}

// Password      godoc
// @Summary      Password
// @Description  Gets Credentials and Returns Auth Token
// @Tags         Auth
// @Produce      json
// @Param        payload  body      structs.UpdateUserPasswordPayload  true  "structs.UpdateUserPasswordPayload JSON"
// @Success      200   {object}  bool
// @Router       /api/v1/auth/password [patch]
// @Security     Bearer
func UpdateUserPassword(c *gin.Context) {
	var body structs.UpdateUserPasswordPayload
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	var user models.User
	err = models.Db.First(&user, c.Param("id")).Error
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	err = services.UpdateUserPassword(&user, body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, true)
}

func AuthRouter(r gin.RouterGroup) {
	r.POST("/login", Login)
	r.POST("/register", Register)
	r.PUT("/deactivate", middleware.WithAuth(), Logout)
	r.GET("/logout", middleware.WithAuth(), middleware.WithUser(), Logout)
	r.GET("/user", middleware.WithAuth(), middleware.WithUser(), GetAuthUser)
	r.PATCH("/password", middleware.WithAuth(), middleware.WithUser(), UpdateUserPassword)
}
