package controllers

import "github.com/gin-gonic/gin"

type Resources interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Store(c *gin.Context)
	Show(c *gin.Context)
	Edit(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
