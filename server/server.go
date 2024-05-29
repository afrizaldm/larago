package server

import (
	"simple-api/routes"
)

type IServer interface {
	InstanceRun()
}

type Server struct {
	router *routes.Router
}

func (server *Server) InstanceRun() {

	// Buat instance dari WebRouter
	server.router = &routes.Router{}

	// r := routes.Router{}

	// Setup router
	server.router.Setup()

	// Jalankan server
	server.router.Engine.Run()
}
