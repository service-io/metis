// Package repository
// @author tabuyos
// @since 2023/8/1
// @description repository
package repository

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"metis/generated/helper"
	"time"
)

type GenBaseRepo struct {
}

func (rec *GenBaseRepo) genInterfaceAutoGen(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Comment("iAutoGen 该接口自动生成, 请勿修改").Line().Type().Id("iAutoGen").Interface(
		helper.InferColumn(
			jen.Id("SelectByID").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(helper.UseEntity(camel))), "id",
			columns,
		),
		helper.InferColumn(
			jen.Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).Params(jen.Index().Op("*").Add(helper.UseEntity(camel))),
			"id", columns,
		),
		helper.InferColumn(
			jen.Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).Params(jen.Index().Op("*").Add(helper.UseEntity(camel))),
			"id", columns,
		),
		jen.Line(),
		helper.InferColumn(
			jen.Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(helper.UseEntity(camel))),
			"name", columns,
		),
		jen.Line(),
		jen.Id("Insert").Params(
			jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
		).Params(jen.Id("int64")),
		jen.Id("InsertNonNil").Params(
			jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
		).Params(jen.Id("int64")),
		jen.Id("InsertWithFunc").Params(
			jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
			helper.GenDeclAnonymousFunc(),
		).Params(jen.Id("int64")),
		jen.Id("BatchInsert").Params(
			jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
			jen.Id(lowerCamel+"s").Index().Op("*").Add(helper.UseEntity(camel)),
		).Params(jen.Index().Id("int64")),
		jen.Id("BatchInsertWithFunc").Params(
			jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
			jen.Id(lowerCamel+"s").Index().Op("*").Add(helper.UseEntity(camel)),
			jen.List(jen.Id("pid"), jen.Id("sid")).Id("int64"), helper.GenDeclAnonymousFunc(),
		).Params(jen.Index().Id("int64")),
		jen.Line(),
		helper.InferColumn(
			jen.Id("DeleteByID").Params(
				jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id("id").Id("int64"),
			).Params(jen.Id("bool")), "id", columns,
		),
		helper.InferColumn(
			jen.Id("DeleteByIDs").Params(
				jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id("ids").Op("...").Id("int64"),
			).Params(jen.Id("bool")), "id", columns,
		),
		helper.InferColumn(
			jen.Id("BatchDeleteByID").Params(
				jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id("ids").Index().Id("int64"),
			).Params(jen.Id("bool")), "id", columns,
		),
		jen.Line(),
		helper.InferColumn(
			jen.Id("UpdateByID").Params(
				jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
			).Params(jen.Id("bool")), "id", columns,
		),
		helper.InferColumn(
			jen.Id("UpdateNonNilByID").Params(
				jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
			).Params(jen.Id("bool")), "id", columns,
		),
		helper.InferColumn(
			jen.Id("UpdateWithFuncByID").Params(
				jen.Id("tx").Op("*").Add(helper.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
				helper.GenDeclAnonymousFunc(),
			).Params(jen.Id("bool")), "id", columns,
		),
		helper.InferColumn(
			jen.Id("BatchUpdateWithFuncByID").Params(
				jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
				jen.Id(lowerCamel+"s").Index().Op("*").Add(helper.UseEntity(camel)), helper.GenDeclAnonymousFunc(),
			).Params(jen.Id("bool")), "id", columns,
		),
	)
}

func (rec *GenBaseRepo) genStructAutoGen() jen.Code {
	return jen.Line().Comment("autoGen 该结构体自动生成, 请勿修改").Line().Type().Id("autoGen").Struct(jen.Id("ctx").Op("*").Add(helper.UseGin("Context")))
}

func (rec *GenBaseRepo) genFuncGetDbCtx() jen.Code {
	return jen.Line().Comment("getDbCtx 获取 DB 的初始上下文").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("getDbCtx").
		Params().Params(helper.UseContext("Context")).
		Block(
			jen.Return().Add(helper.UseContext("WithValue")).Call(
				jen.Add(helper.UseContext("Background")).Call(),
				jen.Add(helper.UseConstant("TraceIdKey")),
				jen.Id("ag").Dot("ctx").Dot("GetString").Call(helper.UseConstant("TraceIdKey")),
			),
		)
}

