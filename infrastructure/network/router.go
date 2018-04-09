package network

import (
	"simple-note-api/interface/controller"

	"github.com/labstack/echo"
)

var Router *echo.Echo

func init() {
	router := echo.New()
	api := router.Group("/v1")

	loginController := controller.NewLoginController()
	api.POST("/login", loginController.Login)

	userController := controller.NewUserController()
	users := api.Group("/users")
	users.GET("", userController.Index)

	Router = router
}
