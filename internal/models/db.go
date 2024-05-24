package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	// var dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	// Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	// Db.Logger.LogMode(logger.Silent)
	Db.AutoMigrate(&User{}, &Client{}, &BlackListedToken{})
}

func DisonnectDb() {
	db, err := Db.DB()
	if err != nil {
		log.Println("Failed to get DB to close it, most probably it is already closed")
	}
	err = db.Close()
	if err != nil {
		log.Println("Failed to get DB to close it, most probably it is already closed")
	} else {
		log.Println("Db Closed")
	}
}
