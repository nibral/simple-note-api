package database

import (
	"testing"

	"simple-note-api/domain"
)

var repository = UserRepository{
	DatabaseHandler: &MockDatabaseHandler{},
}

func TestUserRepository_FindAll(t *testing.T) {
	actual, err := repository.FindAll()

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

func TestUserRepository_FindByName(t *testing.T) {
	actual1, err1 := repository.FindByName("foo")

	if err1 != nil {
		t.Fatal(err1)
	}
	if actual1.ID != 1 {
		t.Fatalf("unexpected user id: expect 1, actual %v", actual1.ID)
	}

	actual2, err2 := repository.FindByName("qux")
	if err2 != nil {
		t.Fatal(err2)
	}
	if actual2.ID != 0 {
		t.Fatalf("got user by invalid name: %+v", actual2)
	}
}

func TestUserRepository_Add(t *testing.T) {
	param := domain.User{
		Name:     "qux",
		Password: "password",
		Admin:    false,
	}

	actual, err := repository.Add(param)

	if err != nil {
		t.Fatal(err)
	}

	if actual.ID != 4 {
		t.Fatalf("unexpected user id: expect 4, actual %v(%+v)", actual.ID, actual)
	}
}
