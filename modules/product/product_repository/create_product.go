package product_repository

import (
	"context"
	"relia_system/modules/product/product_dto"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

type CreateProductStorage interface {
	FindProductBySKU(ctx context.Context,
		sku string) (*product_model.Product, error)
	Create(ctx context.Context, data *product_model.Product) error
}

type createProductRepository struct {
	store CreateProductStorage
}

func NewCreateProductRepository(store CreateProductStorage) *createProductRepository {
	return &createProductRepository{
		store: store,
	}
}

func (repo *createProductRepository) CreateProduct(ctx context.Context, input *product_dto.CreateProductRequest) error {

	product, err := repo.store.FindProductBySKU(ctx, input.SKU)

	if product != nil {
		return shared.ErrEntityExisted(product_model.EntityName, err)
	}
	productData := product_model.NewProduct(input.SKU, input.ProductName,
		input.Quantity, input.Price,
		input.Unit, input.Status)

	if err := repo.store.Create(ctx, &productData); err != nil {
		return shared.ErrCannotCreateEntity(product_model.EntityName, err)
	}
	return nil
}
