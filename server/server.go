package server

import (
	"flag"
	"simple-api/routes"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	InstanceRun() *Server
	SetupRoutes(routes *[]routes.Router) *Server
	LoadHTMLGlob(pattern ...string) *Server
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

func (server *Server) InstanceRun() *Server {

	// Instalasi Router Baru
	server.Engine = gin.Default()

	server.SetupRoutes(server.Routes)

	server.LoadHTMLGlob()

	server.Run()

	return server

}

func (server *Server) SetupRoutes(routes *[]routes.Router) *Server {
	server.Routes = routes
	for _, route := range *routes {
		server.setupRoute(route)
	}

	return server
}

func (server *Server) LoadHTMLGlob(pattern ...string) *Server {

	var p string
	debug := flag.Bool("debug", false, "Mode DEBUG")
	flag.Parse()

	if *debug {
		p = "resources/views/**.html"
		if len(pattern) > 0 {
			p = pattern[0]
		}
	} else {
		p = "views/**.html"
		if len(pattern) > 1 {
			p = pattern[1]
		}
	}

	server.Engine.LoadHTMLGlob(p)

	return server
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
