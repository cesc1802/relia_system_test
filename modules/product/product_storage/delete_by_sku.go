package product_storage

import (
	"context"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

func (s *productMySqlStorage) DeleteBySKU(ctx context.Context, sku string) error {
	db := s.db

	db = db.Table(product_model.Product{}.TableName())

	if err := db.Where("sku", sku).Delete(nil).Error; err != nil {
		return shared.ErrDB(err)
	}

	return nil
}
