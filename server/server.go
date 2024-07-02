package server

import (
	"flag"
	"log"
	"os"
	"simple-api/routes"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	SetupRoutes(routes *[]routes.Router) *Server
	LoadHTMLGlob(pattern ...string) *Server
	Static(relativePath string, root string) *Server
	Run(port string)
}

type Server struct {
	*routes.Engine
}

func NewServer() *Server {
	return &Server{
		&routes.Engine{Engine: gin.Default()},
	}
}

func (server *Server) SetupRoutes(routes ...routes.Router) *Server {
	for _, route := range routes {
		route.Setup(server.Engine)
	}

	return server
}

func (server *Server) LoadHTMLGlob(pattern ...string) *Server {

	var p string
	// debug := flag.Bool("debug", false, "Mode DEBUG")
	debug := os.Getenv("GIN_MODE")
	println("debug", debug)
	flag.Parse()

	if debug == "debug" {
		log.Println("Debug Mode")
		p = "resources/views/**.html"
		if len(pattern) > 0 {
			p = pattern[0]
		}
	} else {
		log.Println("Production Mode")
		p = "views/**.html"
		if len(pattern) > 1 {
			p = pattern[1]
		}
	}

	server.Engine.LoadHTMLGlob(p)

	return server
}

func (server *Server) Static(relativePath string, root string) *Server {

	// Serve static files from the "public" directory
	server.Engine.Static(relativePath, root)

	return server
}

func (server *Server) Run(port ...string) {
	p := ":8080"
	if len(port) > 0 {
		p = port[0]
	}
	server.Engine.Run(p)
}
