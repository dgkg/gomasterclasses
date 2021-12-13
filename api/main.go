package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/dgkg/gomasterclasses/api/db"
	"github.com/dgkg/gomasterclasses/api/db/postgres"
	"github.com/dgkg/gomasterclasses/api/db/sqlite"
	"github.com/dgkg/gomasterclasses/api/model"
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
		db = postgres.New()
	} else {
		db = sqlite.New("test.db")
	}
	db.CreateUser(&model.User{
		FirstName: "Bob",
		LastName:  "L'Eponge",
		Email:     "test2@email.com",
		Password:  model.Password(fmt.Sprintf("%x", sha256.Sum256([]byte("123password")))),
	})
	service.New(db, log).InitRoutes().Run("9090")
}
