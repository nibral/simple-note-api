package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"simple-note-api/domain"
)

type SimpleNoteDatabase struct {
	UserTable dynamo.Table
}

func NewSimpleNoteDatabase() *SimpleNoteDatabase {
	dbSession, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	db := dynamo.New(dbSession, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	return &SimpleNoteDatabase{
		UserTable: db.Table("simple-note_users"),
	}
}

func (database *SimpleNoteDatabase) GetAllUsers() ([]domain.User, error) {
	var results []domain.User
	err := database.UserTable.Scan().All(&results)

	return results, err
}
