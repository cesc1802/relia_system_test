package middleware

import (
	"github.com/gin-gonic/gin"
	"relia_system/app_context"
	"relia_system/shared"
)

func Recover(appCtx app_context.AppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json; charset=UTF-8")

				if appErr, ok := err.(*shared.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}

				appErr := shared.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			}
		}()
		c.Next()
	}
}
