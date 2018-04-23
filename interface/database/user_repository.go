package database

import (
	"sort"

	"simple-note-api/domain"
	"simple-note-api/infrastructure/database"
	"simple-note-api/util"
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

func (repository *UserRepository) FindByID(id int) (domain.User, error) {
	count, err := repository.DatabaseHandler.GetUserCountByID(id)
	if err != nil {
		return domain.User{}, err
	}
	if count == 0 {
		return domain.User{}, nil
	}

	return repository.DatabaseHandler.GetUserByID(id)
}

func (repository *UserRepository) FindByName(name string) (domain.User, error) {
	count, err := repository.DatabaseHandler.GetUserCountByName(name)
	if err != nil {
		return domain.User{}, err
	}
	if count == 0 {
		return domain.User{}, nil
	}

	return repository.DatabaseHandler.GetUserByName(name)
}

func (repository *UserRepository) Add(user domain.User) (domain.User, error) {
	id, err := repository.DatabaseHandler.GetNewUserID()
	if err != nil {
		return domain.User{}, err
	}
	user.ID = id

	err = repository.DatabaseHandler.PutUser(user)
	return user, err
}

func (repository *UserRepository) Update(id int, user domain.User) (domain.User, error) {
	err := repository.DatabaseHandler.UpdateUser(id, user)
	return user, err
}

func (repository *UserRepository) Delete(id int) error {
	return repository.DatabaseHandler.DeleteUser(id)
}
