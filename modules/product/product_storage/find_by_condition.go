package product_storage

import (
	"context"
	"gorm.io/gorm"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

func (s *productMySqlStorage) FindByCondition(ctx context.Context,
	conditions map[string]interface{}) (*product_model.Product, error) {
	var data product_model.Product

	db := s.db.Table(product_model.Product{}.TableName())

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, shared.ErrDB(err)
		}
		return nil, err
	}
	return &data, nil
}
