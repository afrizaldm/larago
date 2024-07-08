package models

import "simple-api/database"

var DB *database.Database = nil

func Setup() *database.Database {

	if DB != nil {
		return DB
	}

	// Initialize the database
	DB = database.NewDatabase()

	DB.AutoMigrate(
		&User{},
		&Car{},
	)

	return DB
}
