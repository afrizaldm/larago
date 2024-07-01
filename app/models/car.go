package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
