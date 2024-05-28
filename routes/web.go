package routes

import (
	"simple-api/controllers/web"
)

func (r *Router) SetupWeb() {

	r.Engine.LoadHTMLGlob("views/**.html")

	r.Engine.GET("/", web.Dashboard)
}
