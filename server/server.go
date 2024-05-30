package server

import (
	"errors"
	"flag"
	"net/http"
	"os"
	"path/filepath"
	"simple-api/routes"
	"strings"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	SetupRoutes(routes *[]routes.Router) *Server
	LoadHTMLGlob(pattern ...string) *Server
	Static(relativePath string, root string) *Server
	Run(port string)
}

type Server struct {
	Engine *routes.ExtendedEngine
	Routes *[]routes.Router
}

func NewServer() *Server {
	return &Server{
		Engine: &routes.ExtendedEngine{Engine: gin.Default()},
		Routes: &[]routes.Router{},
	}
}

func (server *Server) SetupRoutes(routes ...routes.Router) *Server {
	server.Routes = &routes
	for _, route := range routes {
		route.Setup(server.Engine)
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

func (server *Server) StaticNonPrefix(root string) *Server {

	// Serve static files from the "public" directory
	// server.Engine.Static(relativePath, root)

	// Serve static files from the "public" directory at the root URL
	server.Engine.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// Ensure path does not contain ".." or other invalid sequences
		if strings.Contains(path, "..") {
			c.String(http.StatusBadRequest, "Invalid URL")
			return
		}

		// Serve the file if it exists
		filePath := filepath.Join(root, path)
		if _, err := os.Stat("/path/to/whatever"); errors.Is(err, os.ErrNotExist) {
			c.String(http.StatusNotFound, "404 not found")
			return
		}

		c.File(filePath)
	})

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
