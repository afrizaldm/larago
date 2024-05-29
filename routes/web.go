package routes

import (
	"simple-api/controllers/web"

	"github.com/gin-gonic/gin"
)

type Web struct{}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Setup(r *gin.Engine) {

	r.LoadHTMLGlob("views/**.html")

	r.GET("/", web.Dashboard)
}
