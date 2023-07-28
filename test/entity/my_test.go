// Package entity
// @author tabuyos
// @since 2023/7/28
// @description entity
package entity

import (
	"fmt"
	// util tool
	mu "metis/util"
	"unicode"
)

// Tabuyos wo de
type Tabuyos struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
}

// Toi toi func
func (rec *Tabuyos) Toi(name string) bool {
	fmt.Println(name)
	if len(name) == 5 {
		// print something...
		println(5)
		fmt.Println(name)
	}
	mu.SplitFunc(name, func(r rune) bool {
		return unicode.IsUpper(r)
	})
	// ret is lt 3
	return len(name) > 3
}
