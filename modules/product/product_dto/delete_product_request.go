package product_dto

type DeleteProductRequest struct {
	SKU string `json:"sku" form:"sku" binding:"required"`
}
