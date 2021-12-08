package db

import (
	"errors"

	"github.com/dgkg/gomasterclasses/api/model"
)

type Storage interface {
	CreateUser(u *model.User) (*model.User, error)
	Delete(uuid string) error
	Update(uuid string, data map[string]interface{}) (*model.User, error)
	Get(uuid string) (*model.User, error)
	GetAll() (map[string]*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

var (
	ErrUserNotFound      = errors.New("db: user not found")
	ErrUserAlreadyExists = errors.New("db: user allready exists")
)
