package user_storage

import (
	"context"
	"relia_system/modules/user/user_model"
)

func (s *userMySqlStorage) FindUserById(ctx context.Context, id int) (*user_model.User, error) {
	return s.FindUserByCondition(ctx, map[string]interface{}{
		"id": id,
	})
}
