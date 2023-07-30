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

func HandleRollback(err error, tx *sql.Tx, eh func(err error)) {
	if err != nil {
		err := tx.Rollback()
		eh(err)
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

func PanicErr(logger *zap.Logger, err error) {
	if err != nil {
		logger.Error(err.Error())
		panic(err)
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

// SplitFunc 使用函数进行分割, 注意: 并不会移除符合谓词的字符,
// 具体实现参考 strings.FieldsFunc 进行修改的,
// strings.FieldsFunc 会移除符合谓词的字符
func SplitFunc(s string, f func(rune) bool) []string {
	type span struct {
		start int
		end   int
	}
	spans := make([]span, 0, 32)

	start := -1
	for end, char := range s {
		if f(char) {
			if start == -1 {
				start = 0
			} else if start >= 0 {
				spans = append(spans, span{start, end})
				start = end
			}
		}
	}

	if start >= 0 {
		spans = append(spans, span{start, len(s)})
	} else {
		spans = append(spans, span{0, len(s)})
	}

	a := make([]string, len(spans))
	for i, span := range spans {
		a[i] = s[span.start:span.end]
	}

	return a
}

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
