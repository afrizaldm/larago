package routes

import "simple-api/app/http/controllers/api"

type Api struct {
	Router
}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) Setup(r *Engine) {

	// Daftarkan route dan handler
	// r.GET("/api", api.Dashboard)
	// r.GET("/api/hello", api.Hallo)

	// r.RESOURCES("/api/user", &api.UserController{})

	// r.RESOURCES("/api/car", &api.CarController{})

	api_g := r.Group("api")
	{
		api_g.RouterGroup.GET("/hello", api.Hallo)

		api_g.RESOURCES("/user", &api.UserController{})

		api_g.RESOURCES("/car", &api.CarController{})
	}

}
