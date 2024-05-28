package main

import (
	"simple-api/routes"
)

func main() {
	// Buat instance dari WebRouter
	webRouter := &routes.WebRouter{}

	// Setup router
	r := webRouter.SetupRouter()

	// Jalankan server
	r.Engine.Run()
}
