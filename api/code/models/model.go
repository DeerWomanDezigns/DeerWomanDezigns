package models

type User struct {
	ID        string `dynamo:"user_id" json:"id"`
	FirstName string `dynamo:"first_name" json:"firstName"`
	LastName  string `dynamo:"last_name" json:"lastName"`
	Address   string `dynamo:"address" json:"address"`
	Email     string `dynamo:"email" json:"email"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
