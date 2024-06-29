package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	d := &Database{}

	d.Connect()

	return d
}

func (d *Database) Connect() *Database {
	db, err := gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	d.DB = db

	return d
}
