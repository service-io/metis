// Package router
// @author tabuyos
// @since 2023/7/10
// @description router
package router

import (
	"github.com/gin-gonic/gin"
	"metis/middleware"
	"metis/util/logger"
)

func setApiRouter() {
	recorder := logger.UseLogger()
	root := baseRouter.Group("/")
	root.Use(middleware.CheckAuth(), middleware.CheckRBAC())
	root.GET("/test",
		func(ctx *gin.Context) {
			user := ctx.Request.Context().Value("user").(string)
			println(user)
			recorder.Info("into first handler function...")
			ctx.Abort()
		},
		func(c *gin.Context) {
			recorder.Info("into second handler function...")
			c.JSON(200, gin.H{
				"message": "test",
			})
		},
		func(ctx *gin.Context) {
			recorder.Info("into third handler function...")
		})
}
