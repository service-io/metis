package test

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	p_key       = "id"
	r_key       = "right"
	l_key       = "left"
	ll_key      = "level"
	tn_key      = "tree_no"
	d_key       = "deleted"
	n_key       = "name"
	ud_cond_key = "deleted = 0"
	dd_cond_key = "deleted = 1"
	cb_key      = "create_by"
	ca_key      = "create_at"
	mb_key      = "modify_by"
	ma_key      = "modify_at"
)

func genDeclAnonymousFunc() jen.Code {
	return jen.Id("fn").Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool"))
}

func useDto(name string) jen.Code {
	return use("metis/test/second/model/dto", name)
}

func useEntity(name string) jen.Code {
	return use("metis/test/second/model/entity", name)
}

func useTime(name string) jen.Code {
	return use("time", name)
}

func useCopier(name string) jen.Code {
	return use("github.com/jinzhu/copier", name)
}

func useLogger(name string) jen.Code {
	return use("metis/util/logger", name)
}

func useSql(name string) jen.Code {
	return use("database/sql", name)
}

func useErrors(name string) jen.Code {
	return use("errors", name)
}

func useContext(name string) jen.Code {
	return use("context", name)
}

func useStrings(name string) jen.Code {
	return use("strings", name)
}

func useFmt(name string) jen.Code {
	return use("fmt", name)
}

func useGin(name string) jen.Code {
	return use("github.com/gin-gonic/gin", name)
}

func useConstant(name string) jen.Code {
	return use("metis/config/constant", name)
}

func useDatabase(name string) jen.Code {
	return use("metis/database", name)
}

func useUtil(name string) jen.Code {
	return use("metis/util", name)
}

func useZap(name string) jen.Code {
	return use("go.uber.org/zap", name)
}

func use(path, name string) jen.Code {
	return jen.Qual(path, name)
}

func inferColumn(code jen.Code, col string, columns []Column) jen.Code {
	if hasColumn(col, columns) {
		return code
	}
	return jen.Null()
}

func hasColumn(col string, columns []Column) bool {
	for _, column := range columns {
		if column.ColumnName == col {
			return true
		}
	}
	return false
}

func renderAndField(sn, field string) jen.Code {
	return jen.Op("&").Id(sn).Dot(strcase.ToCamel(field))
}

func renderStarField(sn, field string) jen.Code {
	return jen.Op("*").Id(sn).Dot(strcase.ToCamel(field))
}

func genInterfaceAutoGen(table string, columns []Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Type().Id("iAutoGen").Interface(
		jen.Id("SelectByID").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(useEntity(camel))),
		jen.Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectMaxLevel").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")),
		jen.Id("SelectMaxRight").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")),
		jen.Id("SelectMaxLeft").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")),
		jen.Id("SelectMaxTreeNo").Params().Params(jen.Id("int")),
		jen.Id("SelectAllPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectDirectPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectBrother").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectBrotherAndSelf").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectAncestorChain").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectAncestor").Params(
			jen.Id("id").Id("int64"), jen.Id("level").Id("int"),
		).Params(jen.Op("*").Add(useEntity(camel))),
		jen.Id("SelectParent").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(useEntity(camel))),
		jen.Id("SelectByTreeNoAndLevel").Params(
			jen.List(jen.Id("treeNo"), jen.Id("level")).Id("int"),
		).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectByLevel").Params(jen.Id("level").Id("int")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectRoot").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(useEntity(camel))),
		jen.Id("SelectLeaf").Params(
			jen.Id("id").Id("int64"),
			jen.List(jen.Id("page"), jen.Id("size")).Id("uint"),
		).Params(jen.Index().Op("*").Add(useEntity(camel)), jen.Id("int64")),
		jen.Id("SelectAllLeaf").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("SelectAllRoot").Params().Params(jen.Index().Op("*").Add(useEntity(camel))),
		jen.Id("Insert").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		).Params(jen.Id("int64")),
		jen.Id("InsertUnderNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.Id("pid").Id("int64"),
		).Params(jen.Id("int64")),
		jen.Id("InsertBetweenNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.List(
				jen.Id("pid"),
				jen.Id("sid"),
			).Id("int64"),
		).Params(jen.Id("int64")),
		jen.Id("BatchInsert").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
		).Params(jen.Index().Id("int64")),
		jen.Id("BatchInsertUnderNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
			jen.Id("pid").Id("int64"),
		).Params(jen.Index().Id("int64")),
		jen.Id("BatchInsertBetweenNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
			jen.List(
				jen.Id("pid"),
				jen.Id("sid"),
			).Id("int64"),
		).Params(jen.Index().Id("int64")),
		jen.Id("InsertNonNil").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		).Params(jen.Id("int64")),
		jen.Id("InsertNonNilUnderNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.Id("pid").Id("int64"),
		).Params(jen.Id("int64")),
		jen.Id("InsertNonNilBetweenNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.List(
				jen.Id("pid"),
				jen.Id("sid"),
			).Id("int64"),
		).Params(jen.Id("int64")),
		jen.Id("InsertWithFunc").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
		).Params(jen.Id("int64")),
		jen.Id("InsertWithFuncUnderNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.Id("pid").Id("int64"),
			jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
		).Params(jen.Id("int64")),
		jen.Id("InsertWithFuncBetweenNode").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.List(
				jen.Id("pid"),
				jen.Id("sid"),
			).Id("int64"),
			jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
		).Params(jen.Id("int64")),
		jen.Id("BatchInsertWithFunc").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
			jen.List(
				jen.Id("pid"),
				jen.Id("sid"),
			).Id("int64"),
			jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
		).Params(jen.Index().Id("int64")),
		jen.Id("DeleteByID").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id("id").Id("int64"),
		).Params(jen.Id("bool")),
		jen.Id("DeleteByIDs").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id("ids").Op("...").Id("int64"),
		).Params(jen.Id("bool")),
		jen.Id("BatchDeleteByID").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id("ids").Index().Id("int64"),
		).Params(jen.Id("bool")),
		jen.Id("UpdateByID").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		).Params(jen.Id("bool")),
		jen.Id("UpdateNonNilByID").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		).Params(jen.Id("bool")),
		jen.Id("UpdateWithFuncByID").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
			jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
		).Params(jen.Id("bool")),
		jen.Id("BatchUpdateWithFuncByID").Params(
			jen.Id("tx").Op("*").Qual("database/sql", "Tx"),
			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
			jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
		).Params(jen.Id("bool")),
	)
}

