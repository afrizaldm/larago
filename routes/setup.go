package routes

import (
	"github.com/gin-gonic/gin"
)

type IRouter interface {
	GetEngine() *gin.Engine
	Setup() *Router
	SetupApi()
	SetupWeb()
}

type Router struct {
	Engine *gin.Engine
}

func (router *Router) Setup() *Router {

	// Instalasi Router Baru
	r := gin.Default()
	router.Engine = r

	// Daftarkan route web
	router.SetupWeb()

	// Daftarkan route api
	router.SetupApi()

	return router
}

func (router *Router) GetEngine() *gin.Engine {
	return router.Engine
}
