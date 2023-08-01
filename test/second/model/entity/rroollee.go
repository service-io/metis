package entity

import "time"

type Rroollee struct {
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

func (rec Rroollee) PKey() string {
	return "id"
}

func (rec Rroollee) NameKey() string {
	return "name"
}
