package auth_dto

import "relia_system/app_context/tokenprovider"

type LoginResponse struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewLoginResponse(atok, rtok *tokenprovider.Token) *LoginResponse {
	return &LoginResponse{
		AccessToken:  atok,
		RefreshToken: rtok,
	}
}
