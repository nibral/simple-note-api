package main

import (
	"simple-note-api/infrastructure/network"

	"github.com/labstack/echo/middleware"
)

func main() {
	router := network.Router

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Start(":3000")
}
