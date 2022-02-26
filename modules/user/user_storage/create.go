package user_storage

import (
	"golang.org/x/net/context"
	"relia_system/modules/user/user_model"
	"relia_system/shared"
)

func (s *userMySqlStorage) Create(ctx context.Context, data *user_model.User) error {
	db := s.db.Begin()

	db = db.Table(user_model.User{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		db.Rollback()
		return shared.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return shared.ErrDB(err)
	}

	return nil
}
