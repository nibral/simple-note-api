package database

import (
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

func (_ *MockDatabaseHandler) GetUserByID(id int) (domain.User, error) {
	for _, v := range users {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.User{}, fmt.Errorf("user id %v doesn't exists", id)
}

func (_ *MockDatabaseHandler) GetUserByName(name string) (domain.User, error) {
	for _, v := range users {
		if v.Name == name {
			return v, nil
		}
	}
	return domain.User{}, fmt.Errorf("user %v doesn't exists", name)
}

func (_ *MockDatabaseHandler) GetUserCountByID(id int) (int, error) {
	count := 0
	for _, v := range users {
		if v.ID == id {
			count = count + 1
		}
	}
	return count, nil
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

func (_ *MockDatabaseHandler) PutUser(user domain.User) error {
	users = append(users, user)

	return nil
}

func (_ *MockDatabaseHandler) UpdateUser(id int, user domain.User) error {
	for _, v := range users {
		if v.ID == id {
			v = user
			return nil
		}
	}

	return fmt.Errorf("user id %v doesn't exists", user.ID)
}
