package test

// import (
// 	"fmt"
// 	"github.com/dave/jennifer/jen"
// 	"github.com/iancoleman/strcase"
//   "os"
// 	"path/filepath"
// 	"strings"
// 	"testing"
// )
//
// const (
// 	p_key       = "id"
// 	r_key       = "right"
// 	l_key       = "left"
// 	ll_key      = "level"
// 	tn_key      = "tree_no"
// 	d_key       = "deleted"
// 	n_key       = "name"
// 	ud_cond_key = "deleted = 0"
// 	dd_cond_key = "deleted = 1"
// 	cb_key      = "create_by"
// 	ca_key      = "create_at"
// 	mb_key      = "modify_by"
// 	ma_key      = "modify_at"
// )
//
// func genDeclAnonymousFunc() jen.Code {
// 	return jen.Id("fn").Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool"))
// }
//
// func useDto(name string) jen.Code {
// 	return use("metis/test/second/model/dto", name)
// }
//
// func useEntity(name string) jen.Code {
// 	return use("metis/test/second/model/entity", name)
// }
//
// func useTime(name string) jen.Code {
// 	return use("time", name)
// }
//
// func useCopier(name string) jen.Code {
// 	return use("github.com/jinzhu/copier", name)
// }
//
// func useLogger(name string) jen.Code {
// 	return use("metis/util/logger", name)
// }
//
// func useSql(name string) jen.Code {
// 	return use("database/sql", name)
// }
//
// func useErrors(name string) jen.Code {
// 	return use("errors", name)
// }
//
// func useContext(name string) jen.Code {
// 	return use("context", name)
// }
//
// func useStrings(name string) jen.Code {
// 	return use("strings", name)
// }
//
// func useFmt(name string) jen.Code {
// 	return use("fmt", name)
// }
//
// func useGin(name string) jen.Code {
// 	return use("github.com/gin-gonic/gin", name)
// }
//
// func useConstant(name string) jen.Code {
// 	return use("metis/config/constant", name)
// }
//
// func useDatabase(name string) jen.Code {
// 	return use("metis/database", name)
// }
//
// func useUtil(name string) jen.Code {
// 	return use("metis/util", name)
// }
//
// func useZap(name string) jen.Code {
// 	return use("go.uber.org/zap", name)
// }
//
// func use(path, name string) jen.Code {
// 	return jen.Qual(path, name)
// }
//
// func inferColumn(code jen.Code, col string, columns []Column) jen.Code {
// 	if hasColumn(col, columns) {
// 		return code
// 	}
// 	return jen.Null()
// }
//
// func hasColumn(col string, columns []Column) bool {
// 	for _, column := range columns {
// 		if column.ColumnName == col {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func renderAndField(sn, field string) jen.Code {
// 	return jen.Op("&").Id(sn).Dot(strcase.ToCamel(field))
// }
//
// func renderStarField(sn, field string) jen.Code {
// 	return jen.Op("*").Id(sn).Dot(strcase.ToCamel(field))
// }
//
// func genInterfaceAutoGen(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Comment("iAutoGen 该接口自动生成, 请勿修改").Line().Type().Id("iAutoGen").Interface(
// 		inferColumn(jen.Id("SelectByID").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(useEntity(camel))), "id", columns),
// 		inferColumn(jen.Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))), "id", columns),
// 		inferColumn(jen.Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).Params(jen.Index().Op("*").Add(useEntity(camel))), "id", columns),
//     jen.Line(),
// 		inferColumn(jen.Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(useEntity(camel))), "name", columns),
//     jen.Line(),
// 		jen.Id("Insert").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel).Op("*").Add(useEntity(camel))).Params(jen.Id("int64")),
// 		jen.Id("InsertNonNil").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel).Op("*").Add(useEntity(camel))).Params(jen.Id("int64")),
// 		jen.Id("InsertWithFunc").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel).Op("*").Add(useEntity(camel)), genDeclAnonymousFunc()).Params(jen.Id("int64")),
// 		jen.Id("BatchInsert").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel))).Params(jen.Index().Id("int64")),
// 		jen.Id("BatchInsertWithFunc").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)), jen.List(jen.Id("pid"), jen.Id("sid")).Id("int64"), genDeclAnonymousFunc()).Params(jen.Index().Id("int64")),
//     jen.Line(),
// 		inferColumn(jen.Id("DeleteByID").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id("id").Id("int64")).Params(jen.Id("bool")), "id", columns),
// 		inferColumn(jen.Id("DeleteByIDs").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id("ids").Op("...").Id("int64")).Params(jen.Id("bool")), "id", columns),
// 		inferColumn(jen.Id("BatchDeleteByID").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id("ids").Index().Id("int64")).Params(jen.Id("bool")), "id", columns),
//     jen.Line(),
// 		inferColumn(jen.Id("UpdateByID").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel).Op("*").Add(useEntity(camel))).Params(jen.Id("bool")), "id", columns),
// 		inferColumn(jen.Id("UpdateNonNilByID").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel).Op("*").Add(useEntity(camel))).Params(jen.Id("bool")), "id", columns),
// 		inferColumn(jen.Id("UpdateWithFuncByID").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel).Op("*").Add(useEntity(camel)), genDeclAnonymousFunc()).Params(jen.Id("bool")), "id", columns),
// 		inferColumn(jen.Id("BatchUpdateWithFuncByID").Params(jen.Id("tx").Op("*").Add(useSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)), genDeclAnonymousFunc()).Params(jen.Id("bool")), "id", columns),
// 	)
// }
//
// func genStructAutoGen() jen.Code {
// 	return jen.Line().Comment("autoGen 该结构体自动生成, 请勿修改").Line().Type().Id("autoGen").Struct(jen.Id("ctx").Op("*").Add(useGin("Context")))
// }
//
// func genFuncGetDbCtx() jen.Code {
// 	return jen.Line().Comment("getDbCtx 获取 DB 的初始上下文").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("getDbCtx").
// 		Params().Params(useContext("Context")).
// 		Block(
// 			jen.Return().Add(useContext("WithValue")).Call(
// 				jen.Add(useContext("Background")).Call(),
// 				jen.Add(useConstant("TraceIdKey")),
// 				jen.Id("ag").Dot("ctx").Dot("GetString").Call(useConstant("TraceIdKey")),
// 			),
// 		)
// }
//
// func genFuncMapperAll(table string, columns []Column) jen.Code {
// 	var codes = make([]jen.Code, len(columns))
// 	for i, column := range columns {
// 		codes[i] = renderAndField("r", column.ColumnName)
// 	}
//
// 	camel := strcase.ToCamel(table)
//
// 	return jen.Line().Comment("mapperAll 映射实体的所有字体").Line().Func().Id("mapperAll").Params().
// 		Params(
// 			jen.Op("*").Add(useEntity(camel)),
// 			jen.Index().Id("any"),
// 		).
// 		Block(
// 			jen.Var().Id("r").Op("=").Op("&").Add(useEntity(camel)).Values(),
// 			jen.Var().Id("cs").Op("=").Index().Id("any").Values(
// 				codes...,
// 			),
// 			jen.Return().List(
// 				jen.Id("r"),
// 				jen.Id("cs"),
// 			),
// 		)
// }
//
// func genFuncMapperNumeric() jen.Code {
// 	return jen.Line().Comment("mapperNumeric 映射数值型").Line().Func().Id("mapperNumeric").Types(jen.Id("T").Union(jen.Int(), jen.Int64())).Params().Params(
// 		jen.Op("*").Id("T"),
// 		jen.Index().Id("any"),
// 	).Block(
// 		jen.Var().Id("r").Id("T"),
// 		jen.Var().Id("cs").Op("=").Index().Id("any").Values(jen.Op("&").Id("r")),
// 		jen.Return().List(
// 			jen.Op("&").Id("r"),
// 			jen.Id("cs"),
// 		),
// 	)
// }
//
// func allFields(columns []Column) string {
// 	fields := make([]string, len(columns))
// 	for i, column := range columns {
// 		fields[i] = column.ColumnName
// 	}
// 	return strings.Join(fields, ", ")
// }
//
// func genFuncCalcInsertField(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	var codes []jen.Code
// 	codes = append(
// 		codes, jen.Var().Id("fields").Index().Id("string"),
// 		jen.Var().Id("values").Index().Id("any"),
// 		jen.Var().Id("places").Index().Id("string"),
// 	)
//
// 	for _, column := range columns {
// 		columnName := column.ColumnName
// 		fieldName := strcase.ToCamel(columnName)
// 		code := jen.If(jen.Id("fn").Call(jen.Id(lowerCamel).Dot(fieldName))).Block(
// 			jen.Id("fields").Op("=").Id("append").Call(
// 				jen.Id("fields"),
// 				jen.Lit(columnName),
// 			),
// 			jen.Id("places").Op("=").Id("append").Call(
// 				jen.Id("places"),
// 				jen.Lit("?"),
// 			),
// 			jen.Id("values").Op("=").Id("append").Call(
// 				jen.Id("values"),
// 				jen.Id(lowerCamel).Dot(fieldName),
// 			),
// 		)
// 		codes = append(codes, code)
// 	}
//
// 	codes = append(
// 		codes, jen.Return().List(
// 			jen.Add(useStrings("Join")).Call(
// 				jen.Id("fields"),
// 				jen.Lit(", "),
// 			),
// 			jen.Add(useStrings("Join")).Call(
// 				jen.Id("places"),
// 				jen.Lit(", "),
// 			),
// 			jen.Id("values"),
// 		),
// 	)
//
// 	return jen.Line().Comment("calcInsertField 计算待插入的字段").Line().Func().Id("calcInsertField").Params(
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		genDeclAnonymousFunc(),
// 	).Params(
// 		jen.Id("string"),
// 		jen.Id("string"),
// 		jen.Index().Id("any"),
// 	).Block(codes...)
// }
//
// func genFuncCalcUpdateField(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	var codes []jen.Code
// 	codes = append(
// 		codes, jen.Var().Id("fields").Index().Id("string"),
// 		jen.Var().Id("values").Index().Id("any"),
// 	)
//
// 	for _, column := range columns {
// 		columnName := column.ColumnName
// 		fieldName := strcase.ToCamel(columnName)
// 		code := jen.If(jen.Id("fn").Call(jen.Id(lowerCamel).Dot(fieldName))).Block(
// 			jen.Id("fields").Op("=").Id("append").Call(
// 				jen.Id("fields"),
// 				jen.Lit(columnName+" = ?"),
// 			),
// 			jen.Id("values").Op("=").Id("append").Call(
// 				jen.Id("values"),
// 				jen.Id(lowerCamel).Dot(fieldName),
// 			),
// 		)
// 		codes = append(codes, code)
// 	}
//
// 	codes = append(
// 		codes, jen.Return().List(
// 			jen.Add(useStrings("Join")).Call(
// 				jen.Id("fields"),
// 				jen.Lit(", "),
// 			),
// 			jen.Id("values"),
// 		),
// 	)
//
// 	return jen.Line().Comment("calcUpdateField 计算待更新的字段").Line().Func().Id("calcUpdateField").Params(
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		genDeclAnonymousFunc(),
// 	).Params(
// 		jen.Id("string"),
// 		jen.Index().Id("any"),
// 	).Block(codes...)
// }
//
// func genFuncInternalSelectByIDs(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	// lowerCamel := strcase.ToLowerCamel(table)
//
// 	fields := allFields(columns)
//
// 	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s ", fields, table, p_key)
//
// 	return jen.Line().Comment("internalSelectByIDs 根据 ID 列表插入节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalSelectByIDs").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id("db").Op("*").Add(useSql("DB")),
// 		jen.Id("ids").Index().Id("int64"),
// 	).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
// 			jen.Lit("查询 ID 列表: %+v 的数据"),
// 			jen.Id("ids"),
// 		),
// 		jen.Var().Id("sqlBuilder").Add(useStrings("Builder")),
// 		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql)),
// 		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(
// 			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
// 		).Else().Block(
// 			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("IN (")),
// 			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Add(useUtil("GenPlaceholder")).Call(jen.Id("ids"))),
// 			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(")")),
// 		),
// 		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(ud_cond_key+";")),
// 		jen.Id("errorHandler").Op(":=").Add(useUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
// 		jen.Var().Id("stmt").Op("*").Add(useSql("Stmt")),
// 		jen.Var().Id("err").Id("error"),
// 		jen.If(jen.Id("tx").Op("!=").Id("nil")).Block(
// 			jen.List(
// 				jen.Id("stmt"),
// 				jen.Id("err"),
// 			).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("stmt"),
// 				jen.Id("errorHandler"),
// 			),
// 			jen.Id("errorHandler").Call(jen.Id("err")),
// 		).Else().Block(
// 			jen.List(
// 				jen.Id("stmt"),
// 				jen.Id("err"),
// 			).Op("=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("stmt"),
// 				jen.Id("errorHandler"),
// 			),
// 			jen.Id("errorHandler").Call(jen.Id("err")),
// 		),
// 		jen.Id("bindValues").Op(":=").Add(useUtil("ToAnyItems")).Call(jen.Id("ids")),
// 		jen.List(
// 			jen.Id("rows"),
// 			jen.Id("err"),
// 		).Op(":=").Id("stmt").Dot("QueryContext").Call(
// 			jen.Id("ag").Dot("getDbCtx").Call(),
// 			jen.Id("bindValues").Op("..."),
// 		),
// 		jen.Id("errorHandler").Call(jen.Id("err")),
// 		jen.Defer().Add(useUtil("DeferClose")).Call(
// 			jen.Id("rows"),
// 			jen.Id("errorHandler"),
// 		),
// 		jen.Id("ds").Op(":=").Add(useUtil("Rows")).Call(
// 			jen.Id("rows"),
// 			jen.Id("mapperAll"),
// 		),
// 		jen.Return().Id("ds"),
// 	)
// }
//
// func genFuncInternalDirectInsert(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Line().Comment("internalDirectInsert 直接插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectInsert").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		genDeclAnonymousFunc(),
// 	).Params(jen.Id("int64")).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
//     jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
//       jen.Lit("插入数据: %+v"),
//       jen.Op("*").Id(lowerCamel),
//     ),
// 		jen.List(
// 			jen.Id("fields"),
// 			jen.Id("places"),
// 			jen.Id("values"),
// 		).Op(":=").Id("calcInsertField").Call(
// 			jen.Id(lowerCamel),
// 			jen.Id("fn"),
// 		),
// 		jen.Var().Id("sqlBuilder").Add(useStrings("Builder")),
// 		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("INSERT INTO "+table+"(")),
// 		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
// 		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(") VALUES (")),
// 		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("places")),
// 		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(");")),
// 		jen.Id("errorHandler").Op(":=").Add(useUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
// 		jen.List(
// 			jen.Id("stmt"),
// 			jen.Id("err"),
// 		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
// 		jen.Defer().Add(useUtil("DeferClose")).Call(
// 			jen.Id("stmt"),
// 			jen.Id("errorHandler"),
// 		),
// 		jen.Id("errorHandler").Call(jen.Id("err")),
// 		jen.List(
// 			jen.Id("result"),
// 			jen.Id("err"),
// 		).Op(":=").Id("stmt").Dot("ExecContext").Call(
// 			jen.Id("ag").Dot("getDbCtx").Call(),
// 			jen.Id("values").Op("..."),
// 		),
// 		jen.Id("errorHandler").Call(jen.Id("err")),
// 		jen.List(
// 			jen.Id("af"),
// 			jen.Id("err"),
// 		).Op(":=").Id("result").Dot("RowsAffected").Call(),
// 		jen.Id("errorHandler").Call(jen.Id("err")),
// 		jen.List(
// 			jen.Id("id"),
// 			jen.Id("err"),
// 		).Op(":=").Id("result").Dot("LastInsertId").Call(),
// 		jen.Id("errorHandler").Call(jen.Id("err")),
// 		jen.If(jen.Id("af").Op("==").Lit(1)).Block(jen.Return().Id("id")),
// 		jen.Id("panic").Call(jen.Lit("插入失败")),
// 	)
// }
//
// func genFuncInternalDirectDelete(table string, columns []Column) jen.Code {
//
// 	var sql0 string
// 	if hasColumn(d_key, columns) {
// 		sql0 = fmt.Sprintf("UPDATE %s SET %s = 1 WHERE %s = ? AND %s;", table, d_key, p_key, ud_cond_key)
// 	} else {
// 		sql0 = fmt.Sprintf("DELETE FROM %s WHERE %s = ?;", table, p_key)
// 	}
//
// 	return jen.Line().Comment("internalDirectDelete 直接删除(逻辑 or 物理)").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectDelete").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id("id").Id("int64"),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
//
//     jen.Id("errorHandler").Op(":=").Add(useUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
//     jen.List(
//       jen.Id("stmt"),
//       jen.Id("err"),
//     ).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
//     jen.Defer().Add(useUtil("DeferClose")).Call(
//       jen.Id("stmt"),
//       jen.Id("errorHandler"),
//     ),
//     jen.Id("errorHandler").Call(jen.Id("err")),
//     jen.List(
//       jen.Id("result"),
//       jen.Id("err"),
//     ).Op(":=").Id("stmt").Dot("ExecContext").Call(
//       jen.Id("ag").Dot("getDbCtx").Call(),
//       jen.Id("id"),
//     ),
//     jen.Id("errorHandler").Call(jen.Id("err")),
//     jen.List(
//       jen.Id("af"),
//       jen.Id("err"),
//     ).Op(":=").Id("result").Dot("RowsAffected").Call(),
//     jen.Id("errorHandler").Call(jen.Id("err")),
//     jen.If(jen.Id("af").Op("==").Lit(1)).Block(jen.Return().Id("true")),
//     jen.Id("panic").Call(jen.Lit("删除错误")),
// 	)
// }
//
// func genFuncSelectByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	// lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Line().Comment("SelectByID 根据 ID 查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByID").Params(jen.Id("id").Id("int64")).
// 		Params(jen.Op("*").Add(useEntity(camel))).Block(
// 		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Index().Id("int64").Values(jen.Id("id"))),
// 		jen.If(jen.Id("len").Call(jen.Id("ds")).Op("==").Lit(1)).Block(jen.Return().Id("ds").Index(jen.Lit(0))),
// 		jen.Return().Id("nil"),
// 	)
// }
//
// func genFuncSelectByIDs(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	// lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Line().Comment("SelectByIDs 根据 ID 列表查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).
// 		Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
// 		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Id("ids")),
// 		jen.Return().Id("ds"),
// 	)
// }
//
// func genFuncBatchSelectByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	// lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Line().Comment("BatchSelectByID 根据 ID 批量查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).
// 		Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
// 			jen.Lit("查询 ID 列表: %+v 的数据"),
// 			jen.Id("ids"),
// 		),
// 		jen.Id("db").Op(":=").Add(useDatabase("FetchDB")).Call(),
// 		jen.Return().Id("ag").Dot("internalSelectByIDs").Call(
// 			jen.Id("nil"),
// 			jen.Id("db"),
// 			jen.Id("ids"),
// 		),
// 	)
// }
//
// func genFuncSelectByName(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	// lowerCamel := strcase.ToLowerCamel(table)
// 	var sql0 string
// 	if hasColumn(d_key, columns) {
// 		sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s like ? AND %s;", allFields(columns), table, n_key, ud_cond_key)
// 	} else {
// 		sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s like ?;", allFields(columns), table, n_key)
// 	}
// 	return jen.Line().Comment("SelectByName 根据名称查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(useEntity(camel))).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
// 			jen.Lit("查询 NAME: %+v 的数据"),
// 			jen.Id("name"),
// 		),
// 		jen.Id("db").Op(":=").Add(useDatabase("FetchDB")).Call(),
// 		jen.Id("errorHandler").Op(":=").Add(useUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
// 		jen.List(
// 			jen.Id("stmt"),
// 			jen.Id("err"),
// 		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
// 		jen.Defer().Add(useUtil("DeferClose")).Call(
// 			jen.Id("stmt"),
// 			jen.Id("errorHandler"),
// 		),
// 		jen.Id("errorHandler").Call(jen.Id("err")),
// 		jen.List(
// 			jen.Id("rows"),
// 			jen.Id("err"),
// 		).Op(":=").Id("stmt").Dot("QueryContext").Call(
// 			jen.Id("ag").Dot("getDbCtx").Call(),
// 			jen.Id("name"),
// 		),
// 		jen.Id("errorHandler").Call(jen.Id("err")),
// 		jen.Defer().Add(useUtil("DeferClose")).Call(
// 			jen.Id("rows"),
// 			jen.Id("errorHandler"),
// 		),
// 		jen.Id("ds").Op(":=").Add(useUtil("Rows")).Call(
// 			jen.Id("rows"),
// 			jen.Id("mapperAll"),
// 		),
// 		jen.Return().Id("ds"),
// 	)
// }
//
// func genFuncInsert(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("Insert").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 	).Params(jen.Id("int64")).Block(
// 		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
// 			jen.Id("tx"),
// 			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
// 			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
// 		),
// 		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
// 		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
// 	)
// }
//
// func genFuncBatchInsert(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Line().Comment("BatchInsert 批量插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsert").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 	).Params(jen.Index().Id("int64")).Block(
// 		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
// 			jen.Id("tx"),
// 			jen.Id(lowerCamel+"s"),
// 			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
// 		),
// 		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Id("len").Call(jen.Id(lowerCamel+"s"))).Block(jen.Return().Id("ids")),
// 		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回记录数等于插入记录数时成功")),
// 	)
// }
//
// func genFuncInsertNonNil(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Line().Comment("InsertNonNil 插入非空字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNil").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 	).Params(jen.Id("int64")).Block(
// 		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
// 			jen.Id("tx"),
// 			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
// 			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
// 		),
// 		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
// 		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
// 	)
// }
//
// func genFuncInsertWithFunc(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Line().Comment("InsertWithFunc 根据函数插入字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFunc").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		genDeclAnonymousFunc(),
// 	).Params(jen.Id("int64")).Block(
// 		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
// 			jen.Id("tx"),
// 			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
// 			jen.Id("fn"),
// 		),
// 		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
// 		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
// 	)
// }
//
// func genFuncBatchInsertWithFunc(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Line().Comment("BatchInsertWithFunc 根据函数批量插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertWithFunc").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 		genDeclAnonymousFunc(),
// 	).Params(jen.Index().Id("int64")).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 		jen.Id("recorder").Dot("Info").Call(
// 			jen.Lit("批量插入数据"),
// 		),
// 		jen.Id("ids").Op(":=").Id("make").Call(
// 			jen.Index().Id("int64"),
// 			jen.Id("len").Call(jen.Id(lowerCamel+"s")),
// 		),
// 		jen.For(
// 			jen.List(
// 				jen.Id("i"),
// 				jen.Id(lowerCamel),
// 			).Op(":=").Range().Id(lowerCamel+"s"),
// 		).Block(
// 			jen.Id("ids").Index(jen.Id("i")).Op("=").Id("ag").Dot("internalDirectInsert").Call(
// 				jen.Id("tx"),
// 				jen.Id(lowerCamel),
// 				jen.Id("fn"),
// 			),
// 		),
// 		jen.Return().Id("ids"),
// 	)
// }
//
// func genFuncDeleteByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	return jen.Line().Comment("DeleteByID 根据 ID 删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByID").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id("id").Id("int64"),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
// 			jen.Id("tx"),
// 			jen.Index().Id("int64").Values(jen.Id("id")),
// 		),
// 	)
// }
//
// func genFuncDeleteByIDs(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	return jen.Line().Comment("DeleteByIDs 根据 ID 列表删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByIDs").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id("ids").Op("...").Id("int64"),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
// 			jen.Id("tx"),
// 			jen.Id("ids"),
// 		),
// 	)
// }
//
// func genFuncBatchDeleteByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	return jen.Line().Comment("BatchDeleteByID 根据 ID 批量删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchDeleteByID").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id("ids").Index().Id("int64"),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
// 			jen.Lit("删除 ID 列表: %+v 的数据"),
// 			jen.Id("ids"),
// 		),
// 		jen.For(
// 			jen.List(
// 				jen.Id("_"),
// 				jen.Id("id"),
// 			).Op(":=").Range().Id("ids"),
// 		).Block(
// 			jen.Id("ds").Op(":=").Id("ag").Dot("internalDirectDelete").Call(
// 				jen.Id("tx"),
// 				jen.Id("id"),
// 			),
// 			jen.If(jen.Op("!").Id("ds")).Block(jen.Id("panic").Call(jen.Lit("存在数据删除错误"))),
// 		),
// 		jen.Return().Id("true"),
// 	)
// }
//
// func genFuncUpdateByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Line().Comment("UpdateByID 根据 ID 批量更新").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateByID").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
// 			jen.Id("tx"),
// 			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
// 			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
// 		),
// 	)
// }
//
// func genFuncUpdateNonNilByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Line().Comment("UpdateNonNilByID 根据 ID 更新非空字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateNonNilByID").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
// 			jen.Id("tx"),
// 			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
// 			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
// 		),
// 	)
// }
//
// func genFuncUpdateWithFuncByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Line().Comment("UpdateWithFuncByID 根据 ID 更新满足函数的字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateWithFuncByID").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		genDeclAnonymousFunc(),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
// 			jen.Id("tx"),
// 			jen.Index().Op("*").Add(useEntity(camel)).Values(jen.Id(lowerCamel)),
// 			jen.Id("fn"),
// 		),
// 	)
// }
//
// func genFuncBatchUpdateWithFuncByID(table string, columns []Column) jen.Code {
//   if !hasColumn(p_key, columns) {
//     return jen.Null()
//   }
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	var sql0, sql1 string
// 	sql0 = fmt.Sprintf("UPDATE %s SET ", table)
// 	if hasColumn(d_key, columns) {
// 		sql1 = fmt.Sprintf(" WHERE %s = ? AND %s;", p_key, ud_cond_key)
// 	} else {
// 		sql1 = fmt.Sprintf(" WHERE %s = ?;", p_key)
// 	}
//
// 	return jen.Line().Comment("BatchUpdateWithFuncByID 根据 ID 批量更新满足函数的字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchUpdateWithFuncByID").Params(
// 		jen.Id("tx").Op("*").Add(useSql("Tx")),
// 		jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 		genDeclAnonymousFunc(),
// 	).Params(jen.Id("bool")).Block(
// 		jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(jen.Lit("批量更新列表数据")),
// 		jen.For(
// 			jen.List(
// 				jen.Id("_"),
// 				jen.Id(lowerCamel),
// 			).Op(":=").Range().Id(lowerCamel+"s"),
// 		).Block(
// 			jen.If(jen.Id(lowerCamel).Dot("ID").Op("==").Id("nil")).Block(jen.Id("panic").Call(jen.Lit("ID 字段不能为空"))),
// 			jen.Id("id").Op(":=").Op("*").Id(lowerCamel).Dot("ID"),
// 			jen.List(
// 				jen.Id("fields"),
// 				jen.Id("values"),
// 			).Op(":=").Id("calcUpdateField").Call(
// 				jen.Id(lowerCamel),
// 				jen.Id("fn"),
// 			),
// 			jen.Var().Id("sqlBuilder").Add(useStrings("Builder")),
// 			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql0)),
// 			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
// 			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql1)),
// 			jen.Id("values").Op("=").Id("append").Call(
// 				jen.Id("values"),
// 				jen.Id("id"),
// 			),
// 			jen.Id("errorHandler").Op(":=").Add(useUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
// 			jen.List(
// 				jen.Id("stmt"),
// 				jen.Id("err"),
// 			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
// 			jen.Id("errorHandler").Call(jen.Id("err")),
// 			jen.List(
// 				jen.Id("result"),
// 				jen.Id("err"),
// 			).Op(":=").Id("stmt").Dot("ExecContext").Call(
// 				jen.Id("ag").Dot("getDbCtx").Call(),
// 				jen.Id("values").Op("..."),
// 			),
// 			jen.Id("errorHandler").Call(jen.Id("err")),
// 			jen.List(
// 				jen.Id("af"),
// 				jen.Id("err"),
// 			).Op(":=").Id("result").Dot("RowsAffected").Call(),
// 			jen.Id("errorHandler").Call(jen.Id("err")),
// 			jen.If(jen.Id("af").Op("!=").Lit(1)).Block(jen.Id("panic").Call(jen.Lit("更新错误"))),
// 			jen.Id("err").Op("=").Id("stmt").Dot("Close").Call(),
// 		),
// 		jen.Return().Id("true"),
// 	)
// }
//
// func genTreeRepositoryFile(table string, columns []Column) *jen.File {
// 	ret := jen.NewFile(table)
//
// 	ret.Add(genInterfaceAutoGen(table, columns))
// 	ret.Add(genStructAutoGen())
// 	ret.Add(genFuncGetDbCtx())
//
// 	ret.Add(genFuncMapperAll(table, columns))
// 	ret.Add(genFuncMapperNumeric())
//
// 	ret.Add(genFuncCalcInsertField(table, columns))
// 	ret.Add(genFuncCalcUpdateField(table, columns))
//
// 	ret.Add(genFuncInternalSelectByIDs(table, columns))
//   ret.Add(genFuncInternalDirectInsert(table, columns))
// 	ret.Add(genFuncInternalDirectDelete(table, columns))
//
// 	ret.Add(genFuncSelectByID(table, columns))
// 	ret.Add(genFuncSelectByIDs(table, columns))
// 	ret.Add(genFuncBatchSelectByID(table, columns))
// 	ret.Add(genFuncSelectByName(table, columns))
//
// 	ret.Add(genFuncInsert(table, columns))
// 	ret.Add(genFuncInsertNonNil(table, columns))
// 	ret.Add(genFuncInsertWithFunc(table, columns))
// 	ret.Add(genFuncBatchInsert(table, columns))
// 	ret.Add(genFuncBatchInsertWithFunc(table, columns))
//
// 	ret.Add(genFuncDeleteByID(table, columns))
// 	ret.Add(genFuncDeleteByIDs(table, columns))
// 	ret.Add(genFuncBatchDeleteByID(table, columns))
//
// 	ret.Add(genFuncUpdateByID(table, columns))
// 	ret.Add(genFuncUpdateNonNilByID(table, columns))
// 	ret.Add(genFuncUpdateWithFuncByID(table, columns))
// 	ret.Add(genFuncBatchUpdateWithFuncByID(table, columns))
// 	return ret
// }
//
// func getColumns(table string) []Column {
// 	var columns []Column
//
// 	columns = append(
// 		columns, Column{
// 			ColumnName:      "id",
// 			Type:            "bigint",
// 			Nullable:        "NO",
// 			TableName:       table,
// 			ColumnComment:   "主键",
// 			Tag:             "id",
// 			MaxLength:       0,
// 			NumberPrecision: 19,
// 			ColumnType:      "bigint",
// 			ColumnKey:       "PRI",
// 			Default:         "",
// 		},
// 	)
// 	columns = append(
// 		columns, Column{
// 			ColumnName:      "title",
// 			Type:            "varchar",
// 			Nullable:        "NO",
// 			TableName:       table,
// 			ColumnComment:   "主题",
// 			Tag:             "title",
// 			MaxLength:       255,
// 			NumberPrecision: 0,
// 			ColumnType:      "varchar(1000)",
// 			ColumnKey:       "",
// 			Default:         "",
// 		},
// 	)
// 	columns = append(
// 		columns, Column{
// 			ColumnName:      "start_at",
// 			Type:            "timestamp",
// 			Nullable:        "NO",
// 			TableName:       table,
// 			ColumnComment:   "开始时间",
// 			Tag:             "startAt",
// 			MaxLength:       0,
// 			NumberPrecision: 0,
// 			ColumnType:      "timestamp",
// 			ColumnKey:       "",
// 			Default:         "CURRENT_TIMESTAMP",
// 		},
// 	)
//
// 	// columns := make([]Column, 0)
// 	// var params = make(map[string]string)
// 	// params["parseTime"] = "true"
// 	// cfg := mysql.Config{
// 	//	User:   "root",
// 	//	Passwd: "root",
// 	//	Net:    "tcp",
// 	//	Addr:   "localhost:3307",
// 	//	DBName: "metis",
// 	//	Params: params,
// 	// }
// 	//
// 	// var err error
// 	// db, err := sql.Open("mysql", cfg.FormatDSN())
// 	// rows, err := db.Query(
// 	//	fmt.Sprintf(
// 	//		`SELECT
// 	//	COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT,CHARACTER_MAXIMUM_LENGTH,COLUMN_TYPE,NUMERIC_PRECISION,COLUMN_KEY,COLUMN_DEFAULT
// 	//	FROM information_schema.COLUMNS
// 	//	WHERE table_schema = DATABASE()  AND TABLE_NAME = '%s'`, table,
// 	//	),
// 	// )
// 	//
// 	// if err != nil {
// 	//	log.Printf("table rows is nil with table:%s error: %v \n", table, err)
// 	//	return columns
// 	// }
// 	//
// 	// if rows == nil {
// 	//	log.Printf("rows is nil with table:%s \n", table)
// 	//	return columns
// 	// }
// 	//
// 	// defer func() {
// 	//	_ = rows.Close()
// 	// }()
// 	//
// 	// for rows.Next() {
// 	//
// 	//	// todo: mysql bigint => go []byte
// 	//	var maxLength, numberPrecision []byte
// 	//	var t = ""
// 	//
// 	//	col := Column{}
// 	//	err = rows.Scan(
// 	//		&col.ColumnName, &col.Type, &t, &col.TableName, &col.ColumnComment, &maxLength, &col.ColumnType, &numberPrecision,
// 	//		&col.ColumnKey, &col.Default,
// 	//	)
// 	//	col.Nullable = t
// 	//	col.Tag = col.ColumnName
// 	//
// 	//	if maxLength != nil {
// 	//		col.MaxLength = Byte2Int64(maxLength)
// 	//	}
// 	//
// 	//	if numberPrecision != nil {
// 	//		col.NumberPrecision = Byte2Int64(numberPrecision)
// 	//	}
// 	//
// 	//	if err != nil {
// 	//		log.Println(err.Error())
// 	//		continue
// 	//	}
// 	//
// 	//	columns = append(columns, col)
// 	// }
//
// 	return columns
// }
//
// func TestRepository(t *testing.T) {
//
// 	strcase.ConfigureAcronym("ID", "id")
// 	strcase.ConfigureAcronym("id", "ID")
//
// 	table := "rroollee"
// 	columns := getColumns(table)
//
// 	f := genTreeRepositoryFile(table, columns)
// 	fmt.Printf("%#v\n", f)
// 	autogenFilePath := "second/module/user/repository/" + table + "/autogen.go"
// 	if err := os.MkdirAll(filepath.Dir(autogenFilePath), 0766); err != nil {
// 		panic(err)
// 	}
// 	wr, err := os.OpenFile(autogenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = f.Render(wr)
// }
