/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package services

import (
	"errors"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

func GetUserByUid(uid string) (*models.UserModel, error) {
	object := models.UserModel{Uid: uid}
	err := models.Db.FirstOrCreate(&object, object).Error
	return &object, err
}

func GetRoomByUid(uid string) (*models.RoomModel, error) {
	object := models.RoomModel{Uid: uid}
	err := models.Db.FirstOrCreate(&object, object).Error
	return &object, err
}

func ExitOtherRooms(userID uint) error {
	query := `
		UPDATE public.roomsusers
		SET
			expires=noww
		WHERE
			userID=userID
	`
	err := models.Db.Raw(query).Error
	return err
}

func CheckIn(body structs.CheckInBody) (models.RoomUserModel, error) {
	var object models.RoomUserModel
	user, err := GetUserByUid(body.UserUid)
	if err != nil {
		return object, err
	}
	room, err := GetRoomByUid(body.RoomUid)
	if err != nil {
		return object, err
	}
	err = ExitOtherRooms(user.ID)
	if err != nil {
		return object, err
	}
	object = models.RoomUserModel{
		UserID:  int(user.ID),
		RoomID:  int(room.ID),
		Expires: time.Now().Add(time.Minute * 5),
	}
	err = models.Db.Create(&object).Error
	return object, err
}

func Extend(body structs.ExtendBody) (models.RoomUserModel, error) {
	var object models.RoomUserModel
	if body.Expires.Before(time.Now()) {
		return object, errors.New("Expires should be grater than or equal to now")
	}
	err := models.Db.First(&object, body.UserRoomID).Error
	if err != nil {
		return object, err
	}
	if object.Expires.Before(time.Now()) {
		return object, errors.New("Already out of room")
	}
	object.Expires = body.Expires
	err = models.Db.Save(&object).Error
	return object, err
}

func GetRoomsUsers() ([]structs.RoomUsers, error) {
	data := []structs.RoomUsers{}
	var err error
	var rooms []models.RoomModel
	err = models.Db.Find(&rooms).Error
	if err != nil {
		return data, err
	}
	for _, room := range rooms {
		var users []structs.UserInRoom
		query := `
			SELECT
				t2.name,
				t2.id AS user_id,
				t2.uid as user_uid,
				t1.id as id,
				t1.created_at,
				t1.expires
			FROM public.room_user_models t1
			JOIN public.user_models t2
			ON
				t1.user_id = t2.id
			Where
				t1.room_id = ?
				AND t1.expires > ?
		`
		err = models.Db.Raw(query, room.ID, time.Now()).Find(&users).Error
		if err != nil {
			return data, err
		}
		item := structs.RoomUsers{RoomID: room.ID, RoomUid: room.Uid, Users: users}
		data = append(data, item)
	}
	return data, err
}
