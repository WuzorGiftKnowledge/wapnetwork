package models


type LoginRequest struct {
	UserName   string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	
	ResfreshToken string `json:"refreshToken"`
}