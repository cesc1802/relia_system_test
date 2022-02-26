package auth_repository

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"relia_system/app_context/tokenprovider"
	"relia_system/modules/auth/auth_model"
	"relia_system/modules/user/user_model"
	"relia_system/shared"
)

type FindUserStorage interface {
	FindUserById(ctx context.Context, id int) (*user_model.User, error)
	UpdateById(ctx context.Context, userId int, data map[string]interface{}) error
}

type rfTokenRepository struct {
	store       FindUserStorage
	tokProvider tokenprovider.Provider
}

func NewRfTokenRepository(store FindUserStorage, tokProvider tokenprovider.Provider) *rfTokenRepository {
	return &rfTokenRepository{store: store, tokProvider: tokProvider}
}

func (repo *rfTokenRepository) RefreshToken(ctx context.Context,
	data *auth_model.RefreshTokenRequest) (*auth_model.LoginResponse, error) {

	pl, err := repo.tokProvider.Validate(data.RefreshToken)
	if err != nil {
		return nil, shared.ErrInvalidRequest(errors.New("refresh token invalid"))
	}

	user, err := repo.store.FindUserById(ctx, pl.GetUserId())
	if err != nil {
		return nil, shared.ErrCannotGetEntity(user_model.EntityName, err)
	}

	if user.IsDeActivate() {
		return nil, user_model.ErrUserHasBeenBlocked
	}

	//if !user.VerifyRefreshTokenId(pl.GetRefreshTokenId()) {
	//	return nil, shared.ErrInvalidRequest(errors.New("refresh token id invalid"))
	//}

	//user.RefreshTokenProcess()

	//if err = repo.store.UpdateById(ctx, user.Id,
	//	map[string]interface{}{"refresh_token_id": user.RefreshTokenId}); err != nil {
	//	return nil, shared.ErrCannotUpdateEntity(user_model.EntityName, err)
	//}

	accessToken, err := repo.tokProvider.GenAccessToken(user, 15*60*60)
	if err != nil {
		return nil, shared.ErrInternal(err)
	}

	refreshToken, err := repo.tokProvider.GenRefreshToken(user, 60*60*24*2)
	if err != nil {
		return nil, shared.ErrInternal(err)
	}

	return auth_model.NewLoginResponse(accessToken, refreshToken), nil
}
