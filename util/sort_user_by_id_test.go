package util

import (
	"testing"
	"sort"
	"simple-note-api/domain"
)

func TestSortUserById(t *testing.T) {
	users := []domain.User{
		{3, "baz", "", false},
		{1, "foo", "", true},
		{2, "bar", "", false},
	}

	sort.Sort(SortUserById(users))

	if users[0].ID != 1 || users[1].ID != 2 || users[2].ID != 3 {
		t.Fatalf("sorted by incorrect order: %+v", users)
	}
}
