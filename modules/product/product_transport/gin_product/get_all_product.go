package gin_product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"relia_system/app_context"
	"relia_system/modules/product/product_dto"
	"relia_system/modules/product/product_repository"
	"relia_system/modules/product/product_storage"
	"relia_system/shared"
)

func ListProduct(appCtx app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param product_dto.ProductListParam

		if err := c.ShouldBind(&param); err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		param.Fulfill()

		db := appCtx.GetDBConnection()
		store := product_storage.NewProductMysql(db)
		repo := product_repository.NewListProductRepository(store)

		products, err := repo.ListProduct(c.Request.Context(), &param.Paging, param.ProductFilter)
		if err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		for i := 0; i < len(products); i++ {
			products[i].Mask(false)
		}

		c.JSON(http.StatusOK, shared.NewSuccessResponse(products, param.Paging, param.ProductFilter))
		return
	}
}
