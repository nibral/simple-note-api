package controller

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"simple-note-api/usecase"
	"net/http"
)

func TestLoginController_Login(t *testing.T) {
	paramJson := `{"name":"foo","password":"password"}`
	responseJson := `{"id":1,"name":"foo"}`

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(paramJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	controller := LoginController{
		Interactor: usecase.LoginInteractor{
			UserRepository: &usecase.MockUserRepository{},
		},
	}
	err := controller.Login(c)

	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %v", rec.Code)
	}

	if rec.Body.String() != responseJson {
		t.Fatalf("unexpected json response: %v", rec.Body.String())
	}
}
