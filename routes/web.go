package routes

import (
	"simple-api/app/http/controllers/web"
)

type Web struct{}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Setup(r *Engine) {

	r.GET("/", web.Dashboard)
	r.GET("/ping", web.Ping)
}
