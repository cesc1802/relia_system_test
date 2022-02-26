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

func CreateProduct(appCtx app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data product_dto.CreateProductRequest

		if err := c.ShouldBind(&data); err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := product_storage.NewProductMysql(db)
		repo := product_repository.NewCreateProductRepository(store)

		if err := repo.CreateProduct(c.Request.Context(), &data); err != nil {
			panic(shared.ErrInvalidRequest(err))
		}
		c.JSON(http.StatusOK, shared.SimpleSuccessResponse(true))
		return
	}
}
