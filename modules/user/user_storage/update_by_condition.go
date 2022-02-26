package user_storage

import (
	"golang.org/x/net/context"
	"relia_system/modules/user/user_model"
	"relia_system/shared"
)

func (s *userMySqlStorage) UpdateByCondition(ctx context.Context,
	conditions map[string]interface{}, data map[string]interface{}) error {

	db := s.db.Begin()
	db = db.Table(user_model.User{}.TableName())

	if err := db.Where(conditions).
		Updates(&data).Error; err != nil {
		db.Rollback()
		return shared.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return shared.ErrDB(err)
	}

	return nil
}
