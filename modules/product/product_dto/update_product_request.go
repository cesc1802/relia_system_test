package product_dto

type UpdateProductRequest struct {
	SKU         *string  `json:"sku" form:"sku"`
	ProductName *string  `json:"product_name" form:"product_name"`
	Quantity    *uint64  `json:"quantity" form:"quantity"`
	Price       *float64 `json:"price" form:"price"`
	Unit        *string  `json:"unit" form:"unit"`
	Status      *uint32  `json:"status" form:"status"`
}
