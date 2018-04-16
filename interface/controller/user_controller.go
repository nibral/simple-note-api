package controller

import (
	"log"
	"net/http"
	"simple-note-api/interface/database"
	"simple-note-api/usecase"
	"github.com/labstack/echo"
	"simple-note-api/domain"
	"github.com/dgrijalva/jwt-go"
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
		log.Println(err)
		return context.NoContent(http.StatusInternalServerError)
	}
	return context.JSON(http.StatusOK, users)
}

func (controller *UserController) Create(context echo.Context) error {
	sender := context.Get("user").(*jwt.Token)
	claims := sender.Claims.(jwt.MapClaims)
	admin := claims["admin"].(bool)
	if !admin {
		return context.NoContent(http.StatusForbidden)
	}

	userParam := domain.User{}
	if err := context.Bind(&userParam); err != nil {
		return context.NoContent(http.StatusBadRequest)
	}

	user, err := controller.Interactor.Create(userParam)
	if err != nil {
		log.Println(err)
		switch err.(type) {
		case *usecase.UserCreateError:
			return context.NoContent(http.StatusBadRequest)
		default:
			return context.NoContent(http.StatusInternalServerError)
		}
	}

	return context.JSON(http.StatusCreated, user)
}
