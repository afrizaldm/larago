package routes

import "github.com/gin-gonic/gin"

type Router interface {
	SetupRouter() *gin.Engine
	GetEngine() *gin.Engine
}
