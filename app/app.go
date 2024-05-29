package app

import (
	"simple-api/routes"
	"simple-api/server"
)

func Run() {
	server := server.NewServer()

	server.SetupRoutes(&[]routes.Router{
		routes.NewApi(),
		routes.NewWeb(),
	})

	server.LoadHTMLGlob()

	server.Run()

}
