package controller

import (
	"net/http"
	"simple-note-api/usecase"
	"simple-note-api/interface/database"
	"github.com/labstack/echo"
)

type LoginController struct {
	Interactor usecase.LoginInteractor
}

type loginParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func NewLoginController() *LoginController {
	return &LoginController{
		Interactor: usecase.LoginInteractor{
			UserRepository: database.NewUserRepository(),
		},
	}
}

func (controller *LoginController) Login(context echo.Context) error {
	params := loginParams{}
	if err := context.Bind(&params); err != nil {
		return context.NoContent(http.StatusBadRequest)
	}

	user, _, err := controller.Interactor.Login(params.Name, params.Password)

	if err != nil {
		return echo.ErrUnauthorized
	}

	return context.JSON(http.StatusOK, user)
}
