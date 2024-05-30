package app

import (
	"simple-api/app/models"
	"simple-api/routes"
	"simple-api/server"
)

func Run() {

	models.Setup()

	server := server.NewServer()

	server.SetupRoutes(
		routes.NewApi(),
		routes.NewWeb(),
	)

	server.LoadHTMLGlob()

	server.Static("public", "./public")

	server.Run()

}
