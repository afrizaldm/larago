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

	auth := r.Group("/auth")
	{

		auth.GET("/login", webController.AuthLogin)
		auth.GET("/logout", webController.AuthLogout)
		auth.GET("/check", webController.AuthCheck)
		auth.GET("/ping", webController.AuthPing)
	}

	// r.GET("users", web.Users)
}