func genStructAutoGen() jen.Code {
	return jen.Type().Id("autoGen").Struct(jen.Id("ctx").Op("*").Add(useGin("Context")))
}

func genFuncGetDbCtx() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("getDbCtx").
		Params().Params(useContext("Context")).
		Block(
			jen.Return().Add(useContext("WithValue")).Call(
				jen.Add(useContext("Background")).Call(),
				jen.Add(useConstant("TraceIdKey")),
				jen.Id("ag").Dot("ctx").Dot("GetString").Call(useConstant("TraceIdKey")),
			),
		)
}

func genFuncMapperAll(table string, columns []Column) jen.Code {
	var codes = make([]jen.Code, len(columns))
	for i, column := range columns {
		codes[i] = renderAndField("r", column.ColumnName)
	}

	camel := strcase.ToCamel(table)

	return jen.Func().Id("mapperAll").Params().
		Params(
			jen.Op("*").Add(useEntity(camel)),
			jen.Index().Id("any"),
		).
		Block(
			jen.Var().Id("r").Op("=").Op("&").Add(useEntity(camel)).Values(),
			jen.Var().Id("cs").Op("=").Index().Id("any").Values(
				codes...,
			),
			jen.Return().List(
				jen.Id("r"),
				jen.Id("cs"),
			),
		)
}

func genFuncMapperNumeric() jen.Code {
	return jen.Func().Id("mapperNumeric").Params().Params(
		jen.Op("*").Id("T"),
		jen.Index().Id("any"),
	).Block(
		jen.Var().Id("r").Id("T"),
		jen.Var().Id("cs").Op("=").Index().Id("any").Values(jen.Op("&").Id("r")),
		jen.Return().List(
			jen.Op("&").Id("r"),
			jen.Id("cs"),
		),
	)
}

func allFields(columns []Column) string {
	fields := make([]string, len(columns))
	for i, column := range columns {
		fields[i] = column.ColumnName
	}
	return strings.Join(fields, ", ")
}

func genFuncTreeFields(table string, columns []Column) jen.Code {
	var sql string
	if hasColumn(d_key, columns) {
		sql = fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s = ? AND %s;", l_key, r_key, tn_key, table, p_key, ud_cond_key)
	} else {
		sql = fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s = ?;", l_key, r_key, tn_key, table, p_key)
	}
	return jen.Func().Id("treeFields").Params().Params(jen.Id("string")).Block(jen.Return().Id(sql))
}

