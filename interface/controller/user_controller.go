package controller

import (
	"log"
	"net/http"

	"simple-note-api/domain"
	"simple-note-api/interface/database"
	"simple-note-api/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

type userParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

func NewUserController() *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: database.NewUserRepository(),
		},
	}
}

func (controller *UserController) Index(context echo.Context) error {
	sender := ParseToken(context.Get("user").(*jwt.Token))

	users, err := controller.Interactor.Users(sender)
	if err != nil {
		log.Println(err)
		return context.NoContent(http.StatusInternalServerError)
	}
	return context.JSON(http.StatusOK, users)
}

func (controller *UserController) Create(context echo.Context) error {
	sender := ParseToken(context.Get("user").(*jwt.Token))

	params := userParams{}
	if err := context.Bind(&params); err != nil {
		return context.NoContent(http.StatusBadRequest)
	}
	userParam := domain.User{
		Name:     params.Name,
		Password: params.Password,
		Admin:    params.Admin,
	}

	user, err := controller.Interactor.Create(sender, userParam)
	if err != nil {
		log.Println(err)
		switch e := err.(type) {
		case *usecase.NotPermittedError:
			return context.String(http.StatusForbidden, e.Msg)
		case *usecase.UserCreateError:
			return context.String(http.StatusBadRequest, e.Msg)
		default:
			return context.NoContent(http.StatusInternalServerError)
		}
	}

	return context.JSON(http.StatusCreated, user)
}
