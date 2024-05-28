package routes

import (
	"simple-api/controllers/api"
)

func (r *Router) SetupApi() {

	// Daftarkan route dan handler
	r.Engine.POST("/api/register", api.RegisterUser)
	r.Engine.GET("/api", api.Dashboard)
	r.Engine.GET("/api/hello", api.Hallo)
}
