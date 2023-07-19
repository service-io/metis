// Package page
// @author tabuyos
// @since 2023/7/18
// @description model
package page

import (
	"bytes"
	"strconv"
)

type Info struct {
	Page uint `json:"page"`
	Size uint `json:"size"`
}

func New() *Info {
	return &Info{}
}

func (receiver Info) String() string {
	var b bytes.Buffer
	b.WriteString("(")
	b.WriteString("page -> ")
	b.WriteString(strconv.FormatUint(uint64(receiver.Page), 10))
	b.WriteString(", ")
	b.WriteString("size -> ")
	b.WriteString(strconv.FormatUint(uint64(receiver.Size), 10))
	b.WriteString(")")
	return b.String()
}
