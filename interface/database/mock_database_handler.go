package database

import (
	"errors"
	"fmt"

	"simple-note-api/domain"
)

type MockDatabaseHandler struct{}

var users []domain.User

func init() {
	users = []domain.User{
		{3, "baz", "", false},
		{1, "foo", "", true},
		{2, "bar", "", false},
	}
}

func (_ *MockDatabaseHandler) GetAllUsers() ([]domain.User, error) {
	return users, nil
}

func (_ *MockDatabaseHandler) GetNewUserID() (int, error) {
	return len(users) + 1, nil
}

func (_ *MockDatabaseHandler) GetUserByName(name string) (domain.User, error) {
	for _, v := range users {
		if v.Name == name {
			return v, nil
		}
	}
	return domain.User{}, errors.New(fmt.Sprintf("%v doesn't exists", name))
}

func (_ *MockDatabaseHandler) GetUserCountByName(name string) (int, error) {
	count := 0
	for _, v := range users {
		if v.Name == name {
			count = count + 1
		}
	}
	return count, nil
}

func (_ *MockDatabaseHandler) AddUser(param domain.User) error {
	return nil
}
