package app

import (
	"simple-api/app/models"
	"simple-api/config"
	"simple-api/routes"
	"simple-api/server"
)

func Run() {

	appConfig := config.NewAppConfig().Load()

	models.Setup()

	server := server.NewServer()

	server.SetTrustedProxies(appConfig.APP_TRUSTED_PROXIES)

	server.SetupRoutes(
		routes.NewWeb(),
		routes.NewApi(),
	)

	server.LoadHTMLGlob()

	server.Static(appConfig.APP_PUBLIC, "./"+appConfig.APP_PUBLIC)

	server.Run(appConfig.APP_PORT)

}
