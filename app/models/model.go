package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	*gorm.DB
}

var database *Database

func (d *Database) Connection() {
	db, err := gorm.Open(sqlite.Open("./database/database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.DB = db
}

func (d *Database) AutoMigration() {
	d.DB.AutoMigrate(&User{})
}
