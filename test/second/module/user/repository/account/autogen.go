package account

import (
	"context"
	"database/sql"
	gin "github.com/gin-gonic/gin"
	constant "metis/config/constant"
	database "metis/database"
	entity "metis/test/second/model/entity"
	util "metis/util"
	logger "metis/util/logger"
	"strings"
)

// iAutoGen 该接口自动生成, 请勿修改
type iAutoGen interface {
	SelectByID(id int64) *entity.Account
	SelectByIDs(ids ...int64) []*entity.Account
	BatchSelectByID(ids []int64) []*entity.Account

	Insert(tx *sql.Tx, account *entity.Account) int64
	InsertNonNil(tx *sql.Tx, account *entity.Account) int64
	InsertWithFunc(tx *sql.Tx, account *entity.Account, fn func(f any) bool) int64
	BatchInsert(tx *sql.Tx, accounts []*entity.Account) []int64
	BatchInsertWithFunc(tx *sql.Tx, accounts []*entity.Account, pid, sid int64, fn func(f any) bool) []int64

	DeleteByID(tx *sql.Tx, id int64) bool
	DeleteByIDs(tx *sql.Tx, ids ...int64) bool
	BatchDeleteByID(tx *sql.Tx, ids []int64) bool

	UpdateByID(tx *sql.Tx, account *entity.Account) bool
	UpdateNonNilByID(tx *sql.Tx, account *entity.Account) bool
	UpdateWithFuncByID(tx *sql.Tx, account *entity.Account, fn func(f any) bool) bool
	BatchUpdateWithFuncByID(tx *sql.Tx, accounts []*entity.Account, fn func(f any) bool) bool
}

// autoGen 该结构体自动生成, 请勿修改
type autoGen struct {
	ctx *gin.Context
}

// getDbCtx 获取 DB 的初始上下文
func (ag *autoGen) getDbCtx() context.Context {
	return context.WithValue(context.Background(), constant.TraceIdKey, ag.ctx.GetString(constant.TraceIdKey))
}

// mapperAll 映射实体的所有字体
func mapperAll() (*entity.Account, []any) {
	var r = &entity.Account{}
	var cs = []any{&r.ID, &r.Title, &r.StartAt, &r.NsId}
	return r, cs
}

// mapperNumeric 映射数值型
func mapperNumeric[T int | int64]() (*T, []any) {
	var r T
	var cs = []any{&r}
	return &r, cs
}

// calcInsertField 计算待插入的字段
func calcInsertField(account *entity.Account, fn func(f any) bool) (string, string, []any) {
	var fields []string
	var values []any
	var places []string
	if fn(account.ID) {
		fields = append(fields, "id")
		places = append(places, "?")
		values = append(values, *account.ID)
	}
	if fn(account.Title) {
		fields = append(fields, "title")
		places = append(places, "?")
		values = append(values, *account.Title)
	}
	if fn(account.StartAt) {
		fields = append(fields, "start_at")
		places = append(places, "?")
		values = append(values, *account.StartAt)
	}
	if fn(account.NsId) {
		fields = append(fields, "ns_id")
		places = append(places, "?")
		values = append(values, *account.NsId)
	}
	return strings.Join(fields, ", "), strings.Join(places, ", "), values
}

// calcUpdateField 计算待更新的字段
func calcUpdateField(account *entity.Account, fn func(f any) bool) (string, []any) {
	var fields []string
	var values []any
	if fn(account.ID) {
		fields = append(fields, "id = ?")
		values = append(values, *account.ID)
	}
	if fn(account.Title) {
		fields = append(fields, "title = ?")
		values = append(values, *account.Title)
	}
	if fn(account.StartAt) {
		fields = append(fields, "start_at = ?")
		values = append(values, *account.StartAt)
	}
	if fn(account.NsId) {
		fields = append(fields, "ns_id = ?")
		values = append(values, *account.NsId)
	}
	return strings.Join(fields, ", "), values
}

