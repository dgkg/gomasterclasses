package main

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/dgkg/gomasterclasses/api/db"
	"github.com/dgkg/gomasterclasses/api/db/mongodb"
	"github.com/dgkg/gomasterclasses/api/db/sqlite"
	"github.com/dgkg/gomasterclasses/api/service"
)

var (
	ENV string = "local"
	log *zap.SugaredLogger
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	ENV = viper.GetString("ENV")
	loggerp, err := zap.NewProduction()
	log = loggerp.Sugar()
	if err != nil {
		panic(err)
	}
	defer log.Sync()
	log.Info(fmt.Sprintf("application env:%v", ENV))
}

func main() {
	var db db.Storage
	if ENV == "production" {
		db = mongodb.New()
	} else {
		db = sqlite.New("test.db")
	}
	service.New(db, log).InitRoutes().Run("9090")
}
