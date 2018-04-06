package util

import "simple-note-api/domain"

type SortUserById []domain.User

func (a SortUserById) Len() int {
	return len(a)
}

func (a SortUserById) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SortUserById) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}