func (rec *GenBaseRepo) genFuncMapperAll(table string, columns []helper.Column) jen.Code {
	var codes = make([]jen.Code, len(columns))
	for i, column := range columns {
		codes[i] = helper.RenderAndField("r", column.ColumnName)
	}

	camel := strcase.ToCamel(table)

	return jen.Line().Comment("mapperAll 映射实体的所有字体").Line().Func().Id("mapperAll").Params().
		Params(
			jen.Op("*").Add(helper.UseEntity(camel)),
			jen.Index().Id("any"),
		).
		Block(
			jen.Var().Id("r").Op("=").Op("&").Add(helper.UseEntity(camel)).Values(),
			jen.Var().Id("cs").Op("=").Index().Id("any").Values(
				codes...,
			),
			jen.Return().List(
				jen.Id("r"),
				jen.Id("cs"),
			),
		)
}

func (rec *GenBaseRepo) genFuncMapperNumeric() jen.Code {
	return jen.Line().Comment("mapperNumeric 映射数值型").Line().Func().Id("mapperNumeric").Types(
		jen.Id("T").Union(
			jen.Int(), jen.Int64(),
		),
	).Params().Params(
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

func (rec *GenBaseRepo) genFuncCalcInsertField(table string, columns []helper.Column) jen.Code {
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
				helper.RenderStarField(lowerCamel, columnName),
			),
		)
		codes = append(codes, code)
	}

	codes = append(
		codes, jen.Return().List(
			jen.Add(helper.UseStrings("Join")).Call(
				jen.Id("fields"),
				jen.Lit(", "),
			),
			jen.Add(helper.UseStrings("Join")).Call(
				jen.Id("places"),
				jen.Lit(", "),
			),
			jen.Id("values"),
		),
	)

	return jen.Line().Comment("calcInsertField 计算待插入的字段").Line().Func().Id("calcInsertField").Params(
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
		helper.GenDeclAnonymousFunc(),
	).Params(
		jen.Id("string"),
		jen.Id("string"),
		jen.Index().Id("any"),
	).Block(codes...)
}

func (rec *GenBaseRepo) genFuncCalcUpdateField(table string, columns []helper.Column) jen.Code {
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
				jen.Lit(columnName+" = ?"),
			),
			jen.Id("values").Op("=").Id("append").Call(
				jen.Id("values"),
				helper.RenderStarField(lowerCamel, columnName),
			),
		)
		codes = append(codes, code)
	}

	codes = append(
		codes, jen.Return().List(
			jen.Add(helper.UseStrings("Join")).Call(
				jen.Id("fields"),
				jen.Lit(", "),
			),
			jen.Id("values"),
		),
	)

	return jen.Line().Comment("calcUpdateField 计算待更新的字段").Line().Func().Id("calcUpdateField").Params(
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
		helper.GenDeclAnonymousFunc(),
	).Params(
		jen.Id("string"),
		jen.Index().Id("any"),
	).Block(codes...)
}

