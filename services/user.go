package services

import (
	"simple-api/models"
)

func CreateUser(name string, email string) models.User {
	// Di sini bisa ditambahkan logika untuk menyimpan user ke database
	return models.User{
		ID:       1, // ini seharusnya didapat dari database
		Name:     name,
		Email:    email,
		Password: "password",
	}
}
