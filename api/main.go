package main

import (
	"go.uber.org/zap"

	"github.com/dgkg/gomasterclasses/api/db"
	"github.com/dgkg/gomasterclasses/api/db/postgres"
	"github.com/dgkg/gomasterclasses/api/db/sqlite"
	"github.com/dgkg/gomasterclasses/api/service"
)

func main() {
	//ENV := os.Getenv("API_ENV")
	ENV := "production"
	var db db.Storage
	if ENV == "production" {
		db = postgres.New()
	} else {
		db = sqlite.New("test.db")
	}
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	service.New(db, log).InitRoutes().Run("9090")
}
