package product_model

import "relia_system/shared"

const (
	EntityName = "Product"
)

type Product struct {
	shared.SQLModel
	SKU         string  `gorm:"sku"`
	ProductName string  `gorm:"product_name"`
	Quantity    uint64  `gorm:"quantity"`
	Price       float64 `gorm:"price"`
	Unit        string  `gorm:"unit"`
	Status      uint32  `gorm:"status"`
}

func (Product) TableName() string {
	return "products"
}

func NewProduct(sku, productName string, qty uint64, price float64, unit string, status uint32) Product {
	return Product{
		SKU:         sku,
		ProductName: productName,
		Quantity:    qty,
		Price:       price,
		Unit:        unit,
		Status:      status,
	}
}

func ErrProductNotFound(err error) *shared.AppError {
	return shared.NewCustomError(err, "product not found", "ERR_PRODUCT_NOT_FOUND")
}

func (p *Product) Mask(isAdmin bool) {
	p.GenUID(shared.DBTypeProduct, 1)
}
