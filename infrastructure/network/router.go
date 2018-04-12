package network

import (
	"simple-note-api/interface/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"simple-note-api/domain"
)

func NewRouter(config domain.Config) *echo.Echo {
	router := echo.New()
	api := router.Group("/v1")

	loginController := controller.NewLoginController(config)
	api.POST("/login", loginController.Login)

	userController := controller.NewUserController()
	users := api.Group("/users")
	users.Use(middleware.JWT([]byte(config.JwtSecret)))
	users.GET("", userController.Index)

	return router
}
