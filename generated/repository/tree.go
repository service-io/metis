// Package repository
// @author tabuyos
// @since 2023/8/1
// @description repository
package repository

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"metis/generated"
)

type GenTreeRepo struct {
}

func (rec *GenTreeRepo) genInterfaceAutoGen(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Comment("iAutoGen 该接口自动生成, 请勿修改").Line().Type().Id("iAutoGen").Interface(
		generated.InferColumn(jen.Id("SelectByID").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(generated.UseEntity(camel))), "id", columns),
		generated.InferColumn(jen.Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))), "id", columns),
		generated.InferColumn(jen.Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))), "id", columns),
		jen.Line(),
		generated.InferColumn(jen.Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))), "name", columns),
		jen.Line(),
		jen.Id("SelectMaxLevel").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")),
		jen.Id("SelectMaxRight").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")),
		jen.Id("SelectMaxLeft").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")),
		jen.Id("SelectMaxTreeNo").Params().Params(jen.Id("int")),
		jen.Id("SelectAllPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectDirectPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectBrother").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectBrotherAndSelf").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectAncestorChain").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectAncestor").Params(jen.Id("id").Id("int64"), jen.Id("level").Id("int")).Params(jen.Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectParent").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectByTreeNoAndLevel").Params(jen.List(jen.Id("treeNo"), jen.Id("level")).Id("int")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectByLevel").Params(jen.Id("level").Id("int")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectRoot").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectLeafOfNodeWithPage").Params(jen.Id("id").Id("int64"), jen.List(jen.Id("page"), jen.Id("size")).Id("uint")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel)), jen.Id("int64")),
		jen.Id("SelectAllLeafOfNode").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Id("SelectAllRoot").Params().Params(jen.Index().Op("*").Add(generated.UseEntity(camel))),
		jen.Line(),
		jen.Id("Insert").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel))).Params(jen.Id("int64")),
		jen.Id("InsertUnderNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), jen.Id("pid").Id("int64")).Params(jen.Id("int64")),
		jen.Id("InsertBetweenNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), jen.List(jen.Id("pid"), jen.Id("sid")).Id("int64")).Params(jen.Id("int64")),
		jen.Id("BatchInsert").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel))).Params(jen.Index().Id("int64")),
		jen.Id("BatchInsertUnderNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)), jen.Id("pid").Id("int64")).Params(jen.Index().Id("int64")),
		jen.Id("BatchInsertBetweenNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)), jen.List(jen.Id("pid"), jen.Id("sid")).Id("int64")).Params(jen.Index().Id("int64")),
		jen.Id("InsertNonNil").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel))).Params(jen.Id("int64")),
		jen.Id("InsertNonNilUnderNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), jen.Id("pid").Id("int64")).Params(jen.Id("int64")),
		jen.Id("InsertNonNilBetweenNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), jen.List(jen.Id("pid"), jen.Id("sid")).Id("int64")).Params(jen.Id("int64")),
		jen.Id("InsertWithFunc").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), generated.GenDeclAnonymousFunc()).Params(jen.Id("int64")),
		jen.Id("InsertWithFuncUnderNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), jen.Id("pid").Id("int64"), generated.GenDeclAnonymousFunc()).Params(jen.Id("int64")),
		jen.Id("InsertWithFuncBetweenNode").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), jen.List(jen.Id("pid"), jen.Id("sid")).Id("int64"), generated.GenDeclAnonymousFunc()).Params(jen.Id("int64")),
		jen.Id("BatchInsertWithFunc").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)), jen.List(jen.Id("pid"), jen.Id("sid")).Id("int64"), generated.GenDeclAnonymousFunc()).Params(jen.Index().Id("int64")),
		jen.Line(),
		generated.InferColumn(jen.Id("DeleteByID").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id("id").Id("int64")).Params(jen.Id("bool")), "id", columns),
		generated.InferColumn(jen.Id("DeleteByIDs").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id("ids").Op("...").Id("int64")).Params(jen.Id("bool")), "id", columns),
		generated.InferColumn(jen.Id("BatchDeleteByID").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id("ids").Index().Id("int64")).Params(jen.Id("bool")), "id", columns),
		jen.Line(),
		generated.InferColumn(jen.Id("UpdateByID").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel))).Params(jen.Id("bool")), "id", columns),
		generated.InferColumn(jen.Id("UpdateNonNilByID").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel))).Params(jen.Id("bool")), "id", columns),
		generated.InferColumn(jen.Id("UpdateWithFuncByID").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)), generated.GenDeclAnonymousFunc()).Params(jen.Id("bool")), "id", columns),
		generated.InferColumn(jen.Id("BatchUpdateWithFuncByID").Params(jen.Id("tx").Op("*").Add(generated.UseSql("Tx")), jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)), generated.GenDeclAnonymousFunc()).Params(jen.Id("bool")), "id", columns),
	)
}

