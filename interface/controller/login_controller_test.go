package controller

import (
	"testing"
	"strings"
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"simple-note-api/usecase"
	"github.com/labstack/echo"
)

func TestLoginController_Login(t *testing.T) {
	paramJson := `{"name":"foo","password":"password"}`

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(paramJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/login")

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

	result := new(loginResult)
	jsonErr := json.Unmarshal(rec.Body.Bytes(), result)

	if jsonErr != nil {
		t.Fatal(jsonErr)
	}

	if result.ID != 1 {
		t.Fatalf("user id expected 1, but got %v", result.ID)
	}
}