func genFuncCalcInsertField(table string, columns []Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	var codes []jen.Code
	codes = append(
		codes, jen.Var().Id("fields").Index().Id("string"),
		jen.Var().Id("values").Index().Id("any"),
		jen.Var().Id("places").Index().Id("string"),
	)

	for _, column := range columns {
		columnName := column.ColumnName
		fieldName := strcase.ToCamel(columnName)
		code := jen.If(jen.Id("fn").Call(jen.Id(lowerCamel).Dot(fieldName))).Block(
			jen.Id("fields").Op("=").Id("append").Call(
				jen.Id("fields"),
				jen.Lit(columnName),
			),
			jen.Id("places").Op("=").Id("append").Call(
				jen.Id("places"),
				jen.Lit("?"),
			),
			jen.Id("values").Op("=").Id("append").Call(
				jen.Id("values"),
				jen.Id(lowerCamel).Dot(fieldName),
			),
		)
		codes = append(codes, code)
	}

	codes = append(
		codes, jen.Return().List(
			jen.Qual("strings", "Join").Call(
				jen.Id("fields"),
				jen.Lit(", "),
			),
			jen.Qual("strings", "Join").Call(
				jen.Id("places"),
				jen.Lit(", "),
			),
			jen.Id("values"),
		),
	)

	return jen.Func().Id("calcInsertField").Params(
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(
		jen.Id("string"),
		jen.Id("string"),
		jen.Index().Id("any"),
	).Block(codes...)
}

func genFuncCalcUpdateField(table string, columns []Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	var codes []jen.Code
	codes = append(
		codes, jen.Var().Id("fields").Index().Id("string"),
		jen.Var().Id("values").Index().Id("any"),
	)

	for _, column := range columns {
		columnName := column.ColumnName
		fieldName := strcase.ToCamel(columnName)
		code := jen.If(jen.Id("fn").Call(jen.Id(lowerCamel).Dot(fieldName))).Block(
			jen.Id("fields").Op("=").Id("append").Call(
				jen.Id("fields"),
				jen.Lit("SET "+columnName+" = ?"),
			),
			jen.Id("values").Op("=").Id("append").Call(
				jen.Id("values"),
				jen.Id(lowerCamel).Dot(fieldName),
			),
		)
		codes = append(codes, code)
	}

	codes = append(
		codes, jen.Return().List(
			jen.Qual("strings", "Join").Call(
				jen.Id("fields"),
				jen.Lit(", "),
			),
			jen.Id("values"),
		),
	)

	return jen.Func().Id("calcUpdateField").Params(
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(
		jen.Id("string"),
		jen.Index().Id("any"),
	).Block(codes...)
}

func genFuncInternalSelectNodeByIDs(table string, columns []Column) jen.Code {
	camel := strcase.ToCamel(table)
	//lowerCamel := strcase.ToLowerCamel(table)

	fields := allFields(columns)

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s ", fields, table, p_key)

	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalSelectNodeByIDs").Params(
		jen.Id("tx").Op("*").Add(useSql("Tx")),
		jen.Id("db").Op("*").Add(useSql("DB")),
		jen.Id("ids").Index().Id("int64"),
	).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID 列表: %+v 的数据"),
			jen.Id("ids"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql)),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
		).Else().Block(
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("IN (")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Add(useUtil("GenPlaceholder")).Call(jen.Id("ids"))),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(")")),
		),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(ud_cond_key+";")),
		jen.Id("errorHandler").Op(":=").Add(useUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Var().Id("stmt").Op("*").Add(useSql("Stmt")),
		jen.Var().Id("err").Id("error"),
		jen.If(jen.Id("tx").Op("!=").Id("nil")).Block(
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Defer().Add(useUtil("DeferClose")).Call(
				jen.Id("stmt"),
				jen.Id("errorHandler"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
		).Else().Block(
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op("=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Defer().Add(useUtil("DeferClose")).Call(
				jen.Id("stmt"),
				jen.Id("errorHandler"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
		),
		jen.Id("bindValues").Op(":=").Add(useUtil("ToAnyItems")).Call(jen.Id("ids")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("bindValues").Op("..."),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Add(useUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Add(useUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncInternalDirectInsert(table string, columns []Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectInsert").Params(
		jen.Id("tx").Op("*").Add(useSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Id("int64")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.If(jen.Id(lowerCamel).Dot("TreeNo").Op("==").Id("nil")).Block(jen.Id("panic").Call(jen.Lit("需要填充树号"))),
		jen.If(jen.Id(lowerCamel).Dot("Left").Op("==").Id("nil")).Block(jen.Id("panic").Call(jen.Lit("需要填充左值"))),
		jen.If(jen.Id(lowerCamel).Dot("Right").Op("==").Id("nil")).Block(jen.Id("panic").Call(jen.Lit("需要填充右值"))),
		jen.If(jen.Id(lowerCamel).Dot("Level").Op("==").Id("nil")).Block(jen.Id("panic").Call(jen.Lit("需要填充层级"))),
		jen.List(
			jen.Id("fields"),
			jen.Id("places"),
			jen.Id("values"),
		).Op(":=").Id("calcInsertField").Call(
			jen.Id(lowerCamel),
			jen.Id("fn"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("INSERT INTO "+table+"(")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(") VALUES (")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("places")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(");")),
		jen.Id("errorHandler").Op(":=").Add(useUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Add(useUtil("DeferClose")).Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("values").Op("..."),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("af"),
			jen.Id("err"),
		).Op(":=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("id"),
			jen.Id("err"),
		).Op(":=").Id("result").Dot("LastInsertId").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.If(jen.Id("af").Op("==").Lit(1)).Block(jen.Return().Id("id")),
		jen.Id("panic").Call(jen.Lit("插入失败")),
	)
}

func genFuncInternalUpdateNodeInBothWhenInsert(table string, columns []Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalUpdateNodeInBothWhenInsert").Params(
		jen.Id("tx").Op("*").Add(useSql("Tx")),
		jen.List(
			jen.Id("left"),
			jen.Id("right"),
			jen.Id("treeNo"),
		).Id("int"),
	).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Var().Id("firstSqlBuilder").Qual("strings", "Builder"),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("set <l_key> = <l_key> + 2 ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> > ? and")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Var().Id("secondSqlBuilder").Qual("strings", "Builder"),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("set <r_key> = <r_key> + 2 ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> > ? and")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Var().Id("thirdSqlBuilder").Qual("strings", "Builder"),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("set <l_key> = <l_key> + 1 ")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> = <r_key> + 1 ")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = <ll_key> + 1 ")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> >= ? and")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> <= ? and")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("thirdSqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("firstSqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("right"),
			jen.Id("treeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("secondSqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op("=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("right"),
			jen.Id("treeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("thirdSqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op("=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("left"),
			jen.Id("right"),
			jen.Id("treeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
	)
}

func genFuncinternalUpdateNodeInOnlyPrecursorWhenInsert() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalUpdateNodeInOnlyPrecursorWhenInsert").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.List(
			jen.Id("right"),
			jen.Id("treeNo"),
		).Id("int"),
	).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Var().Id("firstSqlBuilder").Qual("strings", "Builder"),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("set <l_key> = <l_key> + 2 ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> > ? and")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Var().Id("secondSqlBuilder").Qual("strings", "Builder"),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("set <r_key> = <r_key> + 2 ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> >= ? and")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("firstSqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("right"),
			jen.Id("treeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("secondSqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op("=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("right"),
			jen.Id("treeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
	)
}

func genFuncinternalInsertWithFunc() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalInsertWithFunc").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Id("int64")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("插入节点, 前驱: %+v, 后继: %+v, 节点: %+v"),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Op("*").Id(lowerCamel),
		),
		jen.If(jen.Id("pid").Op("==").Lit(0)).Block(
			jen.Return().Id("ag").Dot("internalDirectInsert").Call(
				jen.Id("tx"),
				jen.Id(lowerCamel),
				jen.Id("fn"),
			),
		),
		jen.Id("precursorNodes").Op(":=").Id("ag").Dot("internalSelectNodeByIDs").Call(
			jen.Id("tx"),
			jen.Id("nil"),
			jen.Index().Id("int64").Values(jen.Id("pid")),
		),
		jen.Id("nodeLen").Op(":=").Id("len").Call(jen.Id("precursorNodes")),
		jen.If(jen.Id("nodeLen").Op("==").Lit(1)).Block(
			jen.Id("precursor").Op(":=").Id("precursorNodes").Index(jen.Lit(0)),
			jen.Id(lowerCamel).Dot("TreeNo").Op("=").Id("precursor").Dot("TreeNo"),
			jen.Id("level").Op(":=").Op("*").Id("precursor").Dot("Level").Op("+").Lit(1),
			jen.Id(lowerCamel).Dot("Level").Op("=").Op("&").Id("level"),
			jen.If(jen.Id("sid").Op("==").Lit(0)).Block(
				jen.Id("right").Op(":=").Op("*").Id("precursor").Dot("Right").Op("+").Lit(1),
				jen.Id(lowerCamel).Dot("Left").Op("=").Id("precursor").Dot("Right"),
				jen.Id(lowerCamel).Dot("Right").Op("=").Op("&").Id("right"),
				jen.Id("ag").Dot("internalUpdateNodeInOnlyPrecursorWhenInsert").Call(
					jen.Id("tx"),
					jen.Op("*").Id("precursor").Dot("Right"),
					jen.Op("*").Id("precursor").Dot("TreeNo"),
				),
				jen.Return().Id("ag").Dot("internalDirectInsert").Call(
					jen.Id("tx"),
					jen.Id(lowerCamel),
					jen.Id("fn"),
				),
			).Else().Block(
				jen.Id("successorNodes").Op(":=").Id("ag").Dot("internalSelectNodeByIDs").Call(
					jen.Id("tx"),
					jen.Id("nil"),
					jen.Index().Id("int64").Values(jen.Id("pid")),
				),
				jen.If(jen.Id("len").Call(jen.Id("successorNodes")).Op("==").Lit(1)).Block(
					jen.Id("successor").Op(":=").Id("successorNodes").Index(jen.Lit(0)),
					jen.Id("right").Op(":=").Op("*").Id("successor").Dot("Right").Op("+").Lit(2),
					jen.Id(lowerCamel).Dot("Left").Op("=").Id("successor").Dot("Left"),
					jen.Id(lowerCamel).Dot("Right").Op("=").Op("&").Id("right"),
					jen.Id("ag").Dot("internalUpdateNodeInBothWhenInsert").Call(
						jen.Id("tx"),
						jen.Op("*").Id("successor").Dot("Left"),
						jen.Op("*").Id("successor").Dot("Right"),
						jen.Op("*").Id("successor").Dot("TreeNo"),
					),
					jen.Return().Id("ag").Dot("internalDirectInsert").Call(
						jen.Id("tx"),
						jen.Id(lowerCamel),
						jen.Id("fn"),
					),
				),
				jen.Id("panic").Call(jen.Lit("存在多个或不存在后继节点")),
			),
		),
		jen.Id("panic").Call(jen.Lit("存在多个或不存在前驱节点")),
	)
}

func genFuncinternalDirectDelete() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectDelete").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id("id").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("nodes").Op(":=").Id("ag").Dot("internalSelectNodeByIDs").Call(
			jen.Id("tx"),
			jen.Id("nil"),
			jen.Index().Id("int64").Values(jen.Id("id")),
		),
		jen.If(jen.Id("len").Call(jen.Id("nodes")).Op("==").Lit(1)).Block(
			jen.Id("node").Op(":=").Id("nodes").Index(jen.Lit(0)),
			jen.Id("right").Op(":=").Op("*").Id("node").Dot("Right"),
			jen.Id("left").Op(":=").Op("*").Id("node").Dot("Left"),
			jen.Id("treeNo").Op(":=").Op("*").Id("node").Dot("TreeNo"),
			jen.Id("delta").Op(":=").Id("right").Op("-").Id("left").Op("+").Lit(1),
			jen.Id("ag").Dot("internalUpdateNodeWhenDelete").Call(
				jen.Id("tx"),
				jen.Id("delta"),
				jen.Id("right"),
				jen.Id("treeNo"),
			),
			jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("set <d_key> = 1 ")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> >= ? and")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> <= ? and")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
			jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Defer().Qual("metis/util", "DeferClose").Call(
				jen.Id("stmt"),
				jen.Id("errorHandler"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
			jen.List(
				jen.Id("result"),
				jen.Id("err"),
			).Op(":=").Id("stmt").Dot("ExecContext").Call(
				jen.Id("ag").Dot("getDbCtx").Call(),
				jen.Id("left"),
				jen.Id("right"),
				jen.Id("treeNo"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
			jen.List(
				jen.Id("af"),
				jen.Id("err"),
			).Op(":=").Id("result").Dot("RowsAffected").Call(),
			jen.Id("errorHandler").Call(jen.Id("err")),
			jen.If(jen.Id("af").Op("==").Lit(1)).Block(jen.Return().Id("true")),
			jen.Id("panic").Call(jen.Lit("删除错误")),
		),
		jen.Id("panic").Call(jen.Lit("节点数错误")),
	)
}

func genFuncinternalUpdateNodeWhenDelete() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalUpdateNodeWhenDelete").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.List(
			jen.Id("delta"),
			jen.Id("right"),
			jen.Id("treeNo"),
		).Id("int"),
	).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Var().Id("firstSqlBuilder").Qual("strings", "Builder"),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("set <l_key> = <l_key> - ? ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> > ? and")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("firstSqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Var().Id("secondSqlBuilder").Qual("strings", "Builder"),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<table>")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("set <r_key> = <r_key> - ? ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> > ? and")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("secondSqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("firstSqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("delta"),
			jen.Id("right"),
			jen.Id("treeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("secondSqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op("=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("delta"),
			jen.Id("right"),
			jen.Id("treeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
	)
}

func genFuncgetDbCtx() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("getDbCtx").Params().Params(
		jen.Qual(
			"context", "Context",
		),
	).Block(
		jen.Return().Qual("context", "WithValue").Call(
			jen.Qual("context", "Background").Call(),
			jen.Id("constant").Dot("TraceIdKey"),
			jen.Id("ag").Dot("ctx").Dot("GetString").Call(jen.Id("constant").Dot("TraceIdKey")),
		),
	)
}

func genFuncSelectByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByID").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(useEntity(camel))).Block(
		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Index().Id("int64").Values(jen.Id("id"))),
		jen.If(jen.Id("len").Call(jen.Id("ds")).Op("==").Lit(1)).Block(jen.Return().Id("ds").Index(jen.Lit(0))),
		jen.Return().Id("nil"),
	)
}

func genFuncSelectByIDs() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Id("ids")),
		jen.Return().Id("ds"),
	)
}

func genFuncBatchSelectByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID 列表: %+v 的数据"),
			jen.Id("ids"),
		),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.Return().Id("ag").Dot("internalSelectNodeByIDs").Call(
			jen.Id("nil"),
			jen.Id("db"),
			jen.Id("ids"),
		),
	)
}

