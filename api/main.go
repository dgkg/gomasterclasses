package main

import (
	"go.uber.org/zap"

	"github.com/dgkg/gomasterclasses/api/db/moke"
	"github.com/dgkg/gomasterclasses/api/service"
)

func main() {
	db := moke.New()
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	service.New(db, log).InitRoutes().Run("9090")
}
