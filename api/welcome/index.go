// Package welcome
// @author tabuyos
// @since 2023/7/18
// @description api
package welcome

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"metis/util/logger"
	"net/http"
	"strconv"
	"time"
)

// Hello 你好
// https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @Summary ping example
// @Description 打印 Hello 信息
// @Tags welcome
// @Produce plain
// @Success 200 {string} string "HELLO WORLD, HELLO GOLANG, WELCOME TO METIS"
// @Router /welcome/hello [get]
func Hello(ctx *gin.Context) {
	recorder := logger.UseLogger()
	msg := "HELLO WORLD, HELLO GOLANG, WELCOME TO METIS"
	recorder.Info(msg)
	ctx.String(http.StatusOK, msg)
}

func WhoAmI(ctx *gin.Context) {
	recorder := logger.UseLogger()
	// 1: in go
	// user := ctx.Request.Context().Value("user").(string)

	// 2: in gin
	user := ctx.GetString("user")
	recorder.Info("current user ->", zap.String("user", user))

	randomInt := rand.Intn(10)
	fmt.Println("sleep -> " + strconv.Itoa(randomInt))
	time.Sleep(time.Duration(randomInt) * time.Second)
}
