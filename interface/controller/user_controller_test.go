package controller

import (
	"net/http/httptest"
	"testing"
	"simple-note-api/domain"
	"github.com/labstack/echo"
	"simple-note-api/usecase"
	"net/http"
)

type MockUserRepository struct{}

func (_ MockUserRepository) FindAll() ([]domain.User, error) {
	users := []domain.User{
		{1, "foo", ""},
	}
	return users, nil
}

func TestUserController_Index(t *testing.T) {
	json := `[{"id":1,"name":"foo"}]`

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	controller := UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &MockUserRepository{},
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
