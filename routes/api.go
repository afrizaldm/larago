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

	api.Use(middlewares.Logger())
	// api.Use(middlewares.IsAuthJWT())
	// api.Use(middlewares.IsAuthenticated())
	{
		api.RouterGroup.GET("/", apiController.Ping)
		api.RouterGroup.GET("/ping", apiController.Ping)

		api.RESOURCES("/user", &apiController.UserController{})

		api.RESOURCES("/car", &apiController.CarController{})

		api.RouterGroup.GET("/token/generate", apiController.Generate)
		api.RouterGroup.GET("/token/validate/:token", apiController.Validate)
	}

}
