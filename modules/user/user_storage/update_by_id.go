package user_storage

import (
	"context"
	"relia_system/modules/user/user_model"
)

func (s *userMySqlStorage) UpdateById(ctx context.Context,
	userId int, data map[string]interface{}) error {

	db := s.db.Begin()
	db = db.Table(user_model.User{}.TableName())

	condition := map[string]interface{}{
		"id": userId,
	}

	return s.UpdateByCondition(ctx, condition, data)
}
