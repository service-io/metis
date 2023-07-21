// Package util
// @author tabuyos
// @since 2023/7/20
// @description util
package util

import (
	"database/sql"
	"go.uber.org/zap"
	"io"
)

func DeferClose(closer io.Closer, errHandler ...func(err error)) {
	err := closer.Close()
	if err != nil {
		if len(errHandler) > 0 {
			for _, eh := range errHandler {
				eh(err)
			}
		}
		return
	}
}

func ErrToLog(logger *zap.Logger) func(err error) {
	return func(err error) {
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func LogErr(logger *zap.Logger, err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}

func Rows[T any](rows *sql.Rows, supplier func() (*T, []any)) []T {
	rs := make([]T, 0)
	for rows.Next() {
		r, cs := supplier()
		if err := rows.Scan(cs...); err != nil {
			panic(err)
		}
		rs = append(rs, *r)
	}
	return rs
}

func Row[T any](row *sql.Row, supplier func() (*T, []any)) T {
	r, cs := supplier()
	if err := row.Scan(cs...); err != nil {
		panic(err)
	}
	return *r
}
