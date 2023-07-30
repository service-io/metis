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

type TreeInfo struct {
	Left   *int `json:"left"`
	Right  *int `json:"right"`
	Level  *int `json:"level"`
	TreeNo *int `json:"treeNo"`
}