func (rec *GenTreeRepo) genStructAutoGen() jen.Code {
	return jen.Line().Comment("autoGen 该结构体自动生成, 请勿修改").Line().Type().Id("autoGen").Struct(jen.Id("ctx").Op("*").Add(generated.UseGin("Context")))
}

func (rec *GenTreeRepo) genFuncGetDbCtx() jen.Code {
	return jen.Line().Comment("getDbCtx 获取 DB 的初始上下文").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("getDbCtx").
		Params().Params(generated.UseContext("Context")).
		Block(
			jen.Return().Add(generated.UseContext("WithValue")).Call(
				jen.Add(generated.UseContext("Background")).Call(),
				jen.Add(generated.UseConstant("TraceIdKey")),
				jen.Id("ag").Dot("ctx").Dot("GetString").Call(generated.UseConstant("TraceIdKey")),
			),
		)
}

func (rec *GenTreeRepo) genFuncMapperAll(table string, columns []generated.Column) jen.Code {
	var codes = make([]jen.Code, len(columns))
	for i, column := range columns {
		codes[i] = generated.RenderAndField("r", column.ColumnName)
	}

	camel := strcase.ToCamel(table)

	return jen.Line().Comment("mapperAll 映射实体的所有字体").Line().Func().Id("mapperAll").Params().
		Params(
			jen.Op("*").Add(generated.UseEntity(camel)),
			jen.Index().Id("any"),
		).
		Block(
			jen.Var().Id("r").Op("=").Op("&").Add(generated.UseEntity(camel)).Values(),
			jen.Var().Id("cs").Op("=").Index().Id("any").Values(
				codes...,
			),
			jen.Return().List(
				jen.Id("r"),
				jen.Id("cs"),
			),
		)
}

func (rec *GenTreeRepo) genFuncMapperNumeric() jen.Code {
	return jen.Line().Comment("mapperNumeric 映射数值型").Line().Func().Id("mapperNumeric").Types(jen.Id("T").Union(jen.Int(), jen.Int64())).Params().Params(
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

func (rec *GenTreeRepo) genFuncTreeInfoSelectSql(table string, columns []generated.Column) jen.Code {
	var sql string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql = fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ? AND %s AND %s;", generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, table, generated.P_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql = fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ? AND %s;", generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, table, generated.P_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql = fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ? AND %s;", generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, table, generated.P_KEY, generated.NS_COND_KEY)
		} else {
			sql = fmt.Sprintf("SELECT %s, %s, %s, %s FROM %s WHERE %s = ?;", generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, table, generated.P_KEY)
		}
	}
	return jen.Line().Comment("treeInfoSelectSql 获取树型表的基础信息").Line().Func().Id("treeInfoSelectSql").Params().Params(jen.Id("string")).Block(jen.Return().Lit(sql))
}

func (rec *GenTreeRepo) genFuncCalcInsertField(table string, columns []generated.Column) jen.Code {
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
				generated.RenderStarField(lowerCamel, columnName),
			),
		)
		codes = append(codes, code)
	}

	codes = append(
		codes, jen.Return().List(
			jen.Add(generated.UseStrings("Join")).Call(
				jen.Id("fields"),
				jen.Lit(", "),
			),
			jen.Add(generated.UseStrings("Join")).Call(
				jen.Id("places"),
				jen.Lit(", "),
			),
			jen.Id("values"),
		),
	)

	return jen.Line().Comment("calcInsertField 计算待插入的字段").Line().Func().Id("calcInsertField").Params(
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		generated.GenDeclAnonymousFunc(),
	).Params(
		jen.Id("string"),
		jen.Id("string"),
		jen.Index().Id("any"),
	).Block(codes...)
}

func (rec *GenTreeRepo) genFuncCalcUpdateField(table string, columns []generated.Column) jen.Code {
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
				generated.RenderStarField(lowerCamel, columnName),
			),
		)
		codes = append(codes, code)
	}

	codes = append(
		codes, jen.Return().List(
			jen.Add(generated.UseStrings("Join")).Call(
				jen.Id("fields"),
				jen.Lit(", "),
			),
			jen.Id("values"),
		),
	)

	return jen.Line().Comment("calcUpdateField 计算待更新的字段").Line().Func().Id("calcUpdateField").Params(
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		generated.GenDeclAnonymousFunc(),
	).Params(
		jen.Id("string"),
		jen.Index().Id("any"),
	).Block(codes...)
}