func genFuncSelectByName() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 NAME: %+v 的数据"),
			jen.Id("name"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<name_key> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("like ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("name"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectMaxLevel() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxLevel").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v 的最大层级"),
			jen.Id("treeNo"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("max(<ll_key>)")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("treeNo"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func genFuncSelectMaxRight() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxRight").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v 的最大层级"),
			jen.Id("treeNo"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("max(<r_key>)")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("treeNo"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func genFuncSelectMaxLeft() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxLeft").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v 的最大层级"),
			jen.Id("treeNo"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("max(<l_key>)")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("treeNo"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func genFuncSelectMaxTreeNo() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxTreeNo").Params().Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Info").Call(jen.Lit("查询最大TN")),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("max(<tn_key>)")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.Id("row").Op(":=").Id("db").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("sqlBuilder").Dot("String").Call(),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func genFuncSelectAllPosterity() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAllPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的所有子代(含自身)数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> > ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> < ? and")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("secondStmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectDirectPosterity() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectDirectPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的直系子代数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> > ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> < ? and")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("secondStmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Level").Op("+").Lit(1),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectBrother() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectBrother").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的兄弟数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<p_key> != ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("secondStmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Level").Op("+").Lit(1),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
			jen.Id("id"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectBrotherAndSelf() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectBrotherAndSelf").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的兄弟以及自身数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("secondStmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Level").Op("+").Lit(1),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectAncestorChain() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAncestorChain").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的祖链数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> < ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> > ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("secondStmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectAncestor() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAncestor").Params(
		jen.Id("id").Id("int64"),
		jen.Id("level").Id("int"),
	).Params(jen.Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的祖代(%+v)数据"),
			jen.Id("id"),
			jen.Id("level"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> < ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> > ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op("=").Id("secondStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Id("level"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectParent() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectParent").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的父节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> < ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> > ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op("=").Id("secondStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("Level").Op("-").Lit(1),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectByTreeNoAndLevel() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByTreeNoAndLevel").Params(
		jen.List(
			jen.Id("treeNo"),
			jen.Id("level"),
		).Id("int"),
	).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v LL: %+v 的同代数据"),
			jen.Id("treeNo"),
			jen.Id("level"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = ? and")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("treeNo"),
			jen.Id("level"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectByLevel() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByLevel").Params(jen.Id("level").Id("int")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 LL: %+v 的同代(跨树)数据"),
			jen.Id("level"),
		),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = ? and")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("level"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectRoot() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectRoot").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的根节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = 1 and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op("=").Id("secondStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Id("ds").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectLeaf() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectLeaf").Params(
		jen.Id("id").Id("int64"),
		jen.List(
			jen.Id("page"),
			jen.Id("size"),
		).Id("uint"),
	).Params(
		jen.Index().Op("*").Add(useEntity(camel)),
		jen.Id("int64"),
	).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("分页查询 ID: %+v 的叶子节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> >= ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> <= ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> + 1 = <r_key> and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Var().Id("noCondSql").Op("=").Id("sqlBuilder").Dot("String").Call().Op("+").Lit(";"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(" limit ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(" offset ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("secondStmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
			jen.Id("size"),
			jen.Parens(jen.Id("page").Op("-").Lit(1)).Op("*").Id("size"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("thirdStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("noCondSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("thirdStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op("=").Id("thirdStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Id("total").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int64")),
		),
		jen.Return().List(
			jen.Id("ds"),
			jen.Op("*").Id("total"),
		),
	)
}

func genFuncSelectAllLeaf() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAllLeaf").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的所有叶子节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeFields").Call(),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> >= ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<r_key> <= ? and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<l_key> + 1 = <r_key> and ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<tn_key> = ?")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Qual("metis/util", "HandleTx").Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
		),
		jen.Id("currentNode").Op(":=").Qual("metis/util", "Row").Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("secondStmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
		),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncSelectAllRoot() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAllRoot").Params().Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Info").Call(jen.Lit("查询的所有根节点数据")),
		jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("select ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("allFields").Call()),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("from <table> ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("where ")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<ll_key> = 1")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
		jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Qual("metis/database", "FetchDB").Call(),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(jen.Id("ag").Dot("getDbCtx").Call()),
		jen.Defer().Qual("metis/util", "DeferClose").Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Qual("metis/util", "Rows").Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func genFuncInsert() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("Insert").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Lit(0),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncInsertUnderNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertUnderNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("pid").Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncInsertBetweenNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertBetweenNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncBatchInsert() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsert").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
	).Params(jen.Index().Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Id(lowerCamel+"s"),
			jen.Lit(0),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Id("len").Call(jen.Id(lowerCamel+"s"))).Block(jen.Return().Id("ids")),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回记录数等于插入记录数时成功")),
	)
}

func genFuncBatchInsertUnderNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertUnderNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
		jen.Id("pid").Id("int64"),
	).Params(jen.Index().Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Id(lowerCamel+"s"),
			jen.Id("pid"),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Id("len").Call(jen.Id(lowerCamel+"s"))).Block(jen.Return().Id("ids")),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回记录数等于插入记录数时成功")),
	)
}

func genFuncBatchInsertBetweenNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertBetweenNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
	).Params(jen.Index().Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Id(lowerCamel+"s"),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Id("len").Call(jen.Id(lowerCamel+"s"))).Block(jen.Return().Id("ids")),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回记录数等于插入记录数时成功")),
	)
}

func genFuncInsertNonNil() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNil").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Lit(0),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncInsertNonNilUnderNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNilUnderNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("pid").Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncInsertNonNilBetweenNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNilBetweenNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncInsertWithFunc() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFunc").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Lit(0),
			jen.Lit(0),
			jen.Id("fn"),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncInsertWithFuncUnderNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFuncUnderNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("pid").Id("int64"),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Lit(0),
			jen.Id("fn"),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncInsertWithFuncBetweenNode() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFuncBetweenNode").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Id("fn"),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func genFuncBatchInsertWithFunc() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertWithFunc").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Index().Id("int64")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("插入至 PID: %+v SID: %+v 的同代数据"),
			jen.Id("pid"),
			jen.Id("sid"),
		),
		jen.Id("ids").Op(":=").Id("make").Call(
			jen.Index().Id("int64"),
			jen.Id("len").Call(jen.Id(lowerCamel+"s")),
		),
		jen.For(
			jen.List(
				jen.Id("i"),
				jen.Id(lowerCamel),
			).Op(":=").Range().Id(lowerCamel+"s"),
		).Block(
			jen.Id("ids").Index(jen.Id("i")).Op("=").Id("ag").Dot("internalInsertWithFunc").Call(
				jen.Id("tx"),
				jen.Id(lowerCamel),
				jen.Id("pid"),
				jen.Id("sid"),
				jen.Id("fn"),
			),
		),
		jen.Return().Id("ids"),
	)
}

func genFuncDeleteByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByID").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id("id").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
			jen.Id("tx"),
			jen.Index().Id("int64").Values(jen.Id("id")),
		),
	)
}

func genFuncDeleteByIDs() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByIDs").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id("ids").Op("...").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
			jen.Id("tx"),
			jen.Id("ids"),
		),
	)
}

func genFuncBatchDeleteByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchDeleteByID").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id("ids").Index().Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("删除 ID 列表: %+v 的数据"),
			jen.Id("ids"),
		),
		jen.For(
			jen.List(
				jen.Id("_"),
				jen.Id("id"),
			).Op(":=").Range().Id("ids"),
		).Block(
			jen.Id("ds").Op(":=").Id("ag").Dot("internalDirectDelete").Call(
				jen.Id("tx"),
				jen.Id("id"),
			),
			jen.If(jen.Op("!").Id("ds")).Block(jen.Id("panic").Call(jen.Lit("存在数据删除错误"))),
		),
		jen.Return().Id("true"),
	)
}

func genFuncUpdateByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateByID").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
	)
}

func genFuncUpdateNonNilByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateNonNilByID").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
	)
}

func genFuncUpdateWithFuncByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateWithFuncByID").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("fn"),
		),
	)
}

func genFuncBatchUpdateWithFuncByID() jen.Code {
	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchUpdateWithFuncByID").Params(
		jen.Id("tx").Op("*").Qual(
			"database/sql", "Tx",
		),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
		jen.Id("fn").Params(jen.Id("f").Id("any")).Params(jen.Id("bool")),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(jen.Lit("批量更新列表数据")),
		jen.For(
			jen.List(
				jen.Id("_"),
				jen.Id(lowerCamel),
			).Op(":=").Range().Id(lowerCamel+"s"),
		).Block(
			jen.If(jen.Id(lowerCamel).Dot("ID").Op("==").Id("nil")).Block(jen.Id("panic").Call(jen.Lit("ID 字段不能为空"))),
			jen.Id("id").Op(":=").Op("*").Id(lowerCamel).Dot("ID"),
			jen.List(
				jen.Id("fields"),
				jen.Id("values"),
			).Op(":=").Id("calcUpdateField").Call(
				jen.Id(lowerCamel),
				jen.Id("fn"),
			),
			jen.Var().Id("sqlBuilder").Qual("strings", "Builder"),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("update ")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<table> ")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(" where ")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<p_key> = ?")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("<deleted_cond>")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(";")),
			jen.Id("values").Op("=").Id("append").Call(
				jen.Id("values"),
				jen.Id("id"),
			),
			jen.Id("errorHandler").Op(":=").Qual("metis/util", "ErrToLogAndPanic").Call(jen.Id("recorder")),
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Id("errorHandler").Call(jen.Id("err")),
			jen.List(
				jen.Id("result"),
				jen.Id("err"),
			).Op(":=").Id("stmt").Dot("ExecContext").Call(
				jen.Id("ag").Dot("getDbCtx").Call(),
				jen.Id("values").Op("..."),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
			jen.List(
				jen.Id("af"),
				jen.Id("err"),
			).Op(":=").Id("result").Dot("RowsAffected").Call(),
			jen.Id("errorHandler").Call(jen.Id("err")),
			jen.If(jen.Id("af").Op("!=").Lit(1)).Block(jen.Id("panic").Call(jen.Lit("更新错误"))),
			jen.Id("err").Op("=").Id("stmt").Dot("Close").Call(),
		),
		jen.Return().Id("true"),
	)
}

