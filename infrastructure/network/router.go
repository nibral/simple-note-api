package network

import (
	"simple-note-api/interface/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Router *echo.Echo

func init() {
	notRestrictedEndpoint := []string{
		"login",
	}

	router := echo.New()
	api := router.Group("/v1")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			for _, endpoint := range notRestrictedEndpoint {
				if c.Request().RequestURI == "/v1/"+endpoint {
					return true
				}
			}
			return false
		},
	}))

	loginController := controller.NewLoginController()
	api.POST("/login", loginController.Login)

	userController := controller.NewUserController()
	users := api.Group("/users")
	users.GET("", userController.Index)

	Router = router
}
