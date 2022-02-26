package product_repository

import (
	"context"
	"gorm.io/gorm"
	"relia_system/modules/product/product_dto"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

type UpdateProductStorage interface {
	FindProductBySKU(ctx context.Context,
		sku string) (*product_model.Product, error)
	UpdateBySKU(ctx context.Context, sku string, data map[string]interface{}) error
}

type updateProductRepository struct {
	store UpdateProductStorage
}

func NewUpdateProductRepository(store UpdateProductStorage) *updateProductRepository {
	return &updateProductRepository{
		store: store,
	}
}

func newUpdateData(input *product_dto.UpdateProductRequest) map[string]interface{} {
	var updateData = make(map[string]interface{})

	if input.SKU != nil {
		updateData["sku"] = input.SKU
	}
	if input.ProductName != nil {
		updateData["product_name"] = input.ProductName
	}
	if input.Quantity != nil {
		updateData["quantity"] = input.Quantity
	}

	if input.Price != nil {
		updateData["price"] = input.Price
	}

	if input.Unit != nil {
		updateData["unit"] = input.Unit
	}

	if input.Status != nil {
		updateData["status"] = input.Status
	}
	return updateData
}

func (repo *updateProductRepository) UpdateProduct(ctx context.Context, input *product_dto.UpdateProductRequest) error {
	_, err := repo.store.FindProductBySKU(ctx, *input.SKU)

	if err == gorm.ErrRecordNotFound {
		return product_model.ErrProductNotFound(err)
	}

	updateData := newUpdateData(input)
	if err := repo.store.UpdateBySKU(ctx, *input.SKU, updateData); err != nil {
		return shared.ErrCannotUpdateEntity(product_model.EntityName, err)
	}

	return nil
}
