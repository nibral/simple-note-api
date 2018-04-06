package domain

type User struct {
	ID       int    `dynamo:"id"`
	Name     string `dynamo:"name"`
	Password string `dynamo:"password"`
}
