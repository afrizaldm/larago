package routes

import (
	"simple-api/app/http/controllers"
	"strings"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Setup(r *Engine)
}

type Engine struct {
	*gin.Engine
}

func (e *Engine) RESOURCES(uri string, ctr controllers.Resources) {

	parts := strings.Split(uri, "/")

	lastpart := parts[len(parts)-1]

	e.GET(uri, ctr.Index)                      // index
	e.GET(uri+"/create", ctr.Create)           // create
	e.POST(uri, ctr.Store)                     // store
	e.GET(uri+"/:"+lastpart, ctr.Show)         // show
	e.GET(uri+"/:"+lastpart+"/edit", ctr.Edit) // edit
	e.PUT(uri+"/:"+lastpart, ctr.Update)       // update
	e.PATCH(uri+"/:"+lastpart, ctr.Update)     // update
	e.DELETE(uri+"/:"+lastpart, ctr.Destroy)   // destroy

}

func (e *Engine) Group(uri string, handlers ...gin.HandlerFunc) *Engine {

	routeGroup := e.Engine.Group(uri, handlers...)

	e.RouterGroup = *routeGroup

	return e
}

func (e *Engine) Use(handlers ...gin.HandlerFunc) *Engine {

	e.Engine.Use(handlers...)

	return e
}
