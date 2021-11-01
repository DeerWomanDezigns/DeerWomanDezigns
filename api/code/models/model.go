package models

type User struct {
	ID        string `dynamo:"user_id" json:"id"`
	FirstName string `dynamo:"first_name" json:"firstName"`
	LastName  string `dynamo:"last_name" json:"lastName"`
	Address   string `dynamo:"address" json:"address"`
	Email     string `dynamo:"email" json:"email"`
}

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
