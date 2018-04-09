package controller

import (
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo"
	"simple-note-api/usecase"
	"net/http"
)

func TestUserController_Index(t *testing.T) {
	json := `[{"id":1,"name":"foo"},{"id":2,"name":"bar"},{"id":3,"name":"baz"}]`

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

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
