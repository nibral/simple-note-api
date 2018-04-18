package usecase

import (
	"testing"
)

var loginInteractor = LoginInteractor{
	UserRepository: &MockUserRepository{},
}

func TestLoginInteractor_Login(t *testing.T) {
	actual1, err1 := loginInteractor.Login("foo", "password")

	if err1 != nil {
		t.Fatal(err1)
	}
	if actual1.ID != 1 {
		t.Fatalf("unexpected user id: expect 1, actual %v", actual1.ID)
	}

	_, err2 := loginInteractor.Login("foo", "incorrect")

	if err2 == nil {
		t.Fatalf("no error throws with incorrect password.")
	}
}
