package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"relia_system/app_context"
	"relia_system/middleware"
	"relia_system/modules/auth/auth_transport/gin_auth"
	"relia_system/modules/product/product_transport/gin_product"
)

func SetupRouter(engine *gin.Engine, appCtx app_context.AppCtx) {

	engine.Use(middleware.Recover(appCtx), middleware.Cors())

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := engine.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", gin_auth.UserLogin(appCtx))
		auth.POST("/register", gin_auth.UserRegister(appCtx))
		auth.POST("/refresh_token", gin_auth.UserRefreshToken(appCtx))
	}

	product := v1.Group("/item", middleware.RequiredAuth(appCtx))
	{
		product.GET("/search", gin_product.ListProduct(appCtx))
		product.POST("/add", gin_product.CreateProduct(appCtx))
		product.POST("/update", gin_product.UpdateProduct(appCtx))
		product.POST("/delete", gin_product.DeleteProduct(appCtx))
	}

	searchProduct := v1.Group("/items", middleware.RequiredAuth(appCtx))
	{
		searchProduct.GET("", gin_product.ListProduct(appCtx))
	}
}
