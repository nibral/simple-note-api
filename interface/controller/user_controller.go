package controller

import (
	"simple-note-api/usecase"
	"simple-note-api/interface/database"
	"github.com/labstack/echo"
	"net/http"
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
