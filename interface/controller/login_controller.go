package controller

import (
	"time"
	"log"
	"net/http"
	"simple-note-api/domain"
	"simple-note-api/usecase"
	"simple-note-api/interface/database"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

type LoginController struct {
	Config     domain.Config
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

func NewLoginController(config domain.Config) *LoginController {
	return &LoginController{
		Config: config,
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

	user, err := controller.Interactor.Login(params.Name, params.Password)
	if err != nil {
		log.Println(err)
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Duration(controller.Config.JwtLifetime) * time.Second).Unix()
	t, err := token.SignedString([]byte(controller.Config.JwtSecret))
	if err != nil {
		log.Println("failed to generate jwt")
		return context.NoContent(http.StatusInternalServerError)
	}

	result := loginResult{
		ID:    user.ID,
		Name:  user.Name,
		Token: t,
	}
	return context.JSON(http.StatusOK, result)
}
