package database

import (
	"simple-note-api/domain"
	"simple-note-api/infrastructure/database"
	"simple-note-api/util"
	"sort"
	"golang.org/x/crypto/bcrypt"
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

func (repository *UserRepository) FindByName(name string) (domain.User, error) {
	return repository.DatabaseHandler.GetUserByName(name)
}

func (repository *UserRepository) Add(param domain.User) (domain.User, error) {
	id, err := repository.DatabaseHandler.GetNewUserID()
	if err != nil {
		return domain.User{}, err
	}
	param.ID = id

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	param.Password = string(hashedPassword)

	err = repository.DatabaseHandler.AddUser(param)
	return param, err
}
