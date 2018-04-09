package usecase

import "simple-note-api/domain"

type MockUserRepository struct{}

func (_ MockUserRepository) FindAll() ([]domain.User, error) {
	users := []domain.User{
		{1, "foo", "password"},
		{2, "bar", "password"},
		{3, "baz", "password"},
	}
	return users, nil
}

func (_ MockUserRepository) FindByName(name string) (domain.User, error) {
	if name == "foo" {
		fooUser := domain.User{
			ID:       1,
			Name:     "foo",
			Password: "password"}
		return fooUser, nil
	} else {
		return domain.User{}, nil
	}
}
