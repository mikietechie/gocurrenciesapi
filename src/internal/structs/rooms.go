/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Request Response Models For Rooms
*/

package structs

import "time"

type CheckInBody struct {
	RoomUid string `json:"roomUid" binding:"required"`
	UserUid string `json:"userUid" binding:"required"`
}

type ExtendBody struct {
	UserRoomID uint      `json:"userRoomID" binding:"required"`
	Expires    time.Time `json:"expires" binding:"required"`
}

type UserInRoom struct {
	ID        uint      `json:"id" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	Expires   time.Time `json:"expires" binding:"required"`
	UserID    uint      `json:"userID" binding:"required"`
	UserName  string    `json:"userName" binding:"required"`
	UserUid   string    `json:"userUid" binding:"required"`
}

type RoomUsers struct {
	RoomID  uint         `json:"roomID" binding:"required"`
	RoomUid string       `json:"roomUid" binding:"required"`
	Users   []UserInRoom `json:"users" binding:"required"`
}
