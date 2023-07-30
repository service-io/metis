package account

import (
	"github.com/jinzhu/copier"
	"metis/test/first/entity"
	"time"
)

type Account struct {
	*entity.IDInfo
	Name    *string
	Title   *string
	StartAt *time.Time
	entity.CreateInfo
	entity.ModifyInfo
	entity.DeletedInfo
}

func NewAccount() *Account {
	return &Account{}
}

func (rec *Account) From(from any) error {
	err := copier.Copy(rec, from)
	return err
}

func (rec *Account) Into(into *any) error {
	err := copier.Copy(into, rec)
	return err
}
