package main

import (
	"github.com/labstack/echo/middleware"
	"simple-note-api/interface/controller"
)

func main() {
	router := controller.NewRouter()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Start(":3000")
}