func (rec *GenTreeRepo) genFuncInternalSelectNodeByIDs(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	fields := generated.AllFields(columns)

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s ", fields, table, generated.P_KEY)

	return jen.Line().Comment("internalSelectNodeByIDs 根据 ID 列表插入节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalSelectNodeByIDs").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id("db").Op("*").Add(generated.UseSql("DB")),
		jen.Id("ids").Index().Id("int64"),
	).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID 列表: %+v 的数据"),
			jen.Id("ids"),
		),
		jen.Var().Id("sqlBuilder").Add(generated.UseStrings("Builder")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql)),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("= ?")),
		).Else().Block(
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("IN (")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Add(generated.UseUtil("GenPlaceholder")).Call(jen.Id("ids"))),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(")")),
		),
		generated.AddNsField(columns, "sqlBuilder"),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(generated.UD_COND_KEY+";")),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Var().Id("stmt").Op("*").Add(generated.UseSql("Stmt")),
		jen.Var().Id("err").Id("error"),
		jen.If(jen.Id("tx").Op("!=").Id("nil")).Block(
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op("=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
				jen.Id("stmt"),
				jen.Id("errorHandler"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
		).Else().Block(
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op("=").Id("db").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
			jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
				jen.Id("stmt"),
				jen.Id("errorHandler"),
			),
			jen.Id("errorHandler").Call(jen.Id("err")),
		),
		jen.Id("bindValues").Op(":=").Add(generated.UseUtil("ToAnyItems")).Call(jen.Id("ids")),
		generated.AddNsValueWithName(columns, "bindValues"),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("bindValues").Op("..."),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncInternalDirectInsert(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("internalDirectInsert 直接插入树节点, 需要提前计算好树相关信息").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectInsert").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Id("int64")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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
		jen.Var().Id("sqlBuilder").Add(generated.UseStrings("Builder")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit("INSERT INTO "+table+"(")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(") VALUES (")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("places")),
		jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(");")),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("sqlBuilder").Dot("String").Call()),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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

func (rec *GenTreeRepo) genFuncInternalUpdateNodeInBothWhenInsert(table string, columns []generated.Column) jen.Code {
	// camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0, sql1, sql2 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.NS_KEY, generated.UD_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s = ? AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_KEY, generated.UD_COND_KEY)
			sql2 = fmt.Sprintf("UPDATE %s SET %s = %s + 1, %s = %s + 1, %s = %s + 1 WHERE %s >= ? AND %s <= ? AND %s = ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.R_KEY, generated.R_KEY, generated.LL_KEY, generated.LL_KEY, generated.LL_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.UD_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
			sql2 = fmt.Sprintf("UPDATE %s SET %s = %s + 1, %s = %s + 1, %s = %s + 1 WHERE %s >= ? AND %s <= ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.R_KEY, generated.R_KEY, generated.LL_KEY, generated.LL_KEY, generated.LL_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}

	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s = ?;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.NS_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s = ?;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_KEY)
			sql2 = fmt.Sprintf("UPDATE %s SET %s = %s + 1, %s = %s + 1, %s = %s + 1 WHERE %s >= ? AND %s <= ? AND %s = ? AND %s = ?;", table, generated.L_KEY, generated.L_KEY, generated.R_KEY, generated.R_KEY, generated.LL_KEY, generated.LL_KEY, generated.LL_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_KEY)
		} else {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ?;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ?;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY)
			sql2 = fmt.Sprintf("UPDATE %s SET %s = %s + 1, %s = %s + 1, %s = %s + 1 WHERE %s >= ? AND %s <= ? AND %s = ?;", table, generated.L_KEY, generated.L_KEY, generated.R_KEY, generated.R_KEY, generated.LL_KEY, generated.LL_KEY, generated.LL_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("internalUpdateNodeInBothWhenInsert 在两个节点间插入时更新").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalUpdateNodeInBothWhenInsert").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.List(
			jen.Id("left"),
			jen.Id("right"),
			jen.Id("treeNo"),
		).Id("int"),
	).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
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
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Lit(sql1)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
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
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Lit(sql2)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
	)
}

func (rec *GenTreeRepo) genFuncInternalUpdateNodeInOnlyPrecursorWhenInsert(table string, columns []generated.Column) jen.Code {
	var sql0, sql1 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s >= ? AND %s = ? AND %s AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.UD_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s >= ? AND %s = ? AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.NS_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s >= ? AND %s = ? AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s > ? AND %s = ?;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s + 2 WHERE %s >= ? AND %s = ?;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("internalUpdateNodeInOnlyPrecursorWhenInsert 插入至前驱节点时更新").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalUpdateNodeInOnlyPrecursorWhenInsert").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.List(
			jen.Id("right"),
			jen.Id("treeNo"),
		).Id("int"),
	).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
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
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Lit(sql1)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
	)
}

