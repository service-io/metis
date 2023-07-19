// Package survey
// @author tabuyos
// @since 2023/7/19
// @description survey
package survey

import (
	"database/sql"
	"go.uber.org/zap"
	"metis/database"
	"metis/model/dto"
	"metis/util/logger"
)

type Repository interface {
	SelectWithPage(page uint, size uint) []dto.Survey
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (receiver *repository) SelectWithPage(page uint, size uint) []dto.Survey {
	var surveys []dto.Survey
	useLogger := logger.UseLogger()
	db := database.FetchDB()
	prepare, _ := db.Prepare("select id, title, status, start_at from survey limit ? offset ?")
	rows, err := prepare.Query(size, (page-1)*size)
	if err != nil {
		useLogger.Error(err.Error(), zap.Error(err))
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			useLogger.Error(err.Error(), zap.Error(err))
		}
	}(rows)

	for rows.Next() {
		var survey dto.Survey
		if err := rows.Scan(&survey.Id, &survey.Title, &survey.Status, &survey.StartAt); err != nil {
			useLogger.Error(err.Error(), zap.Error(err))
			return nil
		}
		surveys = append(surveys, survey)
	}
	return surveys
}
