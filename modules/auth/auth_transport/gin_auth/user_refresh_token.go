package gin_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"relia_system/app_context"
	"relia_system/modules/auth/auth_model"
	"relia_system/modules/auth/auth_repository"
	"relia_system/modules/user/user_storage"
	"relia_system/shared"
)

func UserRefreshToken(appCtx app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rft auth_model.RefreshTokenRequest

		if err := c.ShouldBind(&rft); err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()

		store := user_storage.NewUserMysql(db)
		tokProvider := appCtx.GetTokenProvider()

		repo := auth_repository.NewRfTokenRepository(store, tokProvider)
		result, err := repo.RefreshToken(c.Request.Context(), &rft)
		if err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, shared.SimpleSuccessResponse(result))
		return
	}
}
