package domain

type User struct {
	ID       int    `dynamo:"id" json:"id"`
	Name     string `dynamo:"name" json:"name"`
	Password string `dynamo:"password" json:"-"`
}
