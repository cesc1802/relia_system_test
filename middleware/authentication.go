package middleware

import (
	"errors"
	"relia_system/app_context"
	"relia_system/modules/user/user_storage"
	"relia_system/shared"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("invalid auth header")
	}

	return parts[1], nil
}

func RequiredAuth(appCtx app_context.AppCtx) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenProvider := appCtx.GetTokenProvider()
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(shared.NewUnauthorized(err, "unauthorized", "unauthorized"))
			return
		}

		db := appCtx.GetDBConnection()
		store := user_storage.NewUserMysql(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(shared.NewUnauthorized(err, "unauthorized", "ErrUnAuthorize"))
		}

		user, err := store.FindUserById(c.Request.Context(), payload.GetUserId())

		if err != nil {
			panic(shared.NewUnauthorized(err, "user has been block", "ErrUnAuthorize"))
		}

		c.Set(shared.KeyCurrentUser, user)
		c.Next()
	}
}
