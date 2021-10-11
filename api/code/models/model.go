package models

type User struct {
	ID        string `dynamo:"user_id"`
	FirstName string `dynamo:"first_name"`
	LastName  string `dynamo:"last_name"`
	Address   string `dynamo:"address"`
	Email     string `dynamo:"email"`
}