func (rec *GenBaseRepo) genFuncInternalSelectByIDs(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	fields := helper.AllFields(columns)

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s ", fields, table, helper.P_KEY)

	return jen.Line().Comment("internalSelectByIDs 根据 ID 列表插入节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalSelectByIDs").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id("db").Op("*").Add(helper.UseSql("DB")),
		jen.Id("ids").Index().Id("int64"),
	).Params(jen.Index().Op("*").Add(helper.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID 列表: %+v 的数据"),
			jen.Id("ids"),
		),
		jen.Var().Id("sqlBuilder").Add(helper.UseStrings("Builder")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql)),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
		).Else().Block(
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("IN (")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Add(helper.UseUtil("GenPlaceholder")).Call(jen.Id("ids"))),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(")")),
		),
		helper.AddNsField(columns, "sqlBuilder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(helper.UD_COND_KEY+";")),
		jen.Id("errorHandler").Op(":=").Add(helper.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Var().Id("stmt").Op("*").Add(helper.UseSql("Stmt")),
		jen.Var().Id("err").Id("error"),
		jen.If(jen.Id("tx").Op("!=").Id("nil")).Block(
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Defer().Add(helper.UseUtil("DeferClose")).Call(
				jen.Id("stmt"),
				jen.Id("errorHandler"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
		).Else().Block(
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op("=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Defer().Add(helper.UseUtil("DeferClose")).Call(
				jen.Id("stmt"),
				jen.Id("errorHandler"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
		),
		jen.Id("bindValues").Op(":=").Add(helper.UseUtil("ToAnyItems")).Call(jen.Id("ids")),
		helper.AddNsValueWithName(columns, "bindValues"),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("bindValues").Op("..."),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Add(helper.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Add(helper.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenBaseRepo) genFuncInternalDirectInsert(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("internalDirectInsert 直接插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectInsert").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
		helper.GenDeclAnonymousFunc(),
	).Params(jen.Id("int64")).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("插入数据: %+v"),
			jen.Op("*").Id(lowerCamel),
		),
		jen.List(
			jen.Id("fields"),
			jen.Id("places"),
			jen.Id("values"),
		).Op(":=").Id("calcInsertField").Call(
			jen.Id(lowerCamel),
			jen.Id("fn"),
		),
		jen.Var().Id("sqlBuilder").Add(helper.UseStrings("Builder")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("INSERT INTO "+table+"(")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(") VALUES (")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("places")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(");")),
		jen.Id("errorHandler").Op(":=").Add(helper.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Add(helper.UseUtil("DeferClose")).Call(
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

func (rec *GenBaseRepo) genFuncInternalDirectDelete(table string, columns []helper.Column) jen.Code {

	var sql0 string
	if helper.HasColumn(helper.D_KEY, columns) {
		if helper.HasColumn(helper.NS_KEY, columns) {
			sql0 = fmt.Sprintf(
				"UPDATE %s SET %s = 1 WHERE %s = ? AND %s AND %s;", table, helper.D_KEY, helper.P_KEY,
				helper.NS_COND_KEY, helper.UD_COND_KEY,
			)
		} else {
			sql0 = fmt.Sprintf(
				"UPDATE %s SET %s = 1 WHERE %s = ? AND %s;", table, helper.D_KEY, helper.P_KEY,
				helper.UD_COND_KEY,
			)
		}
	} else {
		if helper.HasColumn(helper.NS_KEY, columns) {
			sql0 = fmt.Sprintf("DELETE FROM %s WHERE %s = ? AND %s;", table, helper.P_KEY, helper.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("DELETE FROM %s WHERE %s = ?;", table, helper.P_KEY)
		}
	}

	return jen.Line().Comment("internalDirectDelete 直接删除(逻辑 or 物理)").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectDelete").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id("id").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),

		jen.Id("errorHandler").Op(":=").Add(helper.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(helper.UseUtil("DeferClose")).Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("result"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("ExecContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			helper.AddNsSingleValue(columns),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("af"),
			jen.Id("err"),
		).Op(":=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.If(jen.Id("af").Op("==").Lit(1)).Block(jen.Return().Id("true")),
		jen.Id("panic").Call(jen.Lit("删除错误")),
	)
}

func (rec *GenBaseRepo) genFuncSelectByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("SelectByID 根据 ID 查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByID").Params(jen.Id("id").Id("int64")).
		Params(jen.Op("*").Add(helper.UseEntity(camel))).Block(
		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Index().Id("int64").Values(jen.Id("id"))),
		jen.If(jen.Id("len").Call(jen.Id("ds")).Op("==").Lit(1)).Block(jen.Return().Id("ds").Index(jen.Lit(0))),
		jen.Return().Id("nil"),
	)
}

func (rec *GenBaseRepo) genFuncSelectByIDs(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("SelectByIDs 根据 ID 列表查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).
		Params(jen.Index().Op("*").Add(helper.UseEntity(camel))).Block(
		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Id("ids")),
		jen.Return().Id("ds"),
	)
}

func (rec *GenBaseRepo) genFuncBatchSelectByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("BatchSelectByID 根据 ID 批量查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).
		Params(jen.Index().Op("*").Add(helper.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID 列表: %+v 的数据"),
			jen.Id("ids"),
		),
		jen.Id("db").Op(":=").Add(helper.UseDatabase("FetchDB")).Call(),
		jen.Return().Id("ag").Dot("internalSelectByIDs").Call(
			jen.Id("nil"),
			jen.Id("db"),
			jen.Id("ids"),
		),
	)
}

func (rec *GenBaseRepo) genFuncSelectByName(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if helper.HasColumn(helper.D_KEY, columns) {
		if helper.HasColumn(helper.NS_KEY, columns) {
			sql0 = fmt.Sprintf(
				"SELECT %s FROM %s WHERE %s like ? AND %s AND %s;", helper.AllFields(columns), table,
				helper.N_KEY, helper.NS_COND_KEY, helper.UD_COND_KEY,
			)
		} else {
			sql0 = fmt.Sprintf(
				"SELECT %s FROM %s WHERE %s like ? AND %s;", helper.AllFields(columns), table, helper.N_KEY,
				helper.UD_COND_KEY,
			)
		}
	} else {
		if helper.HasColumn(helper.NS_KEY, columns) {
			sql0 = fmt.Sprintf(
				"SELECT %s FROM %s WHERE %s like ? AND %s;", helper.AllFields(columns), table, helper.N_KEY,
				helper.NS_COND_KEY,
			)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s like ?;", helper.AllFields(columns), table, helper.N_KEY)
		}
	}
	return jen.Line().Comment("SelectByName 根据名称查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(helper.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 NAME: %+v 的数据"),
			jen.Id("name"),
		),
		jen.Id("db").Op(":=").Add(helper.UseDatabase("FetchDB")).Call(),
		jen.Id("errorHandler").Op(":=").Add(helper.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(helper.UseUtil("DeferClose")).Call(
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
			helper.AddNsSingleValue(columns),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Add(helper.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Add(helper.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenBaseRepo) genFuncInsert(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("Insert").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(helper.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenBaseRepo) genFuncBatchInsert(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("BatchInsert 批量插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsert").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(helper.UseEntity(camel)),
	).Params(jen.Index().Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Id(lowerCamel+"s"),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Id("len").Call(jen.Id(lowerCamel+"s"))).Block(jen.Return().Id("ids")),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回记录数等于插入记录数时成功")),
	)
}

func (rec *GenBaseRepo) genFuncInsertNonNil(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertNonNil 插入非空字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNil").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(helper.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenBaseRepo) genFuncInsertWithFunc(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertWithFunc 根据函数插入字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFunc").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
		helper.GenDeclAnonymousFunc(),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(helper.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("fn"),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenBaseRepo) genFuncBatchInsertWithFunc(table string, columns []helper.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("BatchInsertWithFunc 根据函数批量插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertWithFunc").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(helper.UseEntity(camel)),
		helper.GenDeclAnonymousFunc(),
	).Params(jen.Index().Id("int64")).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Info").Call(
			jen.Lit("批量插入数据"),
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
			jen.Id("ids").Index(jen.Id("i")).Op("=").Id("ag").Dot("internalDirectInsert").Call(
				jen.Id("tx"),
				jen.Id(lowerCamel),
				jen.Id("fn"),
			),
		),
		jen.Return().Id("ids"),
	)
}

func (rec *GenBaseRepo) genFuncDeleteByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	return jen.Line().Comment("DeleteByID 根据 ID 删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByID").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id("id").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
			jen.Id("tx"),
			jen.Index().Id("int64").Values(jen.Id("id")),
		),
	)
}

func (rec *GenBaseRepo) genFuncDeleteByIDs(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	return jen.Line().Comment("DeleteByIDs 根据 ID 列表删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByIDs").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id("ids").Op("...").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
			jen.Id("tx"),
			jen.Id("ids"),
		),
	)
}

func (rec *GenBaseRepo) genFuncBatchDeleteByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	return jen.Line().Comment("BatchDeleteByID 根据 ID 批量删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchDeleteByID").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id("ids").Index().Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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

func (rec *GenBaseRepo) genFuncUpdateByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("UpdateByID 根据 ID 批量更新").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateByID").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(helper.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
	)
}

func (rec *GenBaseRepo) genFuncUpdateNonNilByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("UpdateNonNilByID 根据 ID 更新非空字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateNonNilByID").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(helper.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
	)
}

