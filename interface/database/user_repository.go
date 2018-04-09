package database

import (
	"simple-note-api/domain"
	"simple-note-api/infrastructure/database"
	"simple-note-api/util"
	"sort"
)

type UserRepository struct {
	DatabaseHandler HandlerInterface
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DatabaseHandler: database.NewDynamoDBHandler(),
	}
}

func (repository *UserRepository) FindAll() ([]domain.User, error) {
	users, err := repository.DatabaseHandler.GetAllUsers()
	sort.Sort(util.SortUserById(users))
	return users, err
}
