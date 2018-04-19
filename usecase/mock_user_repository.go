package usecase

import (
	"fmt"

	"simple-note-api/domain"
)

type MockUserRepository struct{}

var users []domain.User

func init() {
	// password is "password"
	users = []domain.User{
		{1, "foo", "$2a$10$on1KC3u/UkffLpnAZcGeH.smUxQjPw.QzOv2/JRfgIGroEL/uhXwa", true},
		{2, "bar", "$2a$10$fr9ED1p/3WbgK5K3lOkz6OBy5PGtII/UYw2Fvl/wlT.m4ADnyMPO.", false},
		{3, "baz", "$2a$10$E.I05M604bElt6I8h9gMRuMFl9bq.1CBFToEod4CA/Mp8fTi9EDiK", false},
	}
}

func (_ MockUserRepository) FindAll() ([]domain.User, error) {
	return users, nil
}

func (_ MockUserRepository) FindByID(id int) (domain.User, error) {
	for _, v := range users {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.User{}, nil
}

func (_ MockUserRepository) FindByName(name string) (domain.User, error) {
	for _, v := range users {
		if v.Name == name {
			return v, nil
		}
	}
	return domain.User{}, nil
}

func (_ MockUserRepository) Add(user domain.User) (domain.User, error) {
	users = append(users, domain.User{
		ID:       len(users) + 1,
		Name:     user.Name,
		Password: "",
		Admin:    user.Admin,
	})
	return users[len(users)-1], nil
}

func (_ MockUserRepository) Update(id int, user domain.User) (domain.User, error) {
	for _, v := range users {
		if v.ID == id {
			v = user
			return v, nil
		}
	}

	return domain.User{}, fmt.Errorf("user id %v doesn't exists", user.ID)
}
