package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"simple-note-api/usecase"

	"github.com/labstack/echo"
)

var loginController = LoginController{
	Interactor: usecase.LoginInteractor{
		UserRepository: &usecase.MockUserRepository{},
	},
}

func TestLoginController_Login(t *testing.T) {
	paramJson := `{"name":"foo","password":"password"}`

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(paramJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/login")

	err := loginController.Login(c)

	if err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %v", rec.Code)
	}

	actual := loginResult{}
	jsonErr := json.Unmarshal(rec.Body.Bytes(), &actual)

	if jsonErr != nil {
		t.Fatal(jsonErr)
	}
	if actual.ID != 1 {
		t.Fatalf("unexpedted user id: expect 1, actual %v", actual.ID)
	}
}
