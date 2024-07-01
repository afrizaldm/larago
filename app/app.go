package app

import (
	"simple-api/app/models"
	"simple-api/config"
	"simple-api/routes"
	"simple-api/server"
)

func Run() {

	config := config.NewAppConfig().Load()

	models.Setup()

	server := server.NewServer()

	server.SetupRoutes(
		routes.NewApi(),
		routes.NewWeb(),
	)

	server.LoadHTMLGlob()

	server.Static(config.APP_PUBLIC, "./"+config.APP_PUBLIC)

	server.Run(config.APP_PORT)

}