func (rec *GenTreeRepo) genFuncInternalInsertWithFunc(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("internalInsertWithFunc 根据函数进行插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalInsertWithFunc").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Id("int64")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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

func (rec *GenTreeRepo) genFuncInternalDirectDelete(table string, columns []generated.Column) jen.Code {

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = 1 WHERE %s >= ? AND %s <= ? AND %s = ? AND %s AND %s;", table, generated.D_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = 1 WHERE %s >= ? AND %s <= ? AND %s = ? AND %s;", table, generated.D_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("DELETE FROM %s WHERE %s >= ? AND %s <= ? AND %s = ? AND %s;", table, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("DELETE FROM %s WHERE %s >= ? AND %s <= ? AND %s = ?;", table, generated.L_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("internalDirectDelete 直接删除(逻辑 or 物理)").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalDirectDelete").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id("id").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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
			jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
			jen.List(
				jen.Id("stmt"),
				jen.Id("err"),
			).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
			jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
				generated.AddNsSingleValue(columns),
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

func (rec *GenTreeRepo) genFuncInternalUpdateNodeWhenDelete(table string, columns []generated.Column) jen.Code {
	var sql0, sql1 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ? AND %s AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ? AND %s AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.UD_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ? AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ? AND %s;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY, generated.NS_COND_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ? AND %s;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ?;", table, generated.L_KEY, generated.L_KEY, generated.L_KEY, generated.TN_KEY)
			sql1 = fmt.Sprintf("UPDATE %s SET %s = %s - ? WHERE %s > ? AND %s = ?;", table, generated.R_KEY, generated.R_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("internalUpdateNodeWhenDelete 删除时更新节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("internalUpdateNodeWhenDelete").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.List(
			jen.Id("delta"),
			jen.Id("right"),
			jen.Id("treeNo"),
		).Id("int"),
	).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
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
		).Op("=").Id("tx").Dot("Prepare").Call(jen.Lit(sql1)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("_"),
			jen.Id("err"),
		).Op("=").Id("result").Dot("RowsAffected").Call(),
		jen.Id("errorHandler").Call(jen.Id("err")),
	)
}

func (rec *GenTreeRepo) genFuncSelectByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("SelectByID 根据 ID 查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByID").Params(jen.Id("id").Id("int64")).
		Params(jen.Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Index().Id("int64").Values(jen.Id("id"))),
		jen.If(jen.Id("len").Call(jen.Id("ds")).Op("==").Lit(1)).Block(jen.Return().Id("ds").Index(jen.Lit(0))),
		jen.Return().Id("nil"),
	)
}

func (rec *GenTreeRepo) genFuncSelectByIDs(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("SelectByIDs 根据 ID 列表查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByIDs").Params(jen.Id("ids").Op("...").Id("int64")).
		Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("ds").Op(":=").Id("ag").Dot("BatchSelectByID").Call(jen.Id("ids")),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncBatchSelectByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("BatchSelectByID 根据 ID 批量查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchSelectByID").Params(jen.Id("ids").Index().Id("int64")).
		Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID 列表: %+v 的数据"),
			jen.Id("ids"),
		),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.Return().Id("ag").Dot("internalSelectNodeByIDs").Call(
			jen.Id("nil"),
			jen.Id("db"),
			jen.Id("ids"),
		),
	)
}

