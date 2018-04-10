package database

import (
	"simple-note-api/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoDBHandler struct {
	UserTable dynamo.Table
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
		UserTable: db.Table("simple-note_users"),
	}
}

func (database *DynamoDBHandler) GetAllUsers() ([]domain.User, error) {
	var results []domain.User
	err := database.UserTable.Scan().All(&results)

	return results, err
}

func (database *DynamoDBHandler) GetUserByName(name string) (domain.User, error) {
	var result domain.User
	err := database.UserTable.Get("name", name).Index("name-index").One(&result)

	return result, err
}
