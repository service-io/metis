package entity

import "time"

type Role struct {
	ID      *int64
	Name    *string
	Title   *string
	Left    *int
	Right   *int
	Level   *int
	TreeNo  *int
	StartAt *time.Time
}

// FromDto
// IntoDto

func (rec Role) PKey() string {
	return "id"
}

func (rec Role) NameKey() string {
	return "name"
}
