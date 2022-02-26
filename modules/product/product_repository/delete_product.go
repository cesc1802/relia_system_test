package product_repository

import (
	"context"
	"gorm.io/gorm"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

type DeleteProductStorage interface {
	FindProductBySKU(ctx context.Context,
		sku string) (*product_model.Product, error)
	DeleteBySKU(ctx context.Context, sku string) error
}

type deleteProductRepository struct {
	store DeleteProductStorage
}

func NewDeleteProductRepository(store DeleteProductStorage) *deleteProductRepository {
	return &deleteProductRepository{
		store: store,
	}
}
func (repo *deleteProductRepository) DeleteProduct(ctx context.Context, sku string) error {
	_, err := repo.store.FindProductBySKU(ctx, sku)

	if err == gorm.ErrRecordNotFound {
		return product_model.ErrProductNotFound(err)
	}

	if err := repo.store.DeleteBySKU(ctx, sku); err != nil {
		return shared.ErrCannotDeleteEntity(product_model.EntityName, err)
	}

	return nil
}
