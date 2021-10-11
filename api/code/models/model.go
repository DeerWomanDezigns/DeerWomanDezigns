package models

type User struct {
	ID        string `dynamo:"user_id" form:"id" xml:"id" binding:"required"`
	FirstName string `dynamo:"first_name" form:"firstName" binding:"required"`
	LastName  string `dynamo:"last_name" form:"lastName" binding:"required"`
	Address   string `dynamo:"address" form:"address" binding:"required"`
	Email     string `dynamo:"email" form:"email" binding:"required"`
}
