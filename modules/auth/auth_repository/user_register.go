package auth_repository

import (
	"golang.org/x/net/context"
	"relia_system/app_context/tokenprovider"
	"relia_system/modules/auth/auth_dto"
	"relia_system/modules/user/user_model"
	"relia_system/shared"
)

type RegisterUserMySql interface {
	FindUserByEmail(ctx context.Context, email string) (*user_model.User, error)
	Create(ctx context.Context, data *user_model.User) error
}
type registerRepository struct {
	store         RegisterUserMySql
	tokenProvider tokenprovider.Provider
	hasher        shared.Hasher
}

func NewRegisterRepository(store RegisterUserMySql,
	tokenProvider tokenprovider.Provider,
	hasher shared.Hasher) *registerRepository {
	return &registerRepository{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
	}
}

func (repo *registerRepository) Register(ctx context.Context, data *auth_dto.RegisterUserRequest) error {
	user, err := repo.store.FindUserByEmail(ctx, data.Email)

	if user != nil {
		return shared.ErrEntityExisted(user_model.EntityName, err)
	}

	salt := shared.GenSalt(50)

	encryptedPassword := repo.hasher.Hash(data.Password + salt)

	userData := user_model.NewUser(data.Email, encryptedPassword, salt)

	if err := repo.store.Create(ctx, &userData); err != nil {
		return shared.ErrCannotCreateEntity(user_model.EntityName, err)
	}

	return nil
}
