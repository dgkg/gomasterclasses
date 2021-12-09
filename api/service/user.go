package service

import (
	"net/http"

	"github.com/dgkg/gomasterclasses/api/auth"
	"github.com/dgkg/gomasterclasses/api/db"
	"github.com/dgkg/gomasterclasses/api/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var _ Handler = &UserHandlers{}

type UserHandlers struct {
	db  db.Storage
	log *zap.SugaredLogger
}

func (uh *UserHandlers) InitRoutes(r *gin.Engine) {
	gin.Logger()
	r.GET("/users", uh.GetAll)
	r.POST("/users", uh.Create)
	r.PUT("/users/:user_id", uh.Update)
	r.GET("/users/:user_id", uh.Get)
	r.POST("/login", uh.Login)
}

func (uh *UserHandlers) GetAll(c *gin.Context) {
	us, err := uh.db.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": us,
	})
}

func (uh *UserHandlers) Create(c *gin.Context) {
	var u model.User
	err := c.BindJSON(&u)
	if err != nil {
		uh.log.Info(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err = uh.db.CreateUser(&u)
	if err != nil {
		uh.log.Info(err.Error())
		if err == db.ErrUserAlreadyExists {
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (uh *UserHandlers) Update(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.BindJSON(&data)
	if err != nil {
		if err == db.ErrUserNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := uh.db.Update(c.Param("user_id"), data)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (uh *UserHandlers) Get(c *gin.Context) {
	u, err := uh.db.Get(c.Param("user_id"))
	if err != nil {
		if err == db.ErrUserNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (uh *UserHandlers) Login(c *gin.Context) {
	var payload model.LoginUser
	err := c.BindJSON(&payload)
	if err != nil || len(payload.Email) == 0 || len(payload.Password) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := uh.db.GetByEmail(payload.Email)
	if err != nil || u.Password != payload.Password {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtValue, err := auth.JWTSign(u.UUID, u.FirstName+" "+u.LastName)
	if err != nil || u.Password != payload.Password {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt": jwtValue,
	})
}
