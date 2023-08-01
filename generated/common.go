// Package generated
// @author tabuyos
// @since 2023/8/1
// @description generated
package generated

import (
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"strings"
)

const (
	P_KEY       = "id"
	R_KEY       = "right"
	L_KEY       = "left"
	LL_KEY      = "level"
	TN_KEY      = "tree_no"
	D_KEY       = "deleted"
	N_KEY       = "name"
	NS_KEY      = "ns_id"
	NS_COND_KEY = "ns_id = ?"
	UD_COND_KEY = "deleted = 0"
	DD_COND_KEY = "deleted = 1"
	CB_KEY      = "create_by"
	CA_KEY      = "create_at"
	MB_KEY      = "modify_by"
	MA_KEY      = "modify_at"
)

type Column struct {
	ColumnName      string
	Type            string
	Nullable        string
	TableName       string
	ColumnComment   string
	MaxLength       int64
	NumberPrecision int64
	ColumnType      string
	ColumnKey       string
	Default         interface{}
}

type Generator interface {
	GenFile(table string, columns []Column) *jen.File
}

func GenDeclAnonymousFunc() jen.Code {
	return jen.Id("fn").Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool"))
}

func UseDto(name string) jen.Code {
	return use("metis/test/second/model/dto", name)
}

func UseEntity(name string) jen.Code {
	return use("metis/test/second/model/entity", name)
}

func UseTime(name string) jen.Code {
	return use("time", name)
}

func UseCopier(name string) jen.Code {
	return use("github.com/jinzhu/copier", name)
}

func UseLogger(name string) jen.Code {
	return use("metis/util/logger", name)
}

func UseSql(name string) jen.Code {
	return use("database/sql", name)
}

func UseErrors(name string) jen.Code {
	return use("errors", name)
}

func UseContext(name string) jen.Code {
	return use("context", name)
}

func UseStrings(name string) jen.Code {
	return use("strings", name)
}

func UseFmt(name string) jen.Code {
	return use("fmt", name)
}

func UseGin(name string) jen.Code {
	return use("github.com/gin-gonic/gin", name)
}

func UseConstant(name string) jen.Code {
	return use("metis/config/constant", name)
}

func UseDatabase(name string) jen.Code {
	return use("metis/database", name)
}

func UseUtil(name string) jen.Code {
	return use("metis/util", name)
}

func UseZap(name string) jen.Code {
	return use("go.uber.org/zap", name)
}

func use(path, name string) jen.Code {
	return jen.Qual(path, name)
}

func InferColumn(code jen.Code, col string, columns []Column) jen.Code {
	if HasColumn(col, columns) {
		return code
	}
	return jen.Null()
}

func HasColumn(col string, columns []Column) bool {
	for _, column := range columns {
		if column.ColumnName == col {
			return true
		}
	}
	return false
}

func AllFields(columns []Column) string {
	fields := make([]string, len(columns))
	for i, column := range columns {
		fields[i] = column.ColumnName
	}
	return strings.Join(fields, ", ")
}

func RenderAndField(sn, field string) jen.Code {
	return jen.Op("&").Id(sn).Dot(strcase.ToCamel(field))
}

func RenderStarField(sn, field string) jen.Code {
	return jen.Op("*").Id(sn).Dot(strcase.ToCamel(field))
}

func AddNsValue(columns []Column) jen.Code {
	if HasColumn(NS_KEY, columns) {
		// jen.Var().Id("values").Index().Id("any")
		return jen.Id("values").Op("=").Id("append").Call(jen.Id("values"), jen.Add(UseUtil("GetNsID")).Call(jen.Id("ag").Dot("ctx")))
	}
	return jen.Null()
}

func AddNsSingleValue(columns []Column) jen.Code {
	if HasColumn(NS_KEY, columns) {
		// jen.Var().Id("values").Index().Id("any")
		return jen.Add(UseUtil("GetNsID")).Call(jen.Id("ag").Dot("ctx"))
	}
	return jen.Null()
}

func AddNsValueWithNew(columns []Column) jen.Code {
	if HasColumn(NS_KEY, columns) {
		// jen.Var().Id("values").Index().Id("any")
		return jen.Id("values").Op(":=").Id("append").Call(jen.Id("values"), jen.Add(UseUtil("GetNsID")).Call(jen.Id("ag").Dot("ctx")))
	}
	return jen.Null()
}

func AddNsValueWithName(columns []Column, name string) jen.Code {
	if HasColumn(NS_KEY, columns) {
		return jen.Id(name).Op("=").Id("append").Call(jen.Id(name), jen.Add(UseUtil("GetNsID")).Call(jen.Id("ag").Dot("ctx")))
	}
	return jen.Null()
}

func AddNsField(columns []Column, sqlBuilder string) jen.Code {
	if HasColumn(NS_KEY, columns) {
		return jen.Id(sqlBuilder).Dot("WriteString").Call(jen.Lit(NS_COND_KEY + " "))
	}
	return jen.Null()
}
