package models

type User struct {
	Guid         string
	Email        string
	Ip           *string
	RefreshToken *string
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
