package app

import (
	"simple-api/routes"
)

func Run() {
	// Buat instance dari WebRouter
	r := routes.Router{}

	// Setup router
	r.Setup()

	// Jalankan server
	r.Engine.Run()
}
