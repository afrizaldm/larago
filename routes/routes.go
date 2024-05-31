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

func (e Engine) RESOURCES(uri string, ctr controllers.Resources) {

	part := getLastPartOfPath(uri)

	e.GET(uri, ctr.Index)
	e.GET(uri+"/create", ctr.Create)
	e.POST(uri, ctr.Store)
	e.GET(uri+"/:"+part, ctr.Show)
	e.GET(uri+"/:"+part+"/edit", ctr.Edit)
	e.PUT(uri+"/:"+part, ctr.Update)
	e.PATCH(uri+"/:"+part, ctr.Update)
	e.DELETE(uri+"/:"+part, ctr.Destroy)

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
