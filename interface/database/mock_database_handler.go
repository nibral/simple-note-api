package database

import "simple-note-api/domain"

type MockDatabaseHandler struct{}

func (_ *MockDatabaseHandler) GetAllUsers() ([]domain.User, error) {
	users := []domain.User{
		{3, "baz", "", false},
		{1, "foo", "", true},
		{2, "bar", "", false},
	}
	return users, nil
}

func (_ *MockDatabaseHandler) GetUserByName(name string) (domain.User, error) {
	if name == "foo" {
		fooUser := domain.User{
			ID:       1,
			Name:     "foo",
			Password: "",
			Admin:    true,
		}
		return fooUser, nil
	} else {
		return domain.User{}, nil
	}
}

func (_ *MockDatabaseHandler) GetNewUserID() (int, error) {
	return 4, nil
}

func (_ *MockDatabaseHandler) AddUser(param domain.User) error {
	return nil
}