func (rec *GenTreeRepo) genFuncSelectByName(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s like ? AND %s AND %s;", generated.AllFields(columns), table, generated.N_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s like ? AND %s;", generated.AllFields(columns), table, generated.N_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s like ? AND %s;", generated.AllFields(columns), table, generated.N_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s like ?;", generated.AllFields(columns), table, generated.N_KEY)
		}
	}
	return jen.Line().Comment("SelectByName 根据名称查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByName").Params(jen.Id("name").Id("string")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 NAME: %+v 的数据"),
			jen.Id("name"),
		),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectMaxLevel(table string, columns []generated.Column) jen.Code {
	// camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s AND %s;", generated.LL_KEY, table, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s;", generated.LL_KEY, table, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s;", generated.LL_KEY, table, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ?;", generated.LL_KEY, table, generated.TN_KEY)
		}
	}
	return jen.Line().Comment("SelectMaxLevel 查询最大层级").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxLevel").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v 的最大层级"),
			jen.Id("treeNo"),
		),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("treeNo"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectMaxRight(table string, columns []generated.Column) jen.Code {
	// camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s AND %s;", generated.R_KEY, table, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s;", generated.R_KEY, table, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s;", generated.R_KEY, table, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ?;", generated.R_KEY, table, generated.TN_KEY)
		}
	}
	return jen.Line().Comment("SelectMaxRight 查询最大右值").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxRight").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v 的最大层级"),
			jen.Id("treeNo"),
		),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("treeNo"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectMaxLeft(table string, columns []generated.Column) jen.Code {
	// camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s AND %s;", generated.L_KEY, table, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s;", generated.L_KEY, table, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ? AND %s;", generated.L_KEY, table, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s = ?;", generated.L_KEY, table, generated.TN_KEY)
		}
	}
	return jen.Line().Comment("SelectMaxLeft 查询最大左值").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxLeft").Params(jen.Id("treeNo").Id("int")).Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v 的最大层级"),
			jen.Id("treeNo"),
		),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("stmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("treeNo"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectMaxTreeNo(table string, columns []generated.Column) jen.Code {
	// camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s AND %s;", generated.TN_KEY, table, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s;", generated.TN_KEY, table, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s WHERE %s;", generated.TN_KEY, table, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT MAX(%s) FROM %s;", generated.TN_KEY, table)
		}
	}

	return jen.Line().Comment("SelectMaxTreeNo 查询最大树号").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectMaxTreeNo").Params().Params(jen.Id("int")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Info").Call(jen.Lit("查询最大TN")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.Id("row").Op(":=").Id("db").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Lit(sql0),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int")),
		),
		jen.Return().Op("*").Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectAllPosterity(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s > ? AND %s < ? AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s > ? AND %s < ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s > ? AND %s < ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s > ? AND %s < ? AND %s = ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}
	return jen.Line().Comment("SelectAllPosterity 查询所有子代").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAllPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的所有子代(含自身)数据"),
			jen.Id("id"),
		),
		jen.Comment("recorder.Warn(\"不建议查询全部子代, 如果树比较大, 数据量将会非常大\")"),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectDirectPosterity(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s > ? AND %s < ? AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s > ? AND %s < ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s > ? AND %s < ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s > ? AND %s < ? AND %s = ?;", generated.AllFields(columns), table, generated.LL_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectDirectPosterity 查询直系子代").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectDirectPosterity").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的直系子代数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectBrother(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s != ? AND %s AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.P_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s != ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.P_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s != ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.P_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s != ?;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.P_KEY)
		}
	}

	return jen.Line().Comment("SelectBrother 查询兄弟(不含自身)").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectBrother").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的兄弟数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectBrotherAndSelf(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ?;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectBrotherAndSelf 查询兄弟和自身").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectBrotherAndSelf").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的兄弟以及自身数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectAncestorChain(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectAncestorChain 查询祖链").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAncestorChain").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的祖链数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectAncestor(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectAncestor 查询祖节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAncestor").Params(
		jen.Id("id").Id("int64"),
		jen.Id("level").Id("int"),
	).Params(jen.Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的祖代(%+v)数据"),
			jen.Id("id"),
			jen.Id("level"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectParent(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s < ? AND %s > ? AND %s = ? AND %s = ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.LL_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectParent 查询父节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectParent").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的父节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectByTreeNoAndLevel(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s = ?;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectByTreeNoAndLevel 根据树号和层级查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByTreeNoAndLevel").Params(
		jen.List(
			jen.Id("treeNo"),
			jen.Id("level"),
		).Id("int"),
	).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 TN: %+v LL: %+v 的同代数据"),
			jen.Id("treeNo"),
			jen.Id("level"),
		),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectByLevel(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?;", generated.AllFields(columns), table, generated.LL_KEY)
		}
	}

	return jen.Line().Comment("SelectByLevel 根据层级查询").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectByLevel").Params(jen.Id("level").Id("int")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 LL: %+v 的同代(跨树)数据"),
			jen.Id("level"),
		),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectRoot(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)
	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 AND %s = ? AND %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 AND %s = ?;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectRoot 查询根节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectRoot").Params(jen.Id("id").Id("int64")).Params(jen.Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的根节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("secondStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op("=").Id("secondStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectLeafOfNodeWithPage(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0, sql1 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s AND %s ORDER BY %s LIMIT ? OFFSET ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY, generated.L_KEY)
			sql1 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s ORDER BY %s LIMIT ? OFFSET ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY, generated.L_KEY)
			sql1 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s ORDER BY %s LIMIT ? OFFSET ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.L_KEY)
			sql1 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? ORDER BY %s LIMIT ? OFFSET ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.L_KEY)
			sql1 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ?;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectLeafOfNodeWithPage 查询对应节点叶子节点(分页)").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectLeafOfNodeWithPage").Params(
		jen.Id("id").Id("int64"),
		jen.List(
			jen.Id("page"),
			jen.Id("size"),
		).Id("uint"),
	).Params(
		jen.Index().Op("*").Add(generated.UseEntity(camel)),
		jen.Id("int64"),
	).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("分页查询 ID: %+v 的叶子节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("thirdStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql1)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("thirdStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op("=").Id("thirdStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Op("*").Id("currentNode").Dot("Left"),
			jen.Op("*").Id("currentNode").Dot("Right"),
			jen.Op("*").Id("currentNode").Dot("TreeNo"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("total").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperNumeric").Index(jen.Id("int64")),
		),
		jen.Return().List(
			jen.Id("ds"),
			jen.Op("*").Id("total"),
		),
	)
}

