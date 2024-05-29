package routes

import (
	"simple-api/controllers/api"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) Setup(r *gin.Engine) {

	// Daftarkan route dan handler
	r.POST("/api/register", api.RegisterUser)
	r.GET("/api", api.Dashboard)
	r.GET("/api/hello", api.Hallo)
}
