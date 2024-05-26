package models

func Migrate() {
	Db.AutoMigrate(&User{}, &Client{})
}
