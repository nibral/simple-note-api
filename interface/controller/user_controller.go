package controller

import (
	"net/http"
	"simple-note-api/interface/database"
	"simple-note-api/usecase"

	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController() *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: database.NewUserRepository(),
		},
	}
}

func (controller *UserController) Index(context echo.Context) error {
	users, err := controller.Interactor.Users()
	if err != nil {
		return context.NoContent(http.StatusInternalServerError)
	}
	return context.JSON(http.StatusOK, users)
}
