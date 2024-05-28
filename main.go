package main

import (
	"simple-api/routes"
)

func main() {
	// Buat instance dari WebRouter
	r := routes.Router{}

	// Setup router
	r.Setup()

	// Jalankan server
	r.Engine.Run()
}
