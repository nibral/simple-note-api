package usecase

import (
	"testing"
	"simple-note-api/domain"
)

var sender = domain.User{
	ID:    1,
	Name:  "foo",
	Admin: true,
}

func TestUserInteractor_Users(t *testing.T) {
	interactor := UserInteractor{
		UserRepository: &MockUserRepository{},
	}

	users, err := interactor.Users(sender)

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

func TestUserInteractor_Add(t *testing.T) {
	interactor := UserInteractor{
		UserRepository: &MockUserRepository{},
	}

	param := domain.User{
		Name:     "qux",
		Password: "password",
		Admin:    false,
	}
	user, err := interactor.Create(sender, param)

	if err != nil {
		t.Fatal(err)
	}

	if user.ID != 4 {
		t.Fatalf("ID of new user expected 4, but got %v", user.ID)
	}
}
