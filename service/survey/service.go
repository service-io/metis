// Package survey
// @author tabuyos
// @since 2023/7/18
// @description survey
package survey

import (
	"metis/model/dto"
	"metis/model/page"
)

type Service interface {
	FindAll() []dto.Survey
	FindAllWithPage(pi *page.Info) []dto.Survey
}
type service struct{}

func New() Service {
	return &service{}
}
