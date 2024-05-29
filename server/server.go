package server

import (
	"simple-api/routes"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	InstanceRun() string
	SetupRoutes(routes *[]routes.Router)
	setupRoute()
	Run(port string)
}

type Server struct {
	Engine *gin.Engine
	Routes *[]routes.Router
}

func NewServer() *Server {
	return &Server{
		Engine: gin.Default(),
		Routes: &[]routes.Router{},
	}
}

func (server *Server) InstanceRun() {

	// Instalasi Router Baru
	server.Engine = gin.Default()

	server.SetupRoutes(server.Routes)

	server.Run()

}

func (server *Server) Setup() {
	server.Engine = gin.Default()
}

func (server *Server) SetupRoutes(routes *[]routes.Router) {
	server.Routes = routes
	for _, route := range *routes {
		server.setupRoute(route)
	}

}

func (server *Server) setupRoute(route routes.Router) {
	route.Setup(server.Engine)
}

func (server *Server) Run(port ...string) {
	p := ":8080"
	if len(port) > 0 {
		p = port[0]
	}
	server.Engine.Run(p)
}
