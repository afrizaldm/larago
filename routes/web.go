package routes

import (
	webController "simple-api/app/http/controllers/web"
)

type Web struct {
	Router
}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Setup(r *Engine) {

	r.GET("/", webController.Dashboard)
	r.GET("/ping", webController.Ping)
	// r.GET("users", web.Users)
}
