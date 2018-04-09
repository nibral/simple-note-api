package database

import (
	"testing"
)

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

func TestUserRepository_FindByName(t *testing.T) {
	repository := UserRepository{
		DatabaseHandler: &MockDatabaseHandler{},
	}

	user1, err1 := repository.FindByName("foo")

	if err1 != nil {
		t.Fatal(err1)
	}

	if user1.ID != 1 {
		t.Fatalf("user id expected 1, but got %v.", user1.ID)
	}

	user2, err2 := repository.FindByName("bar")

	if err2 != nil {
		t.Fatal(err2)
	}

	if user2.ID != 0 {
		t.Fatalf("not exist user: %+v", user2)
	}
}
