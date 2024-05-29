package controllers

import "github.com/gin-gonic/gin"

type Resources interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Store(c *gin.Context)
	Show(c *gin.Context)
	Edit(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type ControllerResources struct{}

func (r *ControllerResources) Index(c *gin.Context)   {}
func (r *ControllerResources) Create(c *gin.Context)  {}
func (r *ControllerResources) Store(c *gin.Context)   {}
func (r *ControllerResources) Show(c *gin.Context)    {}
func (r *ControllerResources) Edit(c *gin.Context)    {}
func (r *ControllerResources) Update(c *gin.Context)  {}
func (r *ControllerResources) Destroy(c *gin.Context) {}
