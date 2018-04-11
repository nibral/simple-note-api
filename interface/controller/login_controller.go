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

type loginResult struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
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

	user, token, err := controller.Interactor.Login(params.Name, params.Password)

	if err != nil {
		return echo.ErrUnauthorized
	}

	result := loginResult{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}
	return context.JSON(http.StatusOK, result)
}
