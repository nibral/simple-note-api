package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUser_Marshal(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "foo",
		Password: "password",
		Admin:    true,
	}
	expected := `{"id":1,"name":"foo"}`

	actual, err := json.Marshal(user)

	if err != nil {
		t.Fatal(err)
	}
	if string(actual) != expected {
		t.Fatalf("unexpected json string: expect %v, actual %v", expected, string(actual))
	}
}

func TestUser_Unmarshal(t *testing.T) {
	userJson := `{"ID":1,"name":"foo","password":"password","admin":true}`
	expected := User{
		ID:       1,
		Name:     "foo",
		Password: "",    // marked as ignore
		Admin:    false, // marked as ignore
	}

	var actual User
	err := json.Unmarshal([]byte(userJson), &actual)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected user object: expect %v, actual %v", expected, actual)
	}
}
