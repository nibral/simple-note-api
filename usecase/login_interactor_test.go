package usecase

import (
	"testing"
	"simple-note-api/domain"
)

func TestLoginInteractor_Login(t *testing.T) {
	interactor := LoginInteractor{
		Config: domain.Config{
			JwtSecret: "secret",
		},
		UserRepository: &MockUserRepository{},
	}

	user, _, err := interactor.Login("foo", "password")

	if err != nil {
		t.Fatal(err)
	}

	if user.ID != 1 {
		t.Fatalf("user id expected 1, but got %v", user.ID)
	}
}
