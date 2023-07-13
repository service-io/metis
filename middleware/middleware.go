// Package middleware
// @author tabuyos
// @since 2023/6/30
// @description middleware
package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		parentContext := context.WithValue(ctx.Request.Context(), "user", "tabuyos-user")
		ctx.Request = ctx.Request.WithContext(parentContext)
		ctx.Next()
	}
}

func CheckRBAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
