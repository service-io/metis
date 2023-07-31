package entity

import (
	"time"
)

type IDInfo struct {
	ID *int64 `json:"id"`
}

type DeletedInfo struct {
	Deleted *int64 `json:"deleted"`
}

type CreateInfo struct {
	CreateBy *int64     `json:"createBy"`
	CreateAt *time.Time `json:"createAt"`
}

type ModifyInfo struct {
	ModifyBy *int64     `json:"createBy"`
	ModifyAt *time.Time `json:"createAt"`
}

type IEntity interface {
	// PKey 主键
	PKey() string
	// NameKey 名称
	NameKey() string
	// LDKey 逻辑删除
	LDKey() string
}

type ITreeEntity interface {
	IEntity
	// LKey 左值
	LKey() string
	// RKey 右值
	RKey() string
	// LlKey 层级
	LlKey() string
	// TNKey 树号
	TNKey() string
}
