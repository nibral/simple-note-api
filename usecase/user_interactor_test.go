package usecase

import (
	"reflect"
	"testing"

	"simple-note-api/domain"
)

var senderAdmin = domain.User{
	ID:    1,
	Name:  "foo",
	Admin: true,
}
var senderUser = domain.User{
	ID:    2,
	Name:  "bar",
	Admin: false,
}
var userInteractor = UserInteractor{
	UserRepository: &MockUserRepository{},
}

func TestUserInteractor_List(t *testing.T) {
	actual, err := userInteractor.List(senderAdmin)

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

func TestUserInteractor_Get(t *testing.T) {
	expected1 := domain.User{
		ID:    1,
		Name:  "foo",
		Admin: true,
	}

	actual1, err1 := userInteractor.Get(senderAdmin, 1)

	if err1 != nil {
		t.Fatal(err1)
	}
	actual1.Password = ""
	if !reflect.DeepEqual(actual1, expected1) {
		t.Fatalf("unexpeted user: expected %+v, actual %+v", expected1, actual1)
	}

	expected2 := domain.User{
		ID:    2,
		Name:  "bar",
		Admin: false,
	}

	actual2, err2 := userInteractor.Get(senderUser, 2)

	if err2 != nil {
		t.Fatal(err2)
	}
	actual2.Password = ""
	if !reflect.DeepEqual(actual2, expected2) {
		t.Fatalf("unexpeted user: expected %+v, actual %+v", expected2, actual2)
	}

	actual3, err3 := userInteractor.Get(senderUser, 3)

	if err3 == nil {
		t.Fatalf("user read without admin privileges: %+v", actual3)
	}
	switch err3.(type) {
	case *NotPermittedError:
		break
	default:
		t.Fatal(err2)
	}
}

func TestUserInteractor_Create(t *testing.T) {
	param := domain.User{
		Name:     "qux",
		Password: "password",
		Admin:    false,
	}
	expected1 := domain.User{
		ID:    4,
		Name:  "qux",
		Admin: false,
	}

	actual1, err1 := userInteractor.Create(senderAdmin, param)

	if err1 != nil {
		t.Fatal(err1)
	}
	actual1.Password = ""
	if !reflect.DeepEqual(actual1, expected1) {
		t.Fatalf("unexpected created user: expected %+v, actual %+v", expected1, actual1)
	}

	param.Name = "quux"
	actual2, err2 := userInteractor.Create(senderUser, param)

	if err2 == nil {
		t.Fatalf("user created without admin privileges: %+v", actual2)
	}
	switch err2.(type) {
	case *NotPermittedError:
		break
	default:
		t.Fatal(err2)
	}
}

func TestUserInteractor_Update(t *testing.T) {
	param := domain.User{
		Name:  "foo-ooo",
		Admin: true,
	}
	expected1 := domain.User{
		ID:    1,
		Name:  "foo-ooo",
		Admin: true,
	}

	actual1, err1 := userInteractor.Update(senderAdmin, 1, param)
	if err1 != nil {
		t.Fatal(err1)
	}
	actual1.Password = ""
	if !reflect.DeepEqual(actual1, expected1) {
		t.Fatalf("unexpected updated user: expected %+v, actual %+v", expected1, actual1)
	}

	actual2, err2 := userInteractor.Update(senderUser, 1, param)
	if err2 == nil {
		t.Fatalf("user updated without admin privileges: %+v", actual2)
	}
	switch err2.(type) {
	case *NotPermittedError:
		break
	default:
		t.Fatal(err2)
	}

	actual3, err3 := userInteractor.Update(senderAdmin, 9, param)
	if err3 == nil {
		t.Fatalf("unexpected updated by id = 9 (unused id): %+v", actual3)
	}
	switch e := err3.(type) {
	case *InvalidParameterError:
		t.Log(e.Msg)
		break
	default:
		t.Fatal(err3)
	}

	param4 := domain.User{
		Name:  "bar",
		Admin: false,
	}
	actual4, err4 := userInteractor.Update(senderAdmin, 3, param4)
	if err4 == nil {
		t.Fatalf("unexpected updated by name = bar (same name): %+v", actual4)
	}
	switch e := err4.(type) {
	case *InvalidParameterError:
		t.Log(e.Msg)
		break
	default:
		t.Fatal(err4)
	}

	param5 := domain.User{
		Name:  "foo5",
		Admin: false,
	}
	actual5, err5 := userInteractor.Update(senderAdmin, 1, param5)
	if err5 == nil {
		t.Fatalf("unexpected updated by id = 1, admin = false (degrade myself): %+v", actual5)
	}
	switch e := err5.(type) {
	case *InvalidParameterError:
		t.Log(e.Msg)
		break
	default:
		t.Fatal(err5)
	}
}
