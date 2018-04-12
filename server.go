package main

import (
	"strconv"
	"simple-note-api/infrastructure/config"
	"simple-note-api/infrastructure/network"

	"github.com/labstack/echo/middleware"
)

func main() {
	appConfig, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	router := network.NewRouter(appConfig)

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Start(":" + strconv.Itoa(appConfig.Port))
}
