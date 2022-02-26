package user_storage

import (
	"context"
	"relia_system/modules/user/user_model"
)

func (s *userMySqlStorage) FindUserByEmail(ctx context.Context, email string) (*user_model.User, error) {
	return s.FindUserByCondition(ctx, map[string]interface{}{
		"email": email,
	})
}