func (rec *GenBaseRepo) genFuncUpdateWithFuncByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("UpdateWithFuncByID 根据 ID 更新满足函数的字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateWithFuncByID").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(helper.UseEntity(camel)),
		helper.GenDeclAnonymousFunc(),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(helper.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("fn"),
		),
	)
}

func (rec *GenBaseRepo) genFuncBatchUpdateWithFuncByID(table string, columns []helper.Column) jen.Code {
	if !helper.HasColumn(helper.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	var sql0, sql1 string
	sql0 = fmt.Sprintf("UPDATE %s SET ", table)
	if helper.HasColumn(helper.D_KEY, columns) {
		if helper.HasColumn(helper.NS_KEY, columns) {
			sql1 = fmt.Sprintf(" WHERE %s = ? AND %s AND %s;", helper.P_KEY, helper.NS_COND_KEY, helper.UD_COND_KEY)
		} else {
			sql1 = fmt.Sprintf(" WHERE %s = ? AND %s;", helper.P_KEY, helper.UD_COND_KEY)
		}
	} else {
		if helper.HasColumn(helper.NS_KEY, columns) {
			sql1 = fmt.Sprintf(" WHERE %s = ? AND %s;", helper.P_KEY, helper.NS_COND_KEY)
		} else {
			sql1 = fmt.Sprintf(" WHERE %s = ?;", helper.P_KEY)
		}
	}

	return jen.Line().Comment("BatchUpdateWithFuncByID 根据 ID 批量更新满足函数的字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchUpdateWithFuncByID").Params(
		jen.Id("tx").Op("*").Add(helper.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(helper.UseEntity(camel)),
		helper.GenDeclAnonymousFunc(),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Add(helper.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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
			jen.Var().Id("sqlBuilder").Add(helper.UseStrings("Builder")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql0)),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql1)),
			jen.Id("values").Op("=").Id("append").Call(
				jen.Id("values"),
				jen.Id("id"),
				helper.AddNsSingleValue(columns),
			),
			jen.Id("errorHandler").Op(":=").Add(helper.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
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

func (rec *GenBaseRepo) GenFile(table string, columns []helper.Column) *jen.File {
	file := jen.NewFile(table)

	file.HeaderComment("Code generated by tabuyos. DO NOT EDIT.")
	file.PackageComment("Package repository")
	file.PackageComment("@author tabuyos")
	file.PackageComment("@since " + time.Now().Format("2006/01/02"))
	file.PackageComment("@description " + table)

	file.Add(rec.genInterfaceAutoGen(table, columns))
	file.Add(rec.genStructAutoGen())
	file.Add(rec.genFuncGetDbCtx())

	file.Add(rec.genFuncMapperAll(table, columns))
	file.Add(rec.genFuncMapperNumeric())

	file.Add(rec.genFuncCalcInsertField(table, columns))
	file.Add(rec.genFuncCalcUpdateField(table, columns))

	file.Add(rec.genFuncInternalSelectByIDs(table, columns))
	file.Add(rec.genFuncInternalDirectInsert(table, columns))
	file.Add(rec.genFuncInternalDirectDelete(table, columns))

	file.Add(rec.genFuncSelectByID(table, columns))
	file.Add(rec.genFuncSelectByIDs(table, columns))
	file.Add(rec.genFuncBatchSelectByID(table, columns))

	file.Add(rec.genFuncSelectByName(table, columns))
	file.Add(genFuncSelectWithPage(table, columns))

	file.Add(rec.genFuncInsert(table, columns))
	file.Add(rec.genFuncInsertNonNil(table, columns))
	file.Add(rec.genFuncInsertWithFunc(table, columns))
	file.Add(rec.genFuncBatchInsert(table, columns))
	file.Add(rec.genFuncBatchInsertWithFunc(table, columns))

	file.Add(rec.genFuncDeleteByID(table, columns))
	file.Add(rec.genFuncDeleteByIDs(table, columns))
	file.Add(rec.genFuncBatchDeleteByID(table, columns))

	file.Add(rec.genFuncUpdateByID(table, columns))
	file.Add(rec.genFuncUpdateNonNilByID(table, columns))
	file.Add(rec.genFuncUpdateWithFuncByID(table, columns))
	file.Add(rec.genFuncBatchUpdateWithFuncByID(table, columns))
	return file
}