func (rec *GenTreeRepo) genFuncSelectAllLeafOfNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s AND %s ORDER BY %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY, generated.L_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s ORDER BY %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.UD_COND_KEY, generated.L_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? AND %s ORDER BY %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.NS_COND_KEY, generated.L_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s >= ? AND %s <= ? AND %s + 1 = %s AND %s = ? ORDER BY %s;", generated.AllFields(columns), table, generated.L_KEY, generated.R_KEY, generated.L_KEY, generated.R_KEY, generated.TN_KEY, generated.L_KEY)
		}
	}

	return jen.Line().Comment("SelectAllLeafOfNode 查询对应节点的所有叶子节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAllLeafOfNode").Params(jen.Id("id").Id("int64")).Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Sugar").Call().Dot("Infof").Call(
			jen.Lit("查询 ID: %+v 的所有叶子节点数据"),
			jen.Id("id"),
		),
		jen.Id("treeInfoSql").Op(":=").Id("treeInfoSelectSql").Call(),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("tx"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Begin").Call(),
		jen.Defer().Add(generated.UseUtil("HandleTx")).Call(
			jen.Id("tx"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("firstStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Id("treeInfoSql")),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("firstStmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("row").Op(":=").Id("firstStmt").Dot("QueryRowContext").Call(
			jen.Id("ag").Dot("getDbCtx").Call(),
			jen.Id("id"),
			generated.AddNsSingleValue(columns),
		),
		jen.Id("currentNode").Op(":=").Add(generated.UseUtil("Row")).Call(
			jen.Id("row"),
			jen.Id("mapperAll"),
		),
		jen.List(
			jen.Id("secondStmt"),
			jen.Id("err"),
		).Op(":=").Id("tx").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
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
			generated.AddNsSingleValue(columns),
		),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncSelectAllRoot(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	// lowerCamel := strcase.ToLowerCamel(table)

	var sql0 string
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 AND %s AND %s ORDER BY %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY, generated.TN_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 AND %s ORDER BY %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.UD_COND_KEY, generated.TN_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 AND %s ORDER BY %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.NS_COND_KEY, generated.TN_KEY)
		} else {
			sql0 = fmt.Sprintf("SELECT %s FROM %s WHERE %s = 1 ORDER BY %s;", generated.AllFields(columns), table, generated.LL_KEY, generated.TN_KEY)
		}
	}

	return jen.Line().Comment("SelectAllRoot 查询所有的根节点").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("SelectAllRoot").Params().Params(jen.Index().Op("*").Add(generated.UseEntity(camel))).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
		jen.Id("recorder").Dot("Info").Call(jen.Lit("查询的所有根节点数据")),
		jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
		jen.Id("db").Op(":=").Add(generated.UseDatabase("FetchDB")).Call(),
		jen.List(
			jen.Id("stmt"),
			jen.Id("err"),
		).Op(":=").Id("db").Dot("Prepare").Call(jen.Lit(sql0)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("stmt"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.List(
			jen.Id("rows"),
			jen.Id("err"),
		).Op(":=").Id("stmt").Dot("QueryContext").Call(jen.Id("ag").Dot("getDbCtx").Call(), generated.AddNsSingleValue(columns)),
		jen.Defer().Add(generated.UseUtil("DeferClose")).Call(
			jen.Id("rows"),
			jen.Id("errorHandler"),
		),
		jen.Id("errorHandler").Call(jen.Id("err")),
		jen.Id("ds").Op(":=").Add(generated.UseUtil("Rows")).Call(
			jen.Id("rows"),
			jen.Id("mapperAll"),
		),
		jen.Return().Id("ds"),
	)
}

func (rec *GenTreeRepo) genFuncInsert(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("Insert").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Lit(0),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncInsertUnderNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertUnderNode 插入至节点下方(做叶子节点)").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertUnderNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		jen.Id("pid").Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncInsertBetweenNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertBetweenNode 插入至两节点间").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertBetweenNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncBatchInsert(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("BatchInsert 批量插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsert").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)),
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

func (rec *GenTreeRepo) genFuncBatchInsertUnderNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("BatchInsertUnderNode 批量插入至节点下方").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertUnderNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)),
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

func (rec *GenTreeRepo) genFuncBatchInsertBetweenNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("BatchInsertBetweenNode 批量插入至两节点间(谨慎使用)").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertBetweenNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)),
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

