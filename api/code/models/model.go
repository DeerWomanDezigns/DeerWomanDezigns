package models

type User struct {
	ID        string `dynamo:"user_id" json:"id" binding:"required"`
	FirstName string `dynamo:"first_name" json:"firstName" binding:"required"`
	LastName  string `dynamo:"last_name" json:"lastName" binding:"required"`
	Address   string `dynamo:"address" json:"address" binding:"required"`
	Email     string `dynamo:"email" json:"email" binding:"required"`
}
