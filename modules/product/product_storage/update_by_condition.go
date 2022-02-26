package product_storage

import (
	"context"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

func (s *productMySqlStorage) UpdateByCondition(ctx context.Context,
	conditions map[string]interface{}, data map[string]interface{}) error {

	db := s.db.Begin()
	db = db.Table(product_model.Product{}.TableName())

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
