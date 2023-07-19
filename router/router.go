// Package router
// @author tabuyos
// @since 2023/6/30
// @description router
package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"metis/util/logger"
	"net/http"
	"strings"
	"time"
)

var baseRouter *gin.Engine

func init() {
	baseRouter = gin.New()
	baseRouter.Use(fillTrace(), loggerFunc(), recoveryFunc())
	baseRouter.NoRoute(func(ctx *gin.Context) {
		if ctx.Error(errors.New("未找到指定 API")) != nil {
		}
		ctx.String(http.StatusNotFound, "404 NOT FOUND!")
	})

	setApiRouter()
}

func loggerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		useLogger := logger.UseLogger()
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		// hand over to the next handler
		ctx.Next()
		latency := time.Now().Sub(start)
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		// ems := ctx.Errors.ByType(gin.ErrorTypePrivate)
		// if len(ems) > 0 {
		//
		// }
		// var errorStrings []string
		// for _, em := range ems {
		// 	errorStrings = append(errorStrings, em.Error())
		// }
		// errorMessage := strings.Join(errorStrings, ", ")
		errorMessage := strings.TrimRight(ctx.Errors.ByType(gin.ErrorTypePrivate).String(), "\n")
		msg := fmt.Sprintf("%v -> %v(%v)[%v] -> %v -> |%v|%v", clientIP, path, raw, method, latency, statusCode, errorMessage)
		useLogger.Info(msg)
	}
}

func recoveryFunc() gin.HandlerFunc {
	return gin.Recovery()
}

func fillTrace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := uuid.New().String()
		logCtx := context.WithValue(context.Background(), "traceId", traceId)
		logger.WithCtx(logCtx)
		ctx.Next()
		logger.CleanCtx()
	}
}

func BaseRouter() *gin.Engine {
	return baseRouter
}
