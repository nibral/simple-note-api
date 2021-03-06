package database

import (
	"simple-note-api/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoDBHandler struct {
	SequenceTable dynamo.Table
	UserTable     dynamo.Table
}

type Sequence struct {
	Current int `dynamo:"current"`
}

func NewDynamoDBHandler() *DynamoDBHandler {
	dbSession, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	db := dynamo.New(dbSession, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	return &DynamoDBHandler{
		SequenceTable: db.Table("simple-note_sequences"),
		UserTable:     db.Table("simple-note_users"),
	}
}

func (database *DynamoDBHandler) GetAllUsers() ([]domain.User, error) {
	var results []domain.User
	err := database.UserTable.Scan().All(&results)

	return results, err
}

func (database *DynamoDBHandler) GetNewUserID() (int, error) {
	var result Sequence
	err := database.SequenceTable.Update("name", "simple-note_users").Add("current", 1).Value(&result)

	return result.Current, err
}

func (database *DynamoDBHandler) GetUserByID(id int) (domain.User, error) {
	var result domain.User
	err := database.UserTable.Get("id", id).One(&result)

	return result, err
}

func (database *DynamoDBHandler) GetUserByName(name string) (domain.User, error) {
	var result domain.User
	err := database.UserTable.Get("name", name).Index("name-index").One(&result)

	return result, err
}

func (database *DynamoDBHandler) GetUserCountByID(id int) (int, error) {
	count, err := database.UserTable.Get("id", id).Count()

	return int(count), err
}

func (database *DynamoDBHandler) GetUserCountByName(name string) (int, error) {
	count, err := database.UserTable.Get("name", name).Index("name-index").Count()

	return int(count), err
}

func (database *DynamoDBHandler) PutUser(user domain.User) error {
	err := database.UserTable.Put(user).Run()

	return err
}

func (database *DynamoDBHandler) UpdateUser(id int, user domain.User) error {
	err := database.UserTable.Update("id", id).
		Set("name", user.Name).
		Set("password", user.Password).
		Set("admin", user.Admin).
		Run()

	return err
}

func (database *DynamoDBHandler) DeleteUser(id int) error {
	err := database.UserTable.Delete("id", id).Run()

	return err
}
