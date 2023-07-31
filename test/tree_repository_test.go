package test

//
// import (
// 	"fmt"
// 	"github.com/dave/jennifer/jen"
// 	"github.com/iancoleman/strcase"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"testing"
// 	"time"
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
// func genDeclFuncMapperAll(table string, columns []Column) jen.Code {
// 	var codes = make([]jen.Code, len(columns))
// 	for i, column := range columns {
// 		codes[i] = renderAndField("r", column.ColumnName)
// 	}
//
// 	structName := strcase.ToCamel(table)
//
// 	return jen.Func().Id("mapperAll").Params().
// 		Params(jen.Op("*").Add(useEntity(structName)), jen.Index().Id("any")).
// 		Block(
// 			jen.Null().Var().Id("r").Op("=").Op("&").Add(useEntity(structName)).Values(),
// 			jen.Null().Var().Id("cs").Op("=").Index().Id("any").Values(
// 				codes...,
// 			),
// 			jen.Return().List(jen.Id("r"), jen.Id("cs")),
// 		)
// }
//
// func genDeclInterfaceAutoGen(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Type().Id("IAutoGen").Interface(
// 		inferColumn(
// 			jen.Id("SelectByID").Params(jen.Id("id").Id("int64")).Params(useEntity(camel)), "id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).Params(jen.Index().Add(useEntity(camel))),
// 			"id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).Params(jen.Index().Add(useEntity(camel))),
// 			"id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Add(useEntity(camel))),
// 			"name", columns,
// 		),
// 		jen.Line(),
// 		jen.Id("Insert").Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		).Params(jen.Id("int64")),
// 		jen.Id("BatchInsert").Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 		).Params(jen.Index().Id("int64")),
// 		jen.Id("InsertNonNil").Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		).Params(jen.Id("int64")),
// 		jen.Id("InsertWithFunc").Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 			genDeclAnonymousFunc(),
// 		).Params(jen.Id("int64")),
// 		jen.Id("BatchInsertWithFunc").Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 			genDeclAnonymousFunc(),
// 		).Params(jen.Index().Id("int64")),
// 		jen.Line(),
// 		inferColumn(
// 			jen.Id("DeleteByID").Params(
// 				jen.Id("tx").Op("*").Add(useSql("Tx")),
// 				jen.Id("id").Id("int64"),
// 			).Params(jen.Id("bool")), "id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("DeleteByIDs").Params(
// 				jen.Id("tx").Op("*").Add(useSql("Tx")),
// 				jen.Id("ids").Op("...").Id("int64"),
// 			).Params(jen.Id("bool")), "id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("BatchDeleteByID").Params(
// 				jen.Id("tx").Op("*").Add(useSql("Tx")),
// 				jen.Id("ids").Index().Id("int64"),
// 			).Params(jen.Id("bool")), "id", columns,
// 		),
// 		jen.Line(),
// 		inferColumn(
// 			jen.Id("UpdateByID").Params(
// 				jen.Id("tx").Op("*").Add(useSql("Tx")),
// 				jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 			).Params(jen.Id("bool")), "id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("UpdateNonNilByID").Params(
// 				jen.Id("tx").Op("*").Add(useSql("Tx")),
// 				jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 			).Params(jen.Id("bool")), "id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("UpdateWithFuncByID").Params(
// 				jen.Id("tx").Op("*").Add(useSql("Tx")),
// 				jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 				genDeclAnonymousFunc(),
// 			).Params(jen.Id("bool")), "id", columns,
// 		),
// 		inferColumn(
// 			jen.Id("BatchUpdateWithFuncByID").Params(
// 				jen.Id("tx").Op("*").Add(useSql("Tx")),
// 				jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 				genDeclAnonymousFunc(),
// 			).Params(jen.Id("bool")), "id", columns,
// 		),
// 	)
// }
//
// func genDeclStructAutoGen(table string, columns []Column) jen.Code {
// 	return jen.Null().Type().Id("autoGen").Struct(jen.Id("ctx").Op("*").Add(useGin("Context")))
// }
//
// func genDeclFuncGetDbCtx(table string, columns []Column) jen.Code {
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("getDbCtx").
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
// func genDeclFuncSelectByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	var sql strings.Builder
// 	var fields = make([]string, len(columns))
// 	sql.WriteString("SELECT ")
//
// 	for i, column := range columns {
// 		fields[i] = column.ColumnName
// 	}
// 	sql.WriteString(strings.Join(fields, ", "))
// 	sql.WriteString(" FROM ")
// 	sql.WriteString(table)
// 	sql.WriteString(" WHERE ")
// 	sql.WriteString("id = ?")
// 	if hasColumn("deleted", columns) {
// 		sql.WriteString(" AND deleted = 0")
// 	}
// 	sql.WriteString(";")
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).
// 		Id("SelectByID").Params(jen.Id("id").Id("int64")).
// 		Params(useEntity(camel)).
// 		Block(
// 			jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("db").Op(":=").Add(useDatabase("FetchDB")).Call(),
// 			jen.Id("sqlPlaceholder").Op(":=").Lit(sql.String()),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("_"),
// 			).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.Id("row").Op(":=").Id("prepare").Dot("QueryRowContext").Call(
// 				jen.Id("ag").Dot("getDbCtx").Call(),
// 				jen.Id("id"),
// 			),
// 			jen.Id(lowerCamel).Op(":=").Add(useUtil("Row")).Call(
// 				jen.Id("row"),
// 				jen.Id("mapperAll"),
// 			),
// 			jen.Return().Id(lowerCamel),
// 		)
// }
//
// func genDeclFuncSelectByIDs(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByIDs").
// 		Params(jen.Id("ids").Op("...").Id("int64")).
// 		Params(jen.Index().Add(useEntity(strcase.ToCamel(table)))).
// 		Block(jen.Return().Id("ag").Dot("BatchSelectByID").Call(jen.Id("ids")))
// }
//
// func genDeclFuncBatchSelectByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	var sql strings.Builder
// 	var fields = make([]string, len(columns))
// 	sql.WriteString("SELECT ")
//
// 	for i, column := range columns {
// 		fields[i] = column.ColumnName
// 	}
// 	sql.WriteString(strings.Join(fields, ", "))
// 	sql.WriteString(" FROM ")
// 	sql.WriteString(table)
// 	sql.WriteString(" WHERE ")
// 	sql.WriteString("id = (%s)")
// 	if hasColumn("deleted", columns) {
// 		sql.WriteString(" AND deleted = 0")
// 	}
// 	sql.WriteString(";")
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchSelectByID").
// 		Params(jen.Id("ids").Index().Id("int64")).
// 		Params(jen.Index().Add(useEntity(camel))).
// 		Block(
// 			jen.Id("placeholder").Op(":=").Id("make").Call(
// 				jen.Index().Id("string"),
// 				jen.Id("len").Call(jen.Id("ids")),
// 			),
// 			jen.For(
// 				jen.Id("i").Op(":=").Lit(0),
// 				jen.Id("i").Op("<").Id("len").Call(jen.Id("ids")),
// 				jen.Id("i").Op("++"),
// 			).Block(jen.Id("placeholder").Index(jen.Id("i")).Op("=").Lit("?")),
// 			jen.Id("recorder").Op(":=").Add(useLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("db").Op(":=").Add(useDatabase("FetchDB")).Call(),
// 			jen.Id("sqlPlaceholder").Op(":=").Add(useFmt("Sprintf")).Call(
// 				jen.Lit(sql.String()),
// 				jen.Qual("strings", "Join").Call(
// 					jen.Id("placeholder"),
// 					jen.Lit(", "),
// 				),
// 			),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("_"),
// 			).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.Id("bindValues").Op(":=").Id("make").Call(
// 				jen.Index().Id("any"),
// 				jen.Id("len").Call(jen.Id("ids")),
// 			),
// 			jen.For(
// 				jen.List(
// 					jen.Id("i"),
// 					jen.Id("id"),
// 				).Op(":=").Range().Id("ids"),
// 			).Block(jen.Id("bindValues").Index(jen.Id("i")).Op("=").Id("id")),
// 			jen.List(
// 				jen.Id("rows"),
// 				jen.Id("err"),
// 			).Op(":=").Id("prepare").Dot("QueryContext").Call(
// 				jen.Id("ag").Dot("getDbCtx").Call(),
// 				jen.Id("bindValues").Op("..."),
// 			),
// 			jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
// 				jen.Id("recorder").Dot("Error").Call(
// 					jen.Id("err").Dot("Error").Call(),
// 					jen.Add(useZap("Error")).Call(jen.Id("err")),
// 				),
// 			),
// 			jen.Id(lowerCamel+"s").Op(":=").Add(useUtil("Rows")).Call(
// 				jen.Id("rows"),
// 				jen.Id("mapperAll"),
// 			),
// 			jen.Return().Id(lowerCamel+"s"),
// 		)
// }
//
// func genDeclFuncSelectByName(table string, columns []Column) jen.Code {
// 	if !hasColumn("name", columns) {
// 		return jen.Null()
// 	}
//
// 	var sql strings.Builder
// 	var fields = make([]string, len(columns))
// 	sql.WriteString("SELECT ")
//
// 	for i, column := range columns {
// 		fields[i] = column.ColumnName
// 	}
// 	sql.WriteString(strings.Join(fields, ", "))
// 	sql.WriteString(" FROM ")
// 	sql.WriteString(table)
// 	sql.WriteString(" WHERE ")
// 	sql.WriteString("name LIKE ?")
// 	if hasColumn("deleted", columns) {
// 		sql.WriteString(" AND deleted = 0")
// 	}
// 	sql.WriteString(";")
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByName").
// 		Params(jen.Id("name").Id("string")).
// 		Params(jen.Index().Add(useEntity(camel))).
// 		Block(
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("db").Op(":=").Add(useDatabase("FetchDB")).Call(),
// 			jen.Id("sqlPlaceholder").Op(":=").Lit(sql.String()),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("_"),
// 			).Op(":=").Id("db").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.List(
// 				jen.Id("rows"),
// 				jen.Id("err"),
// 			).Op(":=").Id("prepare").Dot("QueryContext").Call(
// 				jen.Id("ag").Dot("getDbCtx").Call(),
// 				jen.Id("name"),
// 			),
// 			jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
// 				jen.Id("recorder").Dot("Error").Call(
// 					jen.Id("err").Dot("Error").Call(),
// 					jen.Add(useZap("Error")).Call(jen.Id("err")),
// 				),
// 			),
// 			jen.Id(lowerCamel+"s").Op(":=").Add(useUtil("Rows")).Call(
// 				jen.Id("rows"),
// 				jen.Id("mapperAll"),
// 			),
// 			jen.Return().Id(lowerCamel+"s"),
// 		)
// }
//
// func genDeclFuncInternalInsert(table string, columns []Column) jen.Code {
// 	structName := strcase.ToCamel(table)
// 	varName := strcase.ToLowerCamel(table)
// 	codes := make([]jen.Code, len(columns)+1)
// 	codes[0] = jen.Id("ag").Dot("getDbCtx").Call()
// 	for i, column := range columns {
// 		codes[i+1] = renderStarField(varName, column.ColumnName)
// 	}
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalInsert").
// 		Params(
// 			jen.Id("prepare").Op("*").Add(useSql("Stmt")),
// 			jen.Id(varName).Op("*").Add(useEntity(structName)),
// 		).
// 		Params(jen.Id("int64")).
// 		Block(
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.List(
// 				jen.Id("result"),
// 				jen.Id("err"),
// 			).Op(":=").Id("prepare").Dot("ExecContext").Call(
// 				codes...,
// 			),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.List(
// 				jen.Id("id"),
// 				jen.Id("err"),
// 			).Op(":=").Id("result").Dot("LastInsertId").Call(),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.Return().Id("id"),
// 		)
// }
//
// func genDeclFuncInsert(table string, columns []Column) jen.Code {
// 	var sql strings.Builder
// 	var fields = make([]string, len(columns))
// 	var ph = make([]string, len(columns))
// 	sql.WriteString("INSERT ")
// 	sql.WriteString("INTO ")
// 	sql.WriteString(table)
//
// 	for i, column := range columns {
// 		fields[i] = column.ColumnName
// 		ph[i] = "?"
// 	}
// 	sql.WriteString("(")
// 	sql.WriteString(strings.Join(fields, ", "))
// 	sql.WriteString(")")
// 	sql.WriteString(" VALUES (")
// 	sql.WriteString(strings.Join(ph, ", "))
// 	sql.WriteString(");")
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("Insert").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		).
// 		Params(jen.Id("int64")).
// 		Block(
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("sqlPlaceholder").Op(":=").Lit(sql.String()),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("err"),
// 			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.Return().Id("ag").Dot("internalInsert").Call(
// 				jen.Id("prepare"),
// 				jen.Id(lowerCamel),
// 			),
// 		)
// }
//
// func genDeclFuncBatchInsert(table string, columns []Column) jen.Code {
// 	var sql strings.Builder
// 	var fields = make([]string, len(columns))
// 	var ph = make([]string, len(columns))
// 	sql.WriteString("INSERT ")
// 	sql.WriteString("INTO ")
// 	sql.WriteString(table)
//
// 	for i, column := range columns {
// 		fields[i] = column.ColumnName
// 		ph[i] = "?"
// 	}
// 	sql.WriteString("(")
// 	sql.WriteString(strings.Join(fields, ", "))
// 	sql.WriteString(")")
// 	sql.WriteString(" VALUES (")
// 	sql.WriteString(strings.Join(ph, ", "))
// 	sql.WriteString(");")
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsert").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 		).
// 		Params(jen.Index().Id("int64")).
// 		Block(
// 			jen.Id("retids").Op(":=").Id("make").Call(
// 				jen.Index().Id("int64"),
// 				jen.Id("len").Call(jen.Id(lowerCamel+"s")),
// 			),
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("sqlPlaceholder").Op(":=").Lit(sql.String()),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("err"),
// 			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.For(
// 				jen.List(
// 					jen.Id("i"),
// 					jen.Id(lowerCamel),
// 				).Op(":=").Range().Id(lowerCamel+"s"),
// 			).Block(
// 				jen.Id("retids").Index(jen.Id("i")).Op("=").Id("ag").Dot("internalInsert").Call(
// 					jen.Id("prepare"),
// 					jen.Id(lowerCamel),
// 				),
// 			),
// 			jen.Return().Id("retids"),
// 		)
// }
//
// func genDeclFuncInsertNonNil(table string, columns []Column) jen.Code {
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNil").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		).
// 		Params(jen.Id("int64")).
// 		Block(
// 			jen.Return().Id("ag").Dot("InsertWithFunc").Call(
// 				jen.Id("tx"),
// 				jen.Id(lowerCamel),
// 				jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
// 			),
// 		)
// }
//
// func genDeclFuncInsertWithFunc(table string, columns []Column) jen.Code {
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	var codes []jen.Code
//
// 	var sql strings.Builder
// 	sql.WriteString("INSERT ")
// 	sql.WriteString("INTO ")
// 	sql.WriteString(table)
//
// 	codes = append(
// 		codes, jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 		jen.Null().Var().Id("needField").Qual("strings", "Builder"),
// 		jen.Null().Var().Id("needPlace").Qual("strings", "Builder"),
// 		jen.Null().Var().Id("bindValue").Index().Id("any"),
// 	)
//
// 	for _, column := range columns {
// 		codes = append(
// 			codes, jen.If(jen.Id("fn").Call(jen.Id(lowerCamel).Dot(strcase.ToCamel(column.ColumnName)))).Block(
// 				jen.Id("needField").Dot("WriteString").Call(jen.Lit(column.ColumnName+", ")),
// 				jen.Id("needPlace").Dot("WriteString").Call(jen.Lit("?, ")),
// 				jen.Id("bindValue").Op("=").Id("append").Call(
// 					jen.Id("bindValue"),
// 					jen.Op("*").Id(lowerCamel).Dot(strcase.ToCamel(column.ColumnName)),
// 				),
// 			),
// 		)
// 	}
//
// 	sql.WriteString("(%s)")
// 	sql.WriteString(" VALUES ")
// 	sql.WriteString("(%s);")
//
// 	codes = append(
// 		codes, jen.Id("sqlPlaceholder").Op(":=").Qual("fmt", "Sprintf").Call(
// 			jen.Lit(sql.String()),
// 			jen.Id("needField").Dot("String").Call().Index(
// 				jen.Empty(),
// 				jen.Id("needField").Dot("Len").Call().Op("-").Lit(2),
// 			),
// 			jen.Id("needPlace").Dot("String").Call().Index(
// 				jen.Empty(),
// 				jen.Id("needPlace").Dot("Len").Call().Op("-").Lit(2),
// 			),
// 		),
// 		jen.List(
// 			jen.Id("prepare"),
// 			jen.Id("err"),
// 		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 		jen.Defer().Add(useUtil("DeferClose")).Call(
// 			jen.Id("prepare"),
// 			jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 		),
// 		jen.Add(useUtil("PanicErr")).Call(
// 			jen.Id("recorder"),
// 			jen.Id("err"),
// 		),
// 		jen.Return().Id("ag").Dot("internalInsert").Call(
// 			jen.Id("prepare"),
// 			jen.Id(lowerCamel),
// 		),
// 	)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFunc").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 			genDeclAnonymousFunc(),
// 		).
// 		Params(jen.Id("int64")).
// 		Block(codes...)
// }
//
// func genDeclFuncBatchInsertWithFunc(table string, columns []Column) jen.Code {
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertWithFunc").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 			genDeclAnonymousFunc(),
// 		).
// 		Params(jen.Index().Id("int64")).
// 		Block(
// 			jen.Id("retids").Op(":=").Id("make").Call(
// 				jen.Index().Id("int64"),
// 				jen.Id("len").Call(jen.Id(lowerCamel+"s")),
// 			),
// 			jen.For(
// 				jen.List(
// 					jen.Id("i"),
// 					jen.Id(lowerCamel),
// 				).Op(":=").Range().Id(lowerCamel+"s"),
// 			).Block(
// 				jen.Id("retids").Index(jen.Id("i")).Op("=").Id("ag").Dot("InsertWithFunc").Call(
// 					jen.Id("tx"),
// 					jen.Id(lowerCamel),
// 					jen.Id("fn"),
// 				),
// 			),
// 			jen.Return().Id("retids"),
// 		)
// }
//
// func genDeclFuncDeleteByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	var sql strings.Builder
// 	sql.WriteString("UPDATE ")
// 	sql.WriteString(table)
// 	sql.WriteString("SET deleted = 1 ")
// 	sql.WriteString("WHERE ")
// 	sql.WriteString("id = ?")
// 	if hasColumn("deleted", columns) {
// 		sql.WriteString(" AND deleted = 0")
// 	}
// 	sql.WriteString(";")
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByID").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id("id").Id("int64"),
// 		).
// 		Params(jen.Id("bool")).
// 		Block(
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("sqlPlaceholder").Op(":=").Lit(sql.String()),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("err"),
// 			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.List(
// 				jen.Id("result"),
// 				jen.Id("err"),
// 			).Op(":=").Id("prepare").Dot("ExecContext").Call(
// 				jen.Id("ag").Dot("getDbCtx").Call(),
// 				jen.Id("id"),
// 			),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.List(
// 				jen.Id("af"),
// 				jen.Id("err"),
// 			).Op(":=").Id("result").Dot("RowsAffected").Call(),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.Return().Id("af").Op("==").Lit(1),
// 		)
// }
//
// func genDeclFuncDeleteByIDs(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByIDs").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id("ids").Op("...").Id("int64"),
// 		).
// 		Params(jen.Id("bool")).
// 		Block(
// 			jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
// 				jen.Id("tx"),
// 				jen.Id("ids"),
// 			),
// 		)
// }
//
// func genDeclFuncBatchDeleteByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
//
// 	var sql strings.Builder
// 	sql.WriteString("UPDATE ")
// 	sql.WriteString(table)
// 	sql.WriteString("SET deleted = 1 ")
// 	sql.WriteString("WHERE ")
// 	sql.WriteString("id IN (%s)")
// 	if hasColumn("deleted", columns) {
// 		sql.WriteString(" AND deleted = 0")
// 	}
// 	sql.WriteString(";")
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchDeleteByID").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id("ids").Index().Id("int64"),
// 		).
// 		Params(jen.Id("bool")).
// 		Block(
// 			jen.Id("placeholder").Op(":=").Id("make").Call(
// 				jen.Index().Id("string"),
// 				jen.Id("len").Call(jen.Id("ids")),
// 			),
// 			jen.For(
// 				jen.Id("i").Op(":=").Lit(0),
// 				jen.Id("i").Op("<").Id("len").Call(jen.Id("ids")),
// 				jen.Id("i").Op("++"),
// 			).Block(jen.Id("placeholder").Index(jen.Id("i")).Op("=").Lit("?")),
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("sqlPlaceholder").Op(":=").Qual("fmt", "Sprintf").Call(
// 				jen.Lit(sql.String()),
// 				jen.Qual("strings", "Join").Call(
// 					jen.Id("placeholder"),
// 					jen.Lit(", "),
// 				),
// 			),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("_"),
// 			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.Id("bindValues").Op(":=").Id("make").Call(
// 				jen.Index().Id("any"),
// 				jen.Id("len").Call(jen.Id("ids")),
// 			),
// 			jen.For(
// 				jen.List(
// 					jen.Id("i"),
// 					jen.Id("id"),
// 				).Op(":=").Range().Id("ids"),
// 			).Block(jen.Id("bindValues").Index(jen.Id("i")).Op("=").Id("id")),
// 			jen.List(
// 				jen.Id("result"),
// 				jen.Id("err"),
// 			).Op(":=").Id("prepare").Dot("ExecContext").Call(
// 				jen.Id("ag").Dot("getDbCtx").Call(),
// 				jen.Id("bindValues").Op("..."),
// 			),
// 			jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
// 				jen.Id("recorder").Dot("Error").Call(
// 					jen.Id("err").Dot("Error").Call(),
// 					jen.Qual("go.uber.org/zap", "Error").Call(jen.Id("err")),
// 				),
// 			),
// 			jen.List(
// 				jen.Id("af"),
// 				jen.Id("err"),
// 			).Op(":=").Id("result").Dot("RowsAffected").Call(),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.Return().Id("af").Op("==").Id("int64").Call(jen.Id("len").Call(jen.Id("ids"))),
// 		)
// }
//
// func genDeclFuncUpdateByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
//
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	setFields := make([]string, len(columns))
// 	codes := make([]jen.Code, len(columns)+1)
// 	var sql strings.Builder
// 	sql.WriteString("UPDATE ")
// 	sql.WriteString(table)
// 	sql.WriteString(" SET ")
//
// 	codes[0] = jen.Id("ag").Dot("getDbCtx").Call()
// 	for i, column := range columns {
// 		setFields[i] = column.ColumnName + " = ?"
// 		codes[i+1] = renderStarField(lowerCamel, column.ColumnName)
// 	}
//
// 	sql.WriteString(strings.Join(setFields, ", "))
// 	sql.WriteString(" WHERE ")
// 	sql.WriteString("id IN ?")
// 	if hasColumn("deleted", columns) {
// 		sql.WriteString(" AND deleted = 0")
// 	}
// 	sql.WriteString(";")
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateByID").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		).
// 		Params(jen.Id("bool")).
// 		Block(
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("sqlPlaceholder").Op(":=").Lit(sql.String()),
// 			jen.List(
// 				jen.Id("prepare"),
// 				jen.Id("err"),
// 			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 			jen.Defer().Add(useUtil("DeferClose")).Call(
// 				jen.Id("prepare"),
// 				jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 			),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.List(
// 				jen.Id("result"),
// 				jen.Id("err"),
// 			).Op(":=").Id("prepare").Dot("ExecContext").Call(codes...),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.List(
// 				jen.Id("af"),
// 				jen.Id("err"),
// 			).Op(":=").Id("result").Dot("RowsAffected").Call(),
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Id("err"),
// 			),
// 			jen.Return().Id("af").Op("==").Lit(1),
// 		)
// }
//
// func genDeclFuncUpdateNonNilByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateNonNilByID").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 		).
// 		Params(jen.Id("bool")).
// 		Block(
// 			jen.Return().Id("ag").Dot("UpdateWithFuncByID").Call(
// 				jen.Id("tx"),
// 				jen.Id(lowerCamel),
// 				jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("==").Id("nil")),
// 			),
// 		)
// }
//
// func genDeclFuncUpdateWithFuncByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	var codes []jen.Code
//
// 	var sql strings.Builder
// 	sql.WriteString("UPDATE ")
// 	sql.WriteString(table)
// 	sql.WriteString(" SET ")
//
// 	codes = append(
// 		codes, jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 		jen.If(jen.Id(lowerCamel).Dot("ID").Op("==").Id("nil")).Block(
// 			jen.Add(useUtil("PanicErr")).Call(
// 				jen.Id("recorder"),
// 				jen.Add(useErrors("New")).Call(jen.Lit("ID is nil")),
// 			),
// 		),
// 		jen.Null().Var().Id("needFieldAndPlace").Qual("strings", "Builder"),
// 		jen.Null().Var().Id("bindValue").Index().Id("any"),
// 	)
//
// 	for _, column := range columns {
// 		if column.ColumnName == "id" {
// 			continue
// 		}
// 		codes = append(
// 			codes, jen.If(jen.Id("fn").Call(jen.Id(lowerCamel).Dot(strcase.ToCamel(column.ColumnName)))).Block(
// 				jen.Id("needFieldAndPlace").Dot("WriteString").Call(jen.Lit(column.ColumnName+" = ?, ")),
// 				jen.Id("bindValue").Op("=").Id("append").Call(
// 					jen.Id("bindValue"),
// 					jen.Op("*").Id(lowerCamel).Dot(strcase.ToCamel(column.ColumnName)),
// 				),
// 			),
// 		)
// 	}
//
// 	sql.WriteString("%s")
// 	sql.WriteString(" WHERE ")
// 	sql.WriteString("id = ?")
// 	if hasColumn("deleted", columns) {
// 		sql.WriteString(" AND deleted = 0")
// 	}
// 	sql.WriteString(";")
//
// 	codes = append(
// 		codes, jen.Id("bindValue").Op("=").Id("append").Call(
// 			jen.Id("bindValue"),
// 			jen.Op("*").Id(lowerCamel).Dot("ID"),
// 		),
// 		jen.Id("sqlPlaceholder").Op(":=").Qual("fmt", "Sprintf").Call(
// 			jen.Lit(sql.String()),
// 			jen.Id("needFieldAndPlace").Dot("String").Call().Index(
// 				jen.Empty(),
// 				jen.Id("needFieldAndPlace").Dot("Len").Call().Op("-").Lit(2),
// 			),
// 		),
// 		jen.List(
// 			jen.Id("prepare"),
// 			jen.Id("err"),
// 		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlPlaceholder")),
// 		jen.Defer().Add(useUtil("DeferClose")).Call(
// 			jen.Id("prepare"),
// 			jen.Add(useUtil("ErrToLog")).Call(jen.Id("recorder")),
// 		),
// 		jen.Add(useUtil("PanicErr")).Call(
// 			jen.Id("recorder"),
// 			jen.Id("err"),
// 		),
// 		jen.List(
// 			jen.Id("result"),
// 			jen.Id("err"),
// 		).Op(":=").Id("prepare").Dot("ExecContext").Call(
// 			jen.Id("ag").Dot("getDbCtx").Call(),
// 			jen.Id("bindValue").Op("..."),
// 		),
// 		jen.Add(useUtil("PanicErr")).Call(
// 			jen.Id("recorder"),
// 			jen.Id("err"),
// 		),
// 		jen.List(
// 			jen.Id("af"),
// 			jen.Id("err"),
// 		).Op(":=").Id("result").Dot("RowsAffected").Call(),
// 		jen.Add(useUtil("PanicErr")).Call(
// 			jen.Id("recorder"),
// 			jen.Id("err"),
// 		),
// 		jen.Return().Id("af").Op("==").Lit(1),
// 	)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateWithFuncByID").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel).Op("*").Add(useEntity(camel)),
// 			genDeclAnonymousFunc(),
// 		).
// 		Params(jen.Id("bool")).
// 		Block(codes...)
// }
//
// func genDeclFuncBatchUpdateWithFuncByID(table string, columns []Column) jen.Code {
// 	if !hasColumn("id", columns) {
// 		return jen.Null()
// 	}
// 	camel := strcase.ToCamel(table)
// 	lowerCamel := strcase.ToLowerCamel(table)
//
// 	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchUpdateWithFuncByID").
// 		Params(
// 			jen.Id("tx").Op("*").Add(useSql("Tx")),
// 			jen.Id(lowerCamel+"s").Index().Op("*").Add(useEntity(camel)),
// 			genDeclAnonymousFunc(),
// 		).
// 		Params(jen.Id("bool")).
// 		Block(
// 			jen.Id("recorder").Op(":=").Id("logger").Dot("AccessLogger").Call(jen.Id("ag").Dot("ctx")),
// 			jen.Id("af").Op(":=").Id("false"),
// 			jen.For(
// 				jen.List(
// 					jen.Id("_"),
// 					jen.Id(lowerCamel),
// 				).Op(":=").Range().Id(lowerCamel+"s"),
// 			).Block(
// 				jen.Id("af").Op("=").Id("ag").Dot("UpdateWithFuncByID").Call(
// 					jen.Id("tx"),
// 					jen.Id(lowerCamel),
// 					jen.Id("fn"),
// 				),
// 				jen.If(jen.Id("af")).Block(jen.Continue()),
// 				jen.Add(useUtil("PanicErr")).Call(
// 					jen.Id("recorder"),
// 					jen.Add(useErrors("New")).Call(jen.Lit("some entry update failed")),
// 				),
// 				jen.Break(),
// 			),
// 			jen.Return().Id("af"),
// 		)
// }
//
// func genRepositoryFile(table string, columns []Column) *jen.File {
// 	f := jen.NewFile(table)
// 	f.ImportName("metis/test/first/entity", "entity")
// 	f.ImportName("context", "context")
// 	f.ImportName("database/sql", "sql")
// 	f.ImportName("errors", "errors")
// 	f.ImportName("fmt", "fmt")
// 	f.ImportName("time", "time")
// 	f.ImportName("github.com/jinzhu/copier", "copier")
// 	f.ImportName("github.com/gin-gonic/gin", "gin")
// 	f.ImportName("go.uber.org/zap", "zap")
// 	f.ImportName("metis/config/constant", "constant")
// 	f.ImportName("metis/database", "database")
// 	f.ImportName("metis/test/first/model/dto", "dto")
// 	f.ImportName("metis/util", "util")
// 	f.ImportName("metis/util/logger", "logger")
// 	f.ImportName("strings", "strings")
//
// 	f.HeaderComment("Code generated by tabuyos. DO NOT EDIT!")
// 	f.PackageComment("Package " + table)
// 	f.PackageComment("@author tabuyos")
// 	f.PackageComment("@since " + time.Now().Format("2006/01/02"))
// 	f.PackageComment("@description " + table)
//
// 	f.Add(genDeclInterfaceAutoGen(table, columns))
//
// 	f.Add(genDeclStructAutoGen(table, columns))
// 	f.Add(genDeclFuncGetDbCtx(table, columns))
// 	f.Add(genDeclFuncSelectByID(table, columns))
// 	f.Add(genDeclFuncSelectByIDs(table, columns))
// 	f.Add(genDeclFuncBatchSelectByID(table, columns))
// 	f.Add(genDeclFuncSelectByName(table, columns))
// 	f.Add(genDeclFuncInternalInsert(table, columns))
// 	f.Add(genDeclFuncInsert(table, columns))
// 	f.Add(genDeclFuncBatchInsert(table, columns))
// 	f.Add(genDeclFuncInsertNonNil(table, columns))
// 	f.Add(genDeclFuncInsertWithFunc(table, columns))
// 	f.Add(genDeclFuncBatchInsertWithFunc(table, columns))
// 	f.Add(genDeclFuncDeleteByID(table, columns))
// 	f.Add(genDeclFuncDeleteByIDs(table, columns))
// 	f.Add(genDeclFuncBatchDeleteByID(table, columns))
// 	f.Add(genDeclFuncUpdateByID(table, columns))
// 	f.Add(genDeclFuncUpdateNonNilByID(table, columns))
// 	f.Add(genDeclFuncUpdateWithFuncByID(table, columns))
// 	f.Add(genDeclFuncBatchUpdateWithFuncByID(table, columns))
//
// 	f.Add(genDeclFuncMapperAll(table, columns))
// 	return f
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
// 	return columns
// }
//
// func TestRepository(t *testing.T) {
// 	strcase.ConfigureAcronym("ID", "id")
// 	strcase.ConfigureAcronym("id", "ID")
//
// 	table := "survey"
// 	columns := getColumns(table)
//
// 	f := genRepositoryFile(table, columns)
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
