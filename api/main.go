package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", GetAllUser)
	r.POST("/users", CreateUser)
	r.PUT("/users/:user_id", UpdateUser)
	r.GET("/users/:user_id", GetUser)
	r.Run(":8080")
}

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": "list users",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "new",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": c.Param("user_id"),
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": c.Param("user_id"),
	})
}
