package database

import (
	"reflect"
	"testing"

	"simple-note-api/domain"
)

var userRepository = UserRepository{
	DatabaseHandler: &MockDatabaseHandler{},
}

func TestUserRepository_FindAll(t *testing.T) {
	actual, err := userRepository.FindAll()

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

func TestUserRepository_FindByID(t *testing.T) {
	actual1, err1 := userRepository.FindByID(1)

	if err1 != nil {
		t.Fatal(err1)
	}
	if actual1.Name != "foo" {
		t.Fatalf("unexpected user name: expect foo, actual %v", actual1.Name)
	}

	actual2, err2 := userRepository.FindByID(9)
	if err2 != nil {
		t.Fatal(err2)
	}
	if actual2.ID != 0 {
		t.Fatalf("got user by invalid id: %+v", actual2)
	}
}

func TestUserRepository_FindByName(t *testing.T) {
	actual1, err1 := userRepository.FindByName("foo")

	if err1 != nil {
		t.Fatal(err1)
	}
	if actual1.ID != 1 {
		t.Fatalf("unexpected user id: expect 1, actual %v", actual1.ID)
	}

	actual2, err2 := userRepository.FindByName("qux")
	if err2 != nil {
		t.Fatal(err2)
	}
	if actual2.ID != 0 {
		t.Fatalf("got user by invalid name: %+v", actual2)
	}
}

func TestUserRepository_Add(t *testing.T) {
	user := domain.User{
		Name:     "qux",
		Password: "password",
		Admin:    false,
	}
	expected := domain.User{
		ID:       4,
		Name:     "qux",
		Password: "password",
		Admin:    false,
	}

	actual, err := userRepository.Add(user)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected new user: expected %+v, actual %+v", expected, actual)
	}
}

func TestUserRepository_Update(t *testing.T) {
	expected := domain.User{
		ID:       1,
		Name:     "foo-ooo",
		Password: "p@ssw0rd",
		Admin:    true,
	}

	actual, err := userRepository.Update(1, expected)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected updated user: expected %+v, actual %+v", expected, actual)
	}
}
