package auth_repository

import (
	"golang.org/x/net/context"
	"relia_system/app_context/tokenprovider"
	"relia_system/modules/auth/auth_dto"
	"relia_system/modules/auth/auth_model"
	"relia_system/modules/user/user_model"
	"relia_system/shared"
)

type UserMySql interface {
	FindUserByEmail(ctx context.Context, email string) (*user_model.User, error)
}

type loginRepository struct {
	store         UserMySql
	tokenProvider tokenprovider.Provider
	hasher        shared.Hasher
}

func NewLoginRepository(store UserMySql, tokenProvider tokenprovider.Provider, hasher shared.Hasher) *loginRepository {
	return &loginRepository{store: store, tokenProvider: tokenProvider, hasher: hasher}
}

func (repo *loginRepository) UserLogin(ctx context.Context,
	data *auth_dto.LoginRequest) (*auth_model.LoginResponse, error) {

	user, err := repo.store.FindUserByEmail(ctx, data.Email)

	if err != nil {
		return nil, auth_dto.ErrInvalidEmail
	}

	if validPassword := user.Password == repo.hasher.Hash(data.Password+user.Salt); !validPassword {
		return nil, auth_dto.ErrInvalidEmailOrPassword
	}

	accessToken, err := repo.tokenProvider.GenAccessToken(user, 15*60*60)
	if err != nil {
		return nil, shared.ErrInternal(err)
	}

	refreshToken, err := repo.tokenProvider.GenRefreshToken(user, 60*60*24*2)
	if err != nil {
		return nil, shared.ErrInternal(err)
	}

	return auth_model.NewLoginResponse(accessToken, refreshToken), nil
}
