package usecase

import "simple-note-api/domain"

type MockUserRepository struct{}

func (_ MockUserRepository) FindAll() ([]domain.User, error) {
	// password is "password"
	users := []domain.User{
		{1, "foo", "$2a$10$on1KC3u/UkffLpnAZcGeH.smUxQjPw.QzOv2/JRfgIGroEL/uhXwa"},
		{2, "bar", "$2a$10$fr9ED1p/3WbgK5K3lOkz6OBy5PGtII/UYw2Fvl/wlT.m4ADnyMPO."},
		{3, "baz", "$2a$10$E.I05M604bElt6I8h9gMRuMFl9bq.1CBFToEod4CA/Mp8fTi9EDiK"},
	}
	return users, nil
}

func (_ MockUserRepository) FindByName(name string) (domain.User, error) {
	if name == "foo" {
		fooUser := domain.User{
			ID:       1,
			Name:     "foo",
			Password: "$2a$10$on1KC3u/UkffLpnAZcGeH.smUxQjPw.QzOv2/JRfgIGroEL/uhXwa"}
		return fooUser, nil
	} else {
		return domain.User{}, nil
	}
}