func (rec *GenTreeRepo) genFuncInsertNonNil(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertNonNil 插入非空字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNil").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Lit(0),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncInsertNonNilUnderNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertNonNilUnderNode 插入非空字段并挂载到某节点下方").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNilUnderNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		jen.Id("pid").Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Lit(0),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncInsertNonNilBetweenNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertNonNilBetweenNode 插入非空字段并挂载到两节点之间").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertNonNilBetweenNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncInsertWithFunc(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertWithFunc 根据函数插入字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFunc").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Lit(0),
			jen.Lit(0),
			jen.Id("fn"),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncInsertWithFuncUnderNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertWithFuncUnderNode 根据函数插入并挂载到某节点下方").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFuncUnderNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		jen.Id("pid").Id("int64"),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Lit(0),
			jen.Id("fn"),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncInsertWithFuncBetweenNode(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("InsertWithFuncBetweenNode 根据函数插入并挂载到两节点之间").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("InsertWithFuncBetweenNode").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Id("int64")).Block(
		jen.Id("ids").Op(":=").Id("ag").Dot("BatchInsertWithFunc").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("pid"),
			jen.Id("sid"),
			jen.Id("fn"),
		),
		jen.If(jen.Id("len").Call(jen.Id("ids")).Op("==").Lit(1)).Block(jen.Return().Id("ids").Index(jen.Lit(0))),
		jen.Id("panic").Call(jen.Lit("插入失败, 仅返回一条记录时成功")),
	)
}

func (rec *GenTreeRepo) genFuncBatchInsertWithFunc(table string, columns []generated.Column) jen.Code {
	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)

	return jen.Line().Comment("BatchInsertWithFunc 根据函数批量插入").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchInsertWithFunc").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)),
		jen.List(
			jen.Id("pid"),
			jen.Id("sid"),
		).Id("int64"),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Index().Id("int64")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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

func (rec *GenTreeRepo) genFuncDeleteByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	return jen.Line().Comment("DeleteByID 根据 ID 删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByID").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id("id").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
			jen.Id("tx"),
			jen.Index().Id("int64").Values(jen.Id("id")),
		),
	)
}

func (rec *GenTreeRepo) genFuncDeleteByIDs(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	return jen.Line().Comment("DeleteByIDs 根据 ID 列表删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("DeleteByIDs").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id("ids").Op("...").Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchDeleteByID").Call(
			jen.Id("tx"),
			jen.Id("ids"),
		),
	)
}

func (rec *GenTreeRepo) genFuncBatchDeleteByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	return jen.Line().Comment("BatchDeleteByID 根据 ID 批量删除").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchDeleteByID").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id("ids").Index().Id("int64"),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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

func (rec *GenTreeRepo) genFuncUpdateByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("UpdateByID 根据 ID 批量更新").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateByID").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("true")),
		),
	)
}

func (rec *GenTreeRepo) genFuncUpdateNonNilByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("UpdateNonNilByID 根据 ID 更新非空字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateNonNilByID").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Func().Params(jen.Id("f").Id("any")).Params(jen.Id("bool")).Block(jen.Return().Id("f").Op("!=").Id("nil")),
		),
	)
}

func (rec *GenTreeRepo) genFuncUpdateWithFuncByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	return jen.Line().Comment("UpdateWithFuncByID 根据 ID 更新满足函数的字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("UpdateWithFuncByID").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel).Op("*").Add(generated.UseEntity(camel)),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Id("bool")).Block(
		jen.Return().Id("ag").Dot("BatchUpdateWithFuncByID").Call(
			jen.Id("tx"),
			jen.Index().Op("*").Add(generated.UseEntity(camel)).Values(jen.Id(lowerCamel)),
			jen.Id("fn"),
		),
	)
}

func (rec *GenTreeRepo) genFuncBatchUpdateWithFuncByID(table string, columns []generated.Column) jen.Code {
	if !generated.HasColumn(generated.P_KEY, columns) {
		return jen.Null()
	}

	camel := strcase.ToCamel(table)
	lowerCamel := strcase.ToLowerCamel(table)
	var sql0, sql1 string
	sql0 = fmt.Sprintf("UPDATE %s ", table)
	if generated.HasColumn(generated.D_KEY, columns) {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql1 = fmt.Sprintf(" WHERE %s = ? AND %s AND %s;", generated.P_KEY, generated.NS_COND_KEY, generated.UD_COND_KEY)
		} else {
			sql1 = fmt.Sprintf(" WHERE %s = ? AND %s;", generated.P_KEY, generated.UD_COND_KEY)
		}
	} else {
		if generated.HasColumn(generated.NS_KEY, columns) {
			sql1 = fmt.Sprintf(" WHERE %s = ? AND %s;", generated.P_KEY, generated.NS_COND_KEY)
		} else {
			sql1 = fmt.Sprintf(" WHERE %s = ?;", generated.P_KEY)
		}
	}

	return jen.Line().Comment("BatchUpdateWithFuncByID 根据 ID 批量更新满足函数的字段").Line().Func().Params(jen.Id("ag").Op("*").Id("autoGen")).Id("BatchUpdateWithFuncByID").Params(
		jen.Id("tx").Op("*").Add(generated.UseSql("Tx")),
		jen.Id(lowerCamel+"s").Index().Op("*").Add(generated.UseEntity(camel)),
		generated.GenDeclAnonymousFunc(),
	).Params(jen.Id("bool")).Block(
		jen.Id("recorder").Op(":=").Add(generated.UseLogger("AccessLogger")).Call(jen.Id("ag").Dot("ctx")),
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
			jen.Var().Id("sqlBuilder").Add(generated.UseStrings("Builder")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql0)),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Id("fields")),
			jen.Id("sqlBuilder").Dot("WriteString").Call(jen.Lit(sql1)),
			jen.Id("values").Op("=").Id("append").Call(
				jen.Id("values"),
				jen.Id("id"),
				generated.AddNsSingleValue(columns),
			),
			jen.Id("errorHandler").Op(":=").Add(generated.UseUtil("ErrToLogAndPanic")).Call(jen.Id("recorder")),
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

