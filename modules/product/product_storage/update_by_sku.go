package product_storage

import "context"

func (s *productMySqlStorage) UpdateBySKU(ctx context.Context, sku string, data map[string]interface{}) error {
	condition := map[string]interface{}{
		"sku": sku,
	}

	return s.UpdateByCondition(ctx, condition, data)
}
