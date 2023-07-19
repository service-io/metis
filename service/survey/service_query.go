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
	surveyRepository := survey.New()
	return surveyRepository.SelectWithPage(info.Page, info.Size)
}
