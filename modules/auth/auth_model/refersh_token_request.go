package auth_model

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" form:"refreshToken"`
}
