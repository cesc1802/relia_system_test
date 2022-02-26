package product_dto

type CreateProductRequest struct {
	SKU         string  `json:"sku" form:"sku" binding:"required"`
	ProductName string  `json:"product_name" form:"product_name" binding:"required"`
	Quantity    uint64  `json:"quantity" form:"quantity" binding:"required"`
	Price       float64 `json:"price" form:"price" binding:"required"`
	Unit        string  `json:"unit" form:"unit" binding:"required"`
	Status      uint32  `json:"status" form:"status" binding:"required"`
}
