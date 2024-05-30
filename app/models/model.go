package models

import "simple-api/database"

var DB *database.Database

func Setup() {
	// Initialize the database
	DB = database.NewDatabase()

	DB.AutoMigrate(
		&User{},
		&Car{},
	)
}
