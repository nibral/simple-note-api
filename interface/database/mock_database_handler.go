package database

import "simple-note-api/domain"

type MockDatabaseHandler struct{}

func (_ *MockDatabaseHandler) GetAllUsers() ([]domain.User, error) {
	users := []domain.User{
		{3, "baz", ""},
		{1, "foo", ""},
		{2, "bar", ""},
	}
	return users, nil
}

func (_ *MockDatabaseHandler) GetUserByName(name string) (domain.User, error) {
	if name == "foo" {
		fooUser := domain.User{
			ID:       1,
			Name:     "foo",
			Password: ""}
		return fooUser, nil
	} else {
		return domain.User{}, nil
	}
}