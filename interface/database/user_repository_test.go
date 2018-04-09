package database

import (
	"testing"
	"simple-note-api/domain"
)

type MockDatabaseHandler struct{}

func (_ *MockDatabaseHandler) GetAllUsers() ([]domain.User, error) {
	users := []domain.User{
		{3, "baz", ""},
		{1, "foo", ""},
		{2, "bar", ""},
	}
	return users, nil
}

func TestUserRepository_FindAll(t *testing.T) {
	repository := UserRepository{
		DatabaseHandler: &MockDatabaseHandler{},
	}

	users, err := repository.FindAll()

	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 3 {
		t.Fatalf("number of users expected 3, but got %v.", len(users))
	}

	if users[0].ID != 1 || users[1].ID != 2 || users[2].ID != 3 {
		t.Fatalf("sorted by incorrect order: %+v", users)
	}
}
