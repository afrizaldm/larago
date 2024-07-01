package routes

import (
	"simple-api/app/http/controllers"
	"strings"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Setup(r *Engine)
	Group(uri string, handlers ...gin.HandlerFunc) *RouterGroup
}

type Engine struct {
	*gin.Engine
}

type RouterGroup struct {
	*Engine
	*gin.RouterGroup
}

func (e *Engine) RESOURCES(uri string, ctr controllers.Resources) {

	part := getLastPartOfPath(uri)

	e.GET(uri, ctr.Index)                  // index
	e.GET(uri+"/create", ctr.Create)       // create
	e.POST(uri, ctr.Store)                 // store
	e.GET(uri+"/:"+part, ctr.Show)         // show
	e.GET(uri+"/:"+part+"/edit", ctr.Edit) // edit
	e.PUT(uri+"/:"+part, ctr.Update)       // update
	e.PATCH(uri+"/:"+part, ctr.Update)     // update
	e.DELETE(uri+"/:"+part, ctr.Destroy)   // destroy

}

func (e *Engine) Group(uri string, handlers ...gin.HandlerFunc) *RouterGroup {
	routeGroup := e.Engine.Group(uri, handlers...)

	r := &RouterGroup{
		Engine:      e,
		RouterGroup: routeGroup,
	}

	return r
}

func getLastPartOfPath(path string) string {
	// Remove leading and trailing slashes if exist
	path = strings.Trim(path, "/")

	// Split the path
	parts := strings.Split(path, "/")

	// Return the last part if exists
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}
