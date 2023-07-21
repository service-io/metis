// Package survey
// @author tabuyos
// @since 2023/7/18
// @description survey
package survey

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metis/model/page"
	"metis/model/reply"
	"metis/service/survey"
	"metis/util"
	"metis/util/logger"
	"net/http"
	"strconv"
)

// List 列表查询
func List(ctx *gin.Context) {
	pageInfo := page.New()
	surveyService := survey.New(ctx)
	accessLogger := logger.AccessLogger(ctx)

	err := ctx.ShouldBindJSON(pageInfo)
	if err != nil {
		accessLogger.Error(err.Error(), zap.Error(err))
		return
	}

	accessLogger.Info(pageInfo.String())
	surveys := surveyService.FindAllWithPage(pageInfo)

	ctx.JSON(http.StatusOK, reply.OkData(surveys))
}

// Detail 详情
func Detail(ctx *gin.Context) {
	surveyService := survey.New(ctx)
	accessLogger := logger.AccessLogger(ctx)

	var id int64

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	util.LogErr(accessLogger, err)

	accessLogger.Sugar().Infof("查询 %d 的详情", id)
	surveyEnt := surveyService.FindById(id)

	ctx.JSON(http.StatusOK, reply.OkData(surveyEnt))
}

// Create 新增
func Create(ctx *gin.Context) {}

// Modify 修改
func Modify(ctx *gin.Context) {}

// Delete 删除
func Delete(ctx *gin.Context) {}
