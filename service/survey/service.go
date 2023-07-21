// Package survey
// @author tabuyos
// @since 2023/7/18
// @description survey
package survey

import (
	"github.com/gin-gonic/gin"
	"metis/model/dto"
	"metis/model/page"
)

type Service interface {
	FindAll() []dto.Survey
	FindAllWithPage(pi *page.Info) []dto.Survey
	FindById(id int64) dto.Survey
}
type service struct {
	ctx *gin.Context
}

func New(ctx *gin.Context) Service {
	return &service{ctx}
}
