package entity

import "time"

type Account struct {
	ID      *int64
	Name    *string
	Title   *string
	StartAt *time.Time
}
