package product_storage

import (
	"context"
	"relia_system/modules/product/product_model"
)

func (s *productMySqlStorage) FindProductBySKU(ctx context.Context,
	sku string) (*product_model.Product, error) {
	return s.FindByCondition(ctx, map[string]interface{}{
		"sku": sku,
	})
}
