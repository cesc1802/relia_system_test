package product_storage

import (
	"context"
	"relia_system/modules/product/product_dto"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
	"relia_system/shared/dbutil"
)

func (s *productMySqlStorage) List(ctx context.Context, paging *shared.Paging,
	filter *product_dto.ProductFilter, cond map[string]interface{}) ([]product_model.Product, error) {

	var products []product_model.Product

	db := s.db.Table(product_model.Product{}.TableName())

	db = db.Scopes(dbutil.QueryByStatus(1))

	if cond != nil {
		db = db.Where(cond)
	}

	if filter != nil {
		if filter.SKU != nil {
			db = db.Scopes(dbutil.QueryContaining("sku", *filter.SKU))
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, shared.ErrDB(err)
	}

	if err := db.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).
		Find(&products).Error; err != nil {
		return nil, shared.ErrDB(err)
	}

	return products, nil
}
