// Package entity
// @author tabuyos
// @since 2023/6/30
// @description entity
package entity

import "time"

type CreateInfo struct {
	CreatorBy int64     `json:"creatorBy" db_cmt:"创建人员 ID"`
	CreatedAt time.Time `json:"createdAt" db_cmt:"创建时间"`
}

type UpdateInfo struct {
	UpdaterBy int64     `json:"updaterBy" db_cmt:"更新人员 ID"`
	UpdatedAt time.Time `json:"updatedAt" db_cmt:"更新时间"`
}

// Survey 问卷
type Survey struct {
	Id          int64     `json:"id" db_cmt:"主键"`
	Title       string    `json:"title" db_cmt:"主题"`
	Description string    `json:"description" db_cmt:"描述"`
	StartAt     time.Time `json:"startAt" db_cmt:"开始时间"`
	EndAt       time.Time `json:"endAt" db_cmt:"结束时间"`
	Status      int8      `json:"status" db_cmt:"0: 发布 1: 暂存 2: 已结束 3: 已失效"`
	Top         int8      `json:"top" db_cmt:"0: 不置顶 1: 置顶"`
	CreateInfo
	UpdateInfo
}

// Question 问题
type Question struct {
	Id          int64  `json:"id" db_cmt:"主键"`
	Title       string `json:"title" db_cmt:"问题主题"`
	Description string `json:"description" db_cmt:"问题描述"`
	Type        int8   `json:"type" db_cmt:"1: 单选 2: 多选 3: 填空"`
	Sort        int8   `json:"sort" db_cmt:"排序"`
	Required    int8   `json:"required" db_cmt:"0: 必填 1: 非必填"`
	Extra       string `json:"extra" db_cmt:"附加数据"`
	CreateInfo
	UpdateInfo
}

// Option 选项
type Option struct {
	Id         int64  `json:"id" db_cmt:"主键"`
	QuestionId int64  `json:"questionId" db_cmt:"问题 ID"`
	Status     int8   `json:"status" db_cmt:"0: 启用 1: 禁用"`
	Name       string `json:"name" db_cmt:"候选名"`
	Content    string `json:"content" db_cmt:"候选内容"`
	Answer     int8   `json:"answer" db_cmt:"是否是答案"`
	Extra      string `json:"extra" db_cmt:"附加数据"`
	CreateInfo
	UpdateInfo
}

// SurveyResult 问卷结果
type SurveyResult struct {
	Id       int64     `json:"id" db_cmt:"主键"`
	SurveyId int64     `json:"surveyId" db_cmt:"关联调查问卷主表 ID"`
	Ip       string    `json:"ip" db_cmt:"IP 地址"`
	Agent    string    `json:"agent" db_cmt:"代理信息"`
	Extra    string    `json:"extra" db_cmt:"附加数据"`
	Result   string    `json:"result" db_cmt:"结果数据"`
	AnswerAt time.Time `json:"answerAt" db_cmt:"回答时间"`
}

// SurveyTemplate 问卷模版
type SurveyTemplate struct {
	Id       int64  `json:"id" db_cmt:"主键"`
	SurveyId int64  `json:"surveyId" db_cmt:"关联调查问卷主表 ID"`
	Type     int8   `json:"type" db_cmt:"0: PC 1: MOBILE 2: EMBED"`
	Content  string `json:"content" db_cmt:"渲染模版"`
	CreateInfo
	UpdateInfo
}

// SurveyQuestionShip 问卷问题关系
type SurveyQuestionShip struct {
	SurveyId   int64 `json:"surveyId" db_cmt:"关联调查问卷主表 ID"`
	QuestionId int64 `json:"questionId" db_cmt:"关联问题主表 ID"`
}
