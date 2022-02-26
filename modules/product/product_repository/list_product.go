package product_repository

import (
	"context"
	"relia_system/modules/product/product_dto"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
)

type ListProductStorage interface {
	List(ctx context.Context, paging *shared.Paging,
		filter *product_dto.ProductFilter, cond map[string]interface{}) ([]product_model.Product, error)
}

type listProductRepository struct {
	store ListProductStorage
}

func NewListProductRepository(store ListProductStorage) *listProductRepository {
	return &listProductRepository{
		store: store,
	}
}

func (repo *listProductRepository) ListProduct(ctx context.Context,
	paging *shared.Paging, filter *product_dto.ProductFilter) ([]product_model.Product, error) {
	return repo.store.List(ctx, paging, filter, nil)
}
