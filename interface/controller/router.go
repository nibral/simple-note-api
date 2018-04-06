package controller

import "github.com/labstack/echo"

type Router struct{}

func NewRouter() *echo.Echo {
	router := echo.New()

	userController := NewUserController()
	users := router.Group("users")
	users.GET("", userController.Index)

	return router
}
