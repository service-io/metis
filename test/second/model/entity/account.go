// Code generated by tabuyos. DO NOT EDIT!

// Package entity
// @author tabuyos
// @since 2023/08/01
// @description account
package entity

import (
	"time"
)

type Account struct {
	*IDInfo
	Title   *string
	StartAt *time.Time
	NsId    *int64
}

func NewAccount() *Account {
	return &Account{}
}