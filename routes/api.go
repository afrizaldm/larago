package routes

import (
	"simple-api/app/http/controllers/api"
)

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) Setup(r *Engine) {

	// Daftarkan route dan handler
	r.GET("/api", api.Dashboard)
	r.GET("/api/hello", api.Hallo)

	r.RESOURCES("/api/user", &api.UserController{})

	r.RESOURCES("/api/car", &api.CarController{})

}
