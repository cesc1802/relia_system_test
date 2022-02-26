package product_storage

import (
	"golang.org/x/net/context"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

func (s *productMySqlStorage) Create(ctx context.Context, data *product_model.Product) error {
	db := s.db.Begin()

	db = db.Table(product_model.Product{}.TableName())

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
