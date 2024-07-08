package app

import (
	"simple-api/app/models"
	"simple-api/boostrap/module/auth/session"
	"simple-api/config"
	"simple-api/routes"
	"simple-api/server"

	"github.com/gin-contrib/sessions"
)

func Run() {

	// Get Config
	appConfig := config.AppInstance()

	// Setup Model
	models.Setup()

	// Create New Server
	server := server.NewServer()

	// Setup Session
	store := session.Create(appConfig.APP_SECRET_KEY)
	server.Engine.Use(sessions.Sessions("auth-session", store))

	// Setup Allowed Proxies
	server.SetTrustedProxies(appConfig.APP_TRUSTED_PROXIES)

	// Setup Routes
	server.SetupRoutes(
		routes.NewWeb(),
		routes.NewApi(),
	)

	// Load Glob
	server.LoadHTMLGlob()

	// Setup Static (public files)
	server.Static(appConfig.APP_PUBLIC, "./"+appConfig.APP_PUBLIC)

	// Run server
	server.Run(appConfig.APP_PORT)

}
