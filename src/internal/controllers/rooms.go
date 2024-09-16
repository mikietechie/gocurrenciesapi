/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Simple endpoints for demonstrating a users rooms system.
*/

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

// Account			godoc
// @Summary      	Rooms Users
// @Description  	Get rooms and users in them
// @Tags         	rooms
// @Produce      	json
// @Success      	200   {object}  []structs.RoomUsers
// @Router       	/api/v1/rooms/rooms-users [get]
func GetRoomsUsers(c *gin.Context) {
	data, err := services.GetRoomsUsers()
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, data)
}

// Extend			godoc
// @Summary      	Extend
// @Description  	Update users stay in room
// @Tags         	rooms
// @Produce      	json
// @Param        	payload  body   structs.ExtendBody  true  "CheckIn JSON"
// @Success      	200   {object}  models.RoomUserModel
// @Failure      	400   {object}  responses.Error400Response
// @Router       	/api/v1/rooms/extend [patch]
func Extend(c *gin.Context) {
	var body structs.ExtendBody
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	roomUser, err := services.Extend(body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, roomUser)
}

// Check In         godoc
// @Summary      	Check In
// @Description  	Register New Client
// @Tags         	rooms
// @Produce      	json
// @Param        	payload  body   structs.CheckInBody  true  "CheckIn JSON"
// @Success      	200   {object}  models.RoomUserModel
// @Failure      	400   {object}  responses.Error400Response
// @Router       	/api/v1/rooms/check-in [post]
func CheckIn(c *gin.Context) {
	var body structs.CheckInBody
	err := c.BindJSON(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	roomUser, err := services.CheckIn(body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, roomUser)
}

func RoomsRouter(r gin.RouterGroup) {
	r.GET("/rooms-users", GetRoomsUsers)
	r.PATCH("/extend", Extend)
	r.POST("/check-in", CheckIn)
}
