package main

import (
	"github.com/thestrayed/place-api/config"
	"github.com/thestrayed/place-api/models"
	"github.com/thestrayed/place-api/router"
)

func main () {
	if err := config.LoadEnv(); err != nil {
		panic(err)
	}

	if err := models.Init(); err != nil {
		panic(err)
	}

	router.Run()
}
