package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConfig_UnMarshal(t *testing.T) {
	configJson := `{"port":3000,"jwt_secret":"secret","jwt_lifetime_sec":900}`
	expected := Config{
		Port:        3000,
		JwtSecret:   "secret",
		JwtLifetime: 900,
	}

	var actual Config
	err := json.Unmarshal([]byte(configJson), &actual)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected config object: expect %v, actual %v", expected, actual)
	}
}