func (rec *GenTreeRepo) GenFile(table string, columns []generated.Column) *jen.File {
	file := jen.NewFile(table)
	file.Add(rec.genInterfaceAutoGen(table, columns))
	file.Add(rec.genStructAutoGen())
	file.Add(rec.genFuncGetDbCtx())
	file.Add(rec.genFuncMapperAll(table, columns))
	file.Add(rec.genFuncMapperNumeric())
	file.Add(rec.genFuncTreeInfoSelectSql(table, columns))
	file.Add(rec.genFuncCalcInsertField(table, columns))
	file.Add(rec.genFuncCalcUpdateField(table, columns))
	file.Add(rec.genFuncInternalSelectNodeByIDs(table, columns))
	file.Add(rec.genFuncInternalDirectInsert(table, columns))
	file.Add(rec.genFuncInternalUpdateNodeInBothWhenInsert(table, columns))
	file.Add(rec.genFuncInternalUpdateNodeInOnlyPrecursorWhenInsert(table, columns))
	file.Add(rec.genFuncInternalInsertWithFunc(table, columns))
	file.Add(rec.genFuncInternalDirectDelete(table, columns))
	file.Add(rec.genFuncInternalUpdateNodeWhenDelete(table, columns))
	file.Add(rec.genFuncSelectByID(table, columns))
	file.Add(rec.genFuncSelectByIDs(table, columns))
	file.Add(rec.genFuncBatchSelectByID(table, columns))
	file.Add(rec.genFuncSelectByName(table, columns))
	file.Add(rec.genFuncSelectMaxLevel(table, columns))
	file.Add(rec.genFuncSelectMaxRight(table, columns))
	file.Add(rec.genFuncSelectMaxLeft(table, columns))
	file.Add(rec.genFuncSelectMaxTreeNo(table, columns))
	file.Add(rec.genFuncSelectAllPosterity(table, columns))
	file.Add(rec.genFuncSelectDirectPosterity(table, columns))
	file.Add(rec.genFuncSelectBrother(table, columns))
	file.Add(rec.genFuncSelectBrotherAndSelf(table, columns))
	file.Add(rec.genFuncSelectAncestorChain(table, columns))
	file.Add(rec.genFuncSelectAncestor(table, columns))
	file.Add(rec.genFuncSelectParent(table, columns))
	file.Add(rec.genFuncSelectByTreeNoAndLevel(table, columns))
	file.Add(rec.genFuncSelectByLevel(table, columns))
	file.Add(rec.genFuncSelectRoot(table, columns))
	file.Add(rec.genFuncSelectLeafOfNodeWithPage(table, columns))
	file.Add(rec.genFuncSelectAllLeafOfNode(table, columns))
	file.Add(rec.genFuncSelectAllRoot(table, columns))
	file.Add(rec.genFuncInsert(table, columns))
	file.Add(rec.genFuncInsertUnderNode(table, columns))
	file.Add(rec.genFuncInsertBetweenNode(table, columns))
	file.Add(rec.genFuncBatchInsert(table, columns))
	file.Add(rec.genFuncBatchInsertUnderNode(table, columns))
	file.Add(rec.genFuncBatchInsertBetweenNode(table, columns))
	file.Add(rec.genFuncInsertNonNil(table, columns))
	file.Add(rec.genFuncInsertNonNilUnderNode(table, columns))
	file.Add(rec.genFuncInsertNonNilBetweenNode(table, columns))
	file.Add(rec.genFuncInsertWithFunc(table, columns))
	file.Add(rec.genFuncInsertWithFuncUnderNode(table, columns))
	file.Add(rec.genFuncInsertWithFuncBetweenNode(table, columns))
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
