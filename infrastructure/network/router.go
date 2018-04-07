package network

import (
	"simple-note-api/interface/controller"

	"github.com/labstack/echo"
)

var Router *echo.Echo

func init() {
	router := echo.New()

	userController := controller.NewUserController()
	users := router.Group("users")
	users.GET("", userController.Index)

	Router = router
}
