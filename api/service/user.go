package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ Handler = &UserHandlers{}

type UserHandlers struct {
}

func (uh *UserHandlers) InitRoutes(r *gin.Engine) {
	r.GET("/users", uh.GetAll)
	r.POST("/users", uh.Create)
	r.PUT("/users/:user_id", uh.Update)
	r.GET("/users/:user_id", uh.Get)
}

func (uh *UserHandlers) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": "list users",
	})
}

func (uh *UserHandlers) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "new",
	})
}

func (uh *UserHandlers) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": c.Param("user_id"),
	})
}

func (uh *UserHandlers) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": c.Param("user_id"),
	})
}
