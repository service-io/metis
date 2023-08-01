// Package generated
// @author tabuyos
// @since 2023/8/1
// @description generated
package generated

import (
	"github.com/dave/jennifer/jen"
	"metis/generated/repository"
)

func RenderFile(table string) *jen.File {
	columns := getColumns(table)
	r := &repository.GenBaseRepo{}
	return r.GenFile(table, columns)
}

func getColumns(table string) []Column {
	var columns []Column

	columns = append(
		columns, Column{
			ColumnName:      "id",
			Type:            "bigint",
			Nullable:        "NO",
			TableName:       table,
			ColumnComment:   "主键",
			MaxLength:       0,
			NumberPrecision: 19,
			ColumnType:      "bigint",
			ColumnKey:       "PRI",
			Default:         "",
		},
	)
	columns = append(
		columns, Column{
			ColumnName:      "title",
			Type:            "varchar",
			Nullable:        "NO",
			TableName:       table,
			ColumnComment:   "主题",
			MaxLength:       255,
			NumberPrecision: 0,
			ColumnType:      "varchar(1000)",
			ColumnKey:       "",
			Default:         "",
		},
	)
	columns = append(
		columns, Column{
			ColumnName:      "start_at",
			Type:            "timestamp",
			Nullable:        "NO",
			TableName:       table,
			ColumnComment:   "开始时间",
			MaxLength:       0,
			NumberPrecision: 0,
			ColumnType:      "timestamp",
			ColumnKey:       "",
			Default:         "CURRENT_TIMESTAMP",
		},
	)
	columns = append(
		columns, Column{
			ColumnName:      "ns_id",
			Type:            "bigint",
			Nullable:        "NO",
			TableName:       table,
			ColumnComment:   "开始时间",
			MaxLength:       0,
			NumberPrecision: 0,
			ColumnType:      "bigint",
			ColumnKey:       "",
			Default:         "CURRENT_TIMESTAMP",
		},
	)

	// columns := make([]Column, 0)
	// var params = make(map[string]string)
	// params["parseTime"] = "true"
	// cfg := mysql.Config{
	//	User:   "root",
	//	Passwd: "root",
	//	Net:    "tcp",
	//	Addr:   "localhost:3307",
	//	DBName: "metis",
	//	Params: params,
	// }
	//
	// var err error
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// rows, err := db.Query(
	//	fmt.Sprintf(
	//		`SELECT
	//	COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT,CHARACTER_MAXIMUM_LENGTH,COLUMN_TYPE,NUMERIC_PRECISION,COLUMN_KEY,COLUMN_DEFAULT
	//	FROM information_schema.COLUMNS
	//	WHERE table_schema = DATABASE()  AND TABLE_NAME = '%s'`, table,
	//	),
	// )
	//
	// if err != nil {
	//	log.Printf("table rows is nil with table:%s error: %v \n", table, err)
	//	return columns
	// }
	//
	// if rows == nil {
	//	log.Printf("rows is nil with table:%s \n", table)
	//	return columns
	// }
	//
	// defer func() {
	//	_ = rows.Close()
	// }()
	//
	// for rows.Next() {
	//
	//	// todo: mysql bigint => go []byte
	//	var maxLength, numberPrecision []byte
	//	var t = ""
	//
	//	col := Column{}
	//	err = rows.Scan(
	//		&col.ColumnName, &col.Type, &t, &col.TableName, &col.ColumnComment, &maxLength, &col.ColumnType, &numberPrecision,
	//		&col.ColumnKey, &col.Default,
	//	)
	//	col.Nullable = t
	//	col.Tag = col.ColumnName
	//
	//	if maxLength != nil {
	//		col.MaxLength = Byte2Int64(maxLength)
	//	}
	//
	//	if numberPrecision != nil {
	//		col.NumberPrecision = Byte2Int64(numberPrecision)
	//	}
	//
	//	if err != nil {
	//		log.Println(err.Error())
	//		continue
	//	}
	//
	//	columns = append(columns, col)
	// }

	return columns
}
