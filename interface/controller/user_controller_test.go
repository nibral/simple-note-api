package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"simple-note-api/domain"
	"simple-note-api/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var userController = UserController{
	Interactor: usecase.UserInteractor{
		UserRepository: &usecase.MockUserRepository{},
	},
}
var senderAdmin = domain.User{
	ID:    1,
	Name:  "foo",
	Admin: true,
}

func TestUserController_Index(t *testing.T) {
	expected := `[{"id":1,"name":"foo","admin":true},{"id":2,"name":"bar","admin":false},{"id":3,"name":"baz","admin":false}]`

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(senderAdmin.ID)
	claims["name"] = senderAdmin.Name
	claims["admin"] = senderAdmin.Admin
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	c.Set("user", token)

	err := userController.Index(c)

	if err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %v", rec.Code)
	}
	if rec.Body.String() != expected {
		t.Fatalf("unexpected json response: %v", rec.Body.String())
	}
}

func TestUserController_Get(t *testing.T) {
	expected := `{"id":1,"name":"foo","admin":true}`

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(senderAdmin.ID)
	claims["name"] = senderAdmin.Name
	claims["admin"] = senderAdmin.Admin
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")
	c.Set("user", token)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := userController.Get(c)

	if err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %v", rec.Code)
	}
	if rec.Body.String() != expected {
		t.Fatalf("unexpected json response: %v", rec.Body.String())
	}
}

func TestUserController_Create(t *testing.T) {
	paramJson := `{"name":"qux","password":"password"}`
	expected := `{"id":4,"name":"qux","admin":false}`

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(senderAdmin.ID)
	claims["name"] = senderAdmin.Name
	claims["admin"] = senderAdmin.Admin
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(paramJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	c.Set("user", token)

	err := userController.Create(c)

	if err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status: %v", rec.Code)
	}
	if rec.Body.String() != expected {
		t.Fatalf("unexpected json response: %v", rec.Body.String())
	}

	users, err := userController.Interactor.List(senderAdmin)

	if len(users) != 4 {
		t.Fatalf("number of users expected 4, but got %v", len(users))
	}
}

func TestUserController_Update(t *testing.T) {
	paramJson := `{"name":"bar-bar","password":"","admin":true}`
	expected := `{"id":2,"name":"bar-bar","admin":true}`

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(senderAdmin.ID)
	claims["name"] = senderAdmin.Name
	claims["admin"] = senderAdmin.Admin
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(paramJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")
	c.Set("user", token)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := userController.Update(c)

	if err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %v", rec.Code)
	}
	if rec.Body.String() != expected {
		t.Fatalf("unexpected json response: %v", rec.Body.String())
	}

	users, err := userController.Interactor.List(senderAdmin)

	if len(users) != 4 {
		t.Fatalf("number of users expected 4, but got %v", len(users))
	}
}

func TestUserController_Delete(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(senderAdmin.ID)
	claims["name"] = senderAdmin.Name
	claims["admin"] = senderAdmin.Admin
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")
	c.Set("user", token)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := userController.Delete(c)

	if err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %v", rec.Code)
	}

	users, err := userController.Interactor.List(senderAdmin)

	if len(users) != 3 {
		t.Fatalf("number of users expected 3, but got %v", len(users))
	}
}
