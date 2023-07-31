// Package survey
// @author tabuyos
// @since 2023/7/19
// @description survey
package survey

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metis/config/constant"
	"metis/database"
	"metis/model/dto"
	"metis/util"
	"metis/util/logger"
)

type Repository interface {
	SelectWithPage(page uint, size uint) []dto.Survey
	SelectById(id int64) dto.Survey
}

type repository struct {
	ctx *gin.Context
}

func New(ctx *gin.Context) Repository {
	return &repository{ctx}
}

func (receiver *repository) getDbCtx() context.Context {
	return context.WithValue(context.Background(), constant.TraceIdKey, receiver.ctx.GetString(constant.TraceIdKey))
}

func (receiver *repository) SelectWithPage(page uint, size uint) []dto.Survey {
	accessLogger := logger.AccessLogger(receiver.ctx)
	db := database.FetchDB()

	prepare, _ := db.Prepare("select id, title, status, start_at from survey limit ? offset ?")
	defer util.DeferClose(prepare, util.ErrToLogAndPanic(accessLogger))

	rows, err := prepare.QueryContext(receiver.getDbCtx(), size, (page-1)*size)
	defer util.DeferClose(rows, util.ErrToLogAndPanic(accessLogger))

	if err != nil {
		accessLogger.Error(err.Error(), zap.Error(err))
	}

	surveys := util.Rows[dto.Survey](rows, func() (*dto.Survey, []any) {
		var r = &dto.Survey{}
		var cs = []any{&r.Id, &r.Title, &r.Status, &r.StartAt}
		return r, cs
	})
	return surveys
}

func (receiver *repository) SelectById(id int64) dto.Survey {
	accessLogger := logger.AccessLogger(receiver.ctx)
	db := database.FetchDB()

	prepare, _ := db.Prepare("select id, title, status, start_at from survey where id = ?")
	defer util.DeferClose(prepare, util.ErrToLogAndPanic(accessLogger))

	row := prepare.QueryRowContext(receiver.getDbCtx(), id)
	survey := util.Row(row, func() (*dto.Survey, []any) {
		var r = &dto.Survey{}
		var cs = []any{&r.Id, &r.Title, &r.Status, &r.StartAt}
		return r, cs
	})

	return survey
}
