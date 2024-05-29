package routes

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Setup(r *gin.Engine)
}
