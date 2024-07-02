package routes

import (
	apiController "simple-api/app/http/controllers/api"
	"simple-api/app/http/middlewares"
)

type Api struct {
	Router
}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) Setup(r *Engine) {

	api := r.Group("api")

	api.Use(middlewares.Log())
	// api.Use(middlewares.IsAuthenticated())
	{
		api.RouterGroup.GET("/hello", apiController.Hallo)

		api.RESOURCES("/user", &apiController.UserController{})

		api.RESOURCES("/car", &apiController.CarController{})
	}

}
