package usecase

import (
	"testing"

	"simple-note-api/domain"
)

var senderAdmin = domain.User{
	ID:    1,
	Name:  "foo",
	Admin: true,
}
var senderUser = domain.User{
	ID:    2,
	Name:  "bar",
	Admin: false,
}
var userInteractor = UserInteractor{
	UserRepository: &MockUserRepository{},
}

func TestUserInteractor_Users(t *testing.T) {
	actual, err := userInteractor.Users(senderAdmin)

	if err != nil {
		t.Fatal(err)
	}
	if len(actual) != 3 {
		t.Fatalf("unexpected number of users: expect 3, actual %v", len(actual))
	}
	if actual[0].ID != 1 || actual[1].ID != 2 || actual[2].ID != 3 {
		t.Fatalf("sorted by incorrect order: %+v", actual)
	}
}

func TestUserInteractor_Add(t *testing.T) {
	param := domain.User{
		Name:     "qux",
		Password: "password",
		Admin:    false,
	}

	actual1, err1 := userInteractor.Create(senderAdmin, param)

	if err1 != nil {
		t.Fatal(err1)
	}
	if actual1.ID != 4 {
		t.Fatalf("ID of new user expected 4, but got %v", actual1.ID)
	}

	param.Name = "quux"
	actual2, err2 := userInteractor.Create(senderUser, param)

	if err2 == nil {
		t.Fatalf("user created without admin privileges: %+v", actual2)
	}
	switch err2.(type) {
	case *NotPermittedError:
		break
	default:
		t.Fatal(err2)
	}
}
