package routes

import "github.com/gin-gonic/gin"

type Router interface {
	Setup(r *ExtendedEngine)
}

type ExtendedEngine struct {
	*gin.Engine
}

func (e ExtendedEngine) RESOURCES(uri string, handlers ...gin.HandlerFunc) {

	e.GET(uri)
	e.GET(uri + "/create")
	e.POST(uri)
	e.GET(uri + "/:" + uri)
	e.GET(uri + "/:" + uri + "/edit")
	e.PUT(uri + "/:" + uri)
	e.PATCH(uri + "/:" + uri)
	e.DELETE(uri + "/:" + uri)

}