// internalSelectByIDs 根据 ID 列表插入节点
func (ag *autoGen) internalSelectByIDs(tx *sql.Tx, db *sql.DB, ids []int64) []*entity.Account {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID 列表: %+v 的数据", ids)
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString("SELECT id, title, start_at, ns_id FROM account WHERE id ")
	if len(ids) == 1 {
		sqlBuilder.WriteString("= ?")
	} else {
		sqlBuilder.WriteString("IN (")
		sqlBuilder.WriteString(util.GenPlaceholder(ids))
		sqlBuilder.WriteString(")")
	}
	sqlBuilder.WriteString("ns_id = ? ")
	sqlBuilder.WriteString("deleted = 0;")
	errorHandler := util.ErrToLogAndPanic(recorder)
	var stmt *sql.Stmt
	var err error
	if tx != nil {
		stmt, err = tx.Prepare(sqlBuilder.String())
		defer util.DeferClose(stmt, errorHandler)
		errorHandler(err)
	} else {
		stmt, err = db.Prepare(sqlBuilder.String())
		defer util.DeferClose(stmt, errorHandler)
		errorHandler(err)
	}
	bindValues := util.ToAnyItems(ids)
	bindValues = append(bindValues, util.GetNsID(ag.ctx))
	rows, err := stmt.QueryContext(ag.getDbCtx(), bindValues...)
	errorHandler(err)
	defer util.DeferClose(rows, errorHandler)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// internalDirectInsert 直接插入
func (ag *autoGen) internalDirectInsert(tx *sql.Tx, account *entity.Account, fn func(f any) bool) int64 {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("插入数据: %+v", *account)
	fields, places, values := calcInsertField(account, fn)
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString("INSERT INTO account(")
	sqlBuilder.WriteString(fields)
	sqlBuilder.WriteString(") VALUES (")
	sqlBuilder.WriteString(places)
	sqlBuilder.WriteString(");")
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := tx.Prepare(sqlBuilder.String())
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err := stmt.ExecContext(ag.getDbCtx(), values...)
	errorHandler(err)
	af, err := result.RowsAffected()
	errorHandler(err)
	id, err := result.LastInsertId()
	errorHandler(err)
	if af == 1 {
		return id
	}
	panic("插入失败")
}

// internalDirectDelete 直接删除(逻辑 or 物理)
func (ag *autoGen) internalDirectDelete(tx *sql.Tx, id int64) bool {
	recorder := logger.AccessLogger(ag.ctx)
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := tx.Prepare("DELETE FROM account WHERE id = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err := stmt.ExecContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	errorHandler(err)
	af, err := result.RowsAffected()
	errorHandler(err)
	if af == 1 {
		return true
	}
	panic("删除错误")
}

// SelectByID 根据 ID 查询
func (ag *autoGen) SelectByID(id int64) *entity.Account {
	ds := ag.BatchSelectByID([]int64{id})
	if len(ds) == 1 {
		return ds[0]
	}
	return nil
}

// SelectByIDs 根据 ID 列表查询
func (ag *autoGen) SelectByIDs(ids ...int64) []*entity.Account {
	ds := ag.BatchSelectByID(ids)
	return ds
}

// BatchSelectByID 根据 ID 批量查询
func (ag *autoGen) BatchSelectByID(ids []int64) []*entity.Account {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID 列表: %+v 的数据", ids)
	db := database.FetchDB()
	return ag.internalSelectByIDs(nil, db, ids)
}

// SelectByName 根据名称查询
func (ag *autoGen) SelectByName(name string) []*entity.Account {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 NAME: %+v 的数据", name)
	db := database.FetchDB()
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := db.Prepare("SELECT id, title, start_at, ns_id FROM account WHERE name like ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	rows, err := stmt.QueryContext(ag.getDbCtx(), name, util.GetNsID(ag.ctx))
	errorHandler(err)
	defer util.DeferClose(rows, errorHandler)
	ds := util.Rows(rows, mapperAll)
	return ds
}
func (ag *autoGen) Insert(tx *sql.Tx, account *entity.Account) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Account{account}, func(f any) bool {
		return true
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertNonNil 插入非空字段
func (ag *autoGen) InsertNonNil(tx *sql.Tx, account *entity.Account) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Account{account}, func(f any) bool {
		return f != nil
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertWithFunc 根据函数插入字段
func (ag *autoGen) InsertWithFunc(tx *sql.Tx, account *entity.Account, fn func(f any) bool) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Account{account}, fn)
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// BatchInsert 批量插入
func (ag *autoGen) BatchInsert(tx *sql.Tx, accounts []*entity.Account) []int64 {
	ids := ag.BatchInsertWithFunc(tx, accounts, func(f any) bool {
		return true
	})
	if len(ids) == len(accounts) {
		return ids
	}
	panic("插入失败, 仅返回记录数等于插入记录数时成功")
}

// BatchInsertWithFunc 根据函数批量插入
func (ag *autoGen) BatchInsertWithFunc(tx *sql.Tx, accounts []*entity.Account, fn func(f any) bool) []int64 {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Info("批量插入数据")
	ids := make([]int64, len(accounts))
	for i, account := range accounts {
		ids[i] = ag.internalDirectInsert(tx, account, fn)
	}
	return ids
}

// DeleteByID 根据 ID 删除
func (ag *autoGen) DeleteByID(tx *sql.Tx, id int64) bool {
	return ag.BatchDeleteByID(tx, []int64{id})
}

// DeleteByIDs 根据 ID 列表删除
func (ag *autoGen) DeleteByIDs(tx *sql.Tx, ids ...int64) bool {
	return ag.BatchDeleteByID(tx, ids)
}

// BatchDeleteByID 根据 ID 批量删除
func (ag *autoGen) BatchDeleteByID(tx *sql.Tx, ids []int64) bool {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("删除 ID 列表: %+v 的数据", ids)
	for _, id := range ids {
		ds := ag.internalDirectDelete(tx, id)
		if !ds {
			panic("存在数据删除错误")
		}
	}
	return true
}

// UpdateByID 根据 ID 批量更新
func (ag *autoGen) UpdateByID(tx *sql.Tx, account *entity.Account) bool {
	return ag.BatchUpdateWithFuncByID(tx, []*entity.Account{account}, func(f any) bool {
		return true
	})
}

// UpdateNonNilByID 根据 ID 更新非空字段
func (ag *autoGen) UpdateNonNilByID(tx *sql.Tx, account *entity.Account) bool {
	return ag.BatchUpdateWithFuncByID(tx, []*entity.Account{account}, func(f any) bool {
		return f != nil
	})
}

// UpdateWithFuncByID 根据 ID 更新满足函数的字段
func (ag *autoGen) UpdateWithFuncByID(tx *sql.Tx, account *entity.Account, fn func(f any) bool) bool {
	return ag.BatchUpdateWithFuncByID(tx, []*entity.Account{account}, fn)
}

// BatchUpdateWithFuncByID 根据 ID 批量更新满足函数的字段
func (ag *autoGen) BatchUpdateWithFuncByID(tx *sql.Tx, accounts []*entity.Account, fn func(f any) bool) bool {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("批量更新列表数据")
	for _, account := range accounts {
		if account.ID == nil {
			panic("ID 字段不能为空")
		}
		id := *account.ID
		fields, values := calcUpdateField(account, fn)
		var sqlBuilder strings.Builder
		sqlBuilder.WriteString("UPDATE account SET ")
		sqlBuilder.WriteString(fields)
		sqlBuilder.WriteString(" WHERE id = ? AND ns_id = ?;")
		values = append(values, id, util.GetNsID(ag.ctx))
		errorHandler := util.ErrToLogAndPanic(recorder)
		stmt, err := tx.Prepare(sqlBuilder.String())
		errorHandler(err)
		result, err := stmt.ExecContext(ag.getDbCtx(), values...)
		errorHandler(err)
		af, err := result.RowsAffected()
		errorHandler(err)
		if af != 1 {
			panic("更新错误")
		}
		err = stmt.Close()
	}
	return true
}
