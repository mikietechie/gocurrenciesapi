package models

import (
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var err error
	Db, err = gorm.Open(postgres.Open(config.DATABASE_CONNECTION), &gorm.Config{})
	if err != nil && config.DEV {
		log.Println("Postgres Connection failed, trying SQLite")
		Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}
	if err != nil {
		log.Fatalln("Failed to connect to database\n", err)
	}
	// Db.Logger.LogMode(logger.Silent)
	Db.AutoMigrate(&User{}, &Client{}, &BlackToken{})
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
