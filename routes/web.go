package routes

import (
	"simple-api/app/http/controllers/web"
)

type Web struct{}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Setup(r *ExtendedEngine) {

	r.GET("/", web.Dashboard)
}
