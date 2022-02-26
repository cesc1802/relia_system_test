package gin_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"relia_system/app_context"
	"relia_system/hash"
	"relia_system/modules/auth/auth_dto"
	"relia_system/modules/auth/auth_repository"
	"relia_system/modules/user/user_storage"
	"relia_system/shared"
)

func UserRegister(appCtx app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data auth_dto.RegisterUserRequest

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, shared.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()
		tokProvider := appCtx.GetTokenProvider()
		md5 := hash.NewMd5Hash()
		store := user_storage.NewUserMysql(db)
		repo := auth_repository.NewRegisterRepository(store, tokProvider, md5)

		err := repo.Register(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, shared.ErrInvalidRequest(err))
			return
		}

		c.JSON(http.StatusOK, shared.SimpleSuccessResponse(true))
		return
	}
}
