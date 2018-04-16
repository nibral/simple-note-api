package usecase

import "simple-note-api/domain"

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

func (_ MockUserRepository) FindByName(name string) (domain.User, error) {
	if name == "foo" {
		return users[0], nil
	} else {
		return domain.User{}, nil
	}
}

func (_ MockUserRepository) Add(user domain.User) (domain.User, error) {
	users = append(users, domain.User{
		ID:       4,
		Name:     user.Name,
		Password: "",
		Admin:    user.Admin,
	})
	return users[3], nil
}
