package user_storage

import (
	"context"
	"gorm.io/gorm"
	"relia_system/modules/user/user_model"
	"relia_system/shared"
)

func (s *userMySqlStorage) FindUserByCondition(ctx context.Context,
	conditions map[string]interface{}) (*user_model.User, error) {
	var data user_model.User

	db := s.db.Table(user_model.User{}.TableName())

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, shared.ErrDB(err)
		}
		return nil, err
	}
	return &data, nil
}
