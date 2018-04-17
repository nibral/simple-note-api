package controller

import (
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo"
	"simple-note-api/usecase"
	"net/http"
	"strings"
	"simple-note-api/domain"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func TestUserController_Index(t *testing.T) {
	json := `[{"id":1,"name":"foo"},{"id":2,"name":"bar"},{"id":3,"name":"baz"}]`

	sender := domain.User{
		ID:    1,
		Name:  "foo",
		Admin: true,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(sender.ID)
	claims["name"] = sender.Name
	claims["admin"] = sender.Admin
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	c.Set("user", token)

	controller := UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &usecase.MockUserRepository{},
		},
	}
	err := controller.Index(c)

	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %v", rec.Code)
	}

	if rec.Body.String() != json {
		t.Fatalf("unexpected json response: %v", rec.Body.String())
	}
}

func TestUserController_Create(t *testing.T) {
	paramJson := `{"name":"qux","password":"password"}`
	userJson := `{"id":4,"name":"qux"}`

	sender := domain.User{
		ID:    1,
		Name:  "foo",
		Admin: true,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(sender.ID)
	claims["name"] = sender.Name
	claims["admin"] = sender.Admin
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(paramJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	c.Set("user", token)

	controller := UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &usecase.MockUserRepository{},
		},
	}
	err := controller.Create(c)

	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status: %v", rec.Code)
	}

	if rec.Body.String() != userJson {
		t.Fatalf("unexpected json response: %v", rec.Body.String())
	}

	users, err := controller.Interactor.Users(sender)

	if len(users) != 4 {
		t.Fatalf("number of users expected 4, but got %v", len(users))
	}
}
