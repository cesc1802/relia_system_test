package product_dto

import "relia_system/shared"

type ProductFilter struct {
	SKU *string `json:"sku" form:"sku"`
}

type ProductListParam struct {
	*ProductFilter `json:",inline"`
	shared.Paging  `json:",inline"`
}
