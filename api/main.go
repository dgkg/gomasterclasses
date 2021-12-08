package main

import (
	"go.uber.org/zap"

	"github.com/dgkg/gomasterclasses/api/db/sqlite"
	"github.com/dgkg/gomasterclasses/api/service"
)

func main() {
	db := sqlite.New("test.db")
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	service.New(db, log).InitRoutes().Run("9090")
}
