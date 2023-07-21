// Package survey
// @author tabuyos
// @since 2023/7/18
// @description survey
package survey

import (
	"metis/model/dto"
	"metis/model/page"
	"metis/repository/survey"
)

func (receiver *service) FindAll() []dto.Survey {
	panic("implement me")
}

func (receiver *service) FindAllWithPage(info *page.Info) []dto.Survey {
	surveyRepository := survey.New(receiver.ctx)
	return surveyRepository.SelectWithPage(info.Page, info.Size)
}

func (receiver *service) FindById(id int64) dto.Survey {
	surveyRepository := survey.New(receiver.ctx)
	return surveyRepository.SelectById(id)
}
