package main

import (
	"github.com/dgkg/gomasterclasses/api/service"
)

func main() {
	service.New().InitRoutes().Run("9090")
}
