package usecase

import (
	"testing"
	"simple-note-api/domain"
)

type MockUserRepository struct{}

func (_ MockUserRepository) FindAll() ([]domain.User, error) {
	users := []domain.User{
		{1, "foo", ""},
		{2, "bar", ""},
		{3, "baz", ""},
	}
	return users, nil
}

func TestUserInteractor_Users(t *testing.T) {
	interactor := UserInteractor{
		UserRepository: &MockUserRepository{},
	}

	users, err := interactor.Users()

	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 3 {
		t.Fatalf("number of users expected 3, but got %v", len(users))
	}

	if users[0].ID != 1 || users[1].ID != 2 || users[2].ID != 3 {
		t.Fatalf("sorted by incorrect order: %+v", users)
	}
}
