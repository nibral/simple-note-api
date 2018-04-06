package database

import (
	"simple-note-api/domain"
	"simple-note-api/infrastructure"
	"sort"
	"simple-note-api/util"
)

type UserRepository struct {
	SimpleNoteDatabase SimpleNoteDatabaseInterface
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		SimpleNoteDatabase: infrastructure.NewSimpleNoteDatabase(),
	}
}

func (repository *UserRepository) FindAll() ([]domain.User, error) {
	users, err := repository.SimpleNoteDatabase.GetAllUsers()
	sort.Sort(util.SortUserById(users))
	return users, err
}
