package middleware

import (
	"fenny-web/framework/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				ctx.JSON(500, p)
			}
		}()
		ctx.Next()
	}
}
