package moke

import "github.com/dgkg/gomasterclasses/api/model"

type MokeDB struct {
	listUser map[string]*model.User
}

func New() *MokeDB {
	var db MokeDB
	db.listUser = make(map[string]*model.User)
	return &db
}
