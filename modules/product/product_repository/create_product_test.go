package product_repository

import (
	"context"
	"relia_system/modules/product/product_dto"
	"relia_system/modules/product/product_model"
	"relia_system/shared"
	"testing"
)

type mockCreateProductStorage struct {
	findProductBySKU func(ctx context.Context, sku string) (*product_model.Product, error)
	create           func(ctx context.Context, data *product_model.Product) error
}

func (m *mockCreateProductStorage) FindProductBySKU(ctx context.Context, sku string) (*product_model.Product, error) {
	if m != nil && m.findProductBySKU != nil {
		return m.findProductBySKU(ctx, sku)
	}
	return &product_model.Product{}, nil
}

func (m *mockCreateProductStorage) Create(ctx context.Context, data *product_model.Product) error {
	if m != nil && m.create != nil {
		return m.create(ctx, data)
	}
	return nil
}

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		testName         string
		sku              string
		productName      string
		quantity         uint64
		price            float64
		unit             string
		status           uint32
		mockProductStore *mockCreateProductStorage
		expectedError    func(err error) *shared.AppError
	}{
		{
			testName:    "Test case SKU exist in system",
			sku:         "APPLE-101",
			productName: "IP13 PRO MAX",
			quantity:    12,
			price:       1200000,
			unit:        "CAI",
			status:      1,
			mockProductStore: &mockCreateProductStorage{
				findProductBySKU: func(ctx context.Context, sku string) (*product_model.Product, error) {
					product := product_model.NewProduct("APPLE-101", "IP13 Pro MAX",
						10, 120000,
						"CAI", 1)
					return &product, nil
				},

				create: func(ctx context.Context, data *product_model.Product) error {
					return nil
				},
			},
			expectedError: func(err error) *shared.AppError {
				return shared.ErrEntityExisted(product_model.EntityName, err)
			},
		},
	}
	for _, tt := range tests {
		repo := NewCreateProductRepository(tt.mockProductStore)

		input := product_dto.CreateProductRequest{
			SKU:         tt.sku,
			ProductName: tt.productName,
			Quantity:    tt.quantity,
			Unit:        tt.unit,
			Price:       tt.price,
			Status:      tt.status,
		}
		err := repo.CreateProduct(context.TODO(), &input)
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}

		if tt.expectedError(err).Log != errMsg {
			t.Errorf("Unexpected error: %v", errMsg)
		}
	}
}