func genTreeRepositoryFile(table string, columns []Column) *jen.File {
	ret := jen.NewFile(table)
	ret.Add(genInterfaceAutoGen(table, columns))
	ret.Add(genStructAutoGen())
	ret.Add(genFuncGetDbCtx())
	ret.Add(genFuncMapperAll(table, columns))
	ret.Add(genFuncMapperNumeric())
	ret.Add(genFuncTreeFields(table, columns))
	ret.Add(genFuncCalcInsertField(table, columns))
	ret.Add(genFuncCalcUpdateField(table, columns))
	ret.Add(genFuncInternalSelectNodeByIDs(table, columns))
	ret.Add(genFuncInternalDirectInsert(table, columns))
	ret.Add(genFuncInternalUpdateNodeInBothWhenInsert(table, columns))

	//ret.Add(genDeclAt136())
	//ret.Add(genDeclAt348())
	//ret.Add(genDeclAt3223())
	//ret.Add(genFuncmapperAll())
	//ret.Add(genFuncmapperNumeric())
	//ret.Add(genFuncallFields())
	//ret.Add(genFunctreeFields())
	//ret.Add(genFunccalcInsertField())
	//ret.Add(genFunccalcUpdateField())
	//ret.Add(genFuncinternalSelectNodeByIDs())
	//ret.Add(genFuncinternalDirectInsert())
	//ret.Add(genFuncinternalUpdateNodeInBothWhenInsert())
	//ret.Add(genFuncinternalUpdateNodeInOnlyPrecursorWhenInsert())
	//ret.Add(genFuncinternalInsertWithFunc())
	//ret.Add(genFuncinternalDirectDelete())
	//ret.Add(genFuncinternalUpdateNodeWhenDelete())
	//ret.Add(genFuncgetDbCtx())
	//ret.Add(genFuncSelectByID())
	//ret.Add(genFuncSelectByIDs())
	//ret.Add(genFuncBatchSelectByID())
	//ret.Add(genFuncSelectByName())
	//ret.Add(genFuncSelectMaxLevel())
	//ret.Add(genFuncSelectMaxRight())
	//ret.Add(genFuncSelectMaxLeft())
	//ret.Add(genFuncSelectMaxTreeNo())
	//ret.Add(genFuncSelectAllPosterity())
	//ret.Add(genFuncSelectDirectPosterity())
	//ret.Add(genFuncSelectBrother())
	//ret.Add(genFuncSelectBrotherAndSelf())
	//ret.Add(genFuncSelectAncestorChain())
	//ret.Add(genFuncSelectAncestor())
	//ret.Add(genFuncSelectParent())
	//ret.Add(genFuncSelectByTreeNoAndLevel())
	//ret.Add(genFuncSelectByLevel())
	//ret.Add(genFuncSelectRoot())
	//ret.Add(genFuncSelectLeaf())
	//ret.Add(genFuncSelectAllLeaf())
	//ret.Add(genFuncSelectAllRoot())
	//ret.Add(genFuncInsert())
	//ret.Add(genFuncInsertUnderNode())
	//ret.Add(genFuncInsertBetweenNode())
	//ret.Add(genFuncBatchInsert())
	//ret.Add(genFuncBatchInsertUnderNode())
	//ret.Add(genFuncBatchInsertBetweenNode())
	//ret.Add(genFuncInsertNonNil())
	//ret.Add(genFuncInsertNonNilUnderNode())
	//ret.Add(genFuncInsertNonNilBetweenNode())
	//ret.Add(genFuncInsertWithFunc())
	//ret.Add(genFuncInsertWithFuncUnderNode())
	//ret.Add(genFuncInsertWithFuncBetweenNode())
	//ret.Add(genFuncBatchInsertWithFunc())
	//ret.Add(genFuncDeleteByID())
	//ret.Add(genFuncDeleteByIDs())
	//ret.Add(genFuncBatchDeleteByID())
	//ret.Add(genFuncUpdateByID())
	//ret.Add(genFuncUpdateNonNilByID())
	//ret.Add(genFuncUpdateWithFuncByID())
	//ret.Add(genFuncBatchUpdateWithFuncByID())
	return ret
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
			Tag:             "id",
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
			Tag:             "title",
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
			Tag:             "startAt",
			MaxLength:       0,
			NumberPrecision: 0,
			ColumnType:      "timestamp",
			ColumnKey:       "",
			Default:         "CURRENT_TIMESTAMP",
		},
	)

	//columns := make([]Column, 0)
	//var params = make(map[string]string)
	//params["parseTime"] = "true"
	//cfg := mysql.Config{
	//	User:   "root",
	//	Passwd: "root",
	//	Net:    "tcp",
	//	Addr:   "localhost:3307",
	//	DBName: "metis",
	//	Params: params,
	//}
	//
	//var err error
	//db, err := sql.Open("mysql", cfg.FormatDSN())
	//rows, err := db.Query(
	//	fmt.Sprintf(
	//		`SELECT
	//	COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT,CHARACTER_MAXIMUM_LENGTH,COLUMN_TYPE,NUMERIC_PRECISION,COLUMN_KEY,COLUMN_DEFAULT
	//	FROM information_schema.COLUMNS
	//	WHERE table_schema = DATABASE()  AND TABLE_NAME = '%s'`, table,
	//	),
	//)
	//
	//if err != nil {
	//	log.Printf("table rows is nil with table:%s error: %v \n", table, err)
	//	return columns
	//}
	//
	//if rows == nil {
	//	log.Printf("rows is nil with table:%s \n", table)
	//	return columns
	//}
	//
	//defer func() {
	//	_ = rows.Close()
	//}()
	//
	//for rows.Next() {
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
	//}

	return columns
}

func TestRepository(t *testing.T) {

	strcase.ConfigureAcronym("ID", "id")
	strcase.ConfigureAcronym("id", "ID")

	table := lowerCamel
	columns := getColumns(table)

	f := genTreeRepositoryFile(table, columns)
	fmt.Printf("%#v\n", f)
	autogenFilePath := "second/module/user/repository/" + table + "/autogen.go"
	if err := os.MkdirAll(filepath.Dir(autogenFilePath), 0766); err != nil {
		panic(err)
	}
	wr, err := os.OpenFile(autogenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	err = f.Render(wr)
}
