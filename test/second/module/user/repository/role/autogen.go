package role

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
	SelectByID(id int64) *entity.Role
	SelectByIDs(ids ...int64) []*entity.Role
	BatchSelectByID(ids []int64) []*entity.Role

	SelectMaxLevel(treeNo int) int
	SelectMaxRight(treeNo int) int
	SelectMaxLeft(treeNo int) int
	SelectMaxTreeNo() int
	SelectAllPosterity(id int64) []*entity.Role
	SelectDirectPosterity(id int64) []*entity.Role
	SelectBrother(id int64) []*entity.Role
	SelectBrotherAndSelf(id int64) []*entity.Role
	SelectAncestorChain(id int64) []*entity.Role
	SelectAncestor(id int64, level int) *entity.Role
	SelectParent(id int64) *entity.Role
	SelectByTreeNoAndLevel(treeNo, level int) []*entity.Role
	SelectByLevel(level int) []*entity.Role
	SelectRoot(id int64) *entity.Role
	SelectLeafOfNodeWithPage(id int64, page, size uint) ([]*entity.Role, int64)
	SelectAllLeafOfNode(id int64) []*entity.Role
	SelectAllRoot() []*entity.Role

	Insert(tx *sql.Tx, role *entity.Role) int64
	InsertUnderNode(tx *sql.Tx, role *entity.Role, pid int64) int64
	InsertBetweenNode(tx *sql.Tx, role *entity.Role, pid, sid int64) int64
	BatchInsert(tx *sql.Tx, roles []*entity.Role) []int64
	BatchInsertUnderNode(tx *sql.Tx, roles []*entity.Role, pid int64) []int64
	BatchInsertBetweenNode(tx *sql.Tx, roles []*entity.Role, pid, sid int64) []int64
	InsertNonNil(tx *sql.Tx, role *entity.Role) int64
	InsertNonNilUnderNode(tx *sql.Tx, role *entity.Role, pid int64) int64
	InsertNonNilBetweenNode(tx *sql.Tx, role *entity.Role, pid, sid int64) int64
	InsertWithFunc(tx *sql.Tx, role *entity.Role, fn func(f any) bool) int64
	InsertWithFuncUnderNode(tx *sql.Tx, role *entity.Role, pid int64, fn func(f any) bool) int64
	InsertWithFuncBetweenNode(tx *sql.Tx, role *entity.Role, pid, sid int64, fn func(f any) bool) int64
	BatchInsertWithFunc(tx *sql.Tx, roles []*entity.Role, pid, sid int64, fn func(f any) bool) []int64

	DeleteByID(tx *sql.Tx, id int64) bool
	DeleteByIDs(tx *sql.Tx, ids ...int64) bool
	BatchDeleteByID(tx *sql.Tx, ids []int64) bool

	UpdateByID(tx *sql.Tx, role *entity.Role) bool
	UpdateNonNilByID(tx *sql.Tx, role *entity.Role) bool
	UpdateWithFuncByID(tx *sql.Tx, role *entity.Role, fn func(f any) bool) bool
	BatchUpdateWithFuncByID(tx *sql.Tx, roles []*entity.Role, fn func(f any) bool) bool
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
func mapperAll() (*entity.Role, []any) {
	var r = &entity.Role{}
	var cs = []any{&r.ID, &r.Title, &r.StartAt, &r.NsId}
	return r, cs
}

// mapperNumeric 映射数值型
func mapperNumeric[T int | int64]() (*T, []any) {
	var r T
	var cs = []any{&r}
	return &r, cs
}

// treeInfoSelectSql 获取树型表的基础信息
func treeInfoSelectSql() string {
	return "SELECT left, right, level, tree_no FROM role WHERE id = ? AND ns_id = ?;"
}

// calcInsertField 计算待插入的字段
func calcInsertField(role *entity.Role, fn func(f any) bool) (string, string, []any) {
	var fields []string
	var values []any
	var places []string
	if fn(role.ID) {
		fields = append(fields, "id")
		places = append(places, "?")
		values = append(values, *role.ID)
	}
	if fn(role.Title) {
		fields = append(fields, "title")
		places = append(places, "?")
		values = append(values, *role.Title)
	}
	if fn(role.StartAt) {
		fields = append(fields, "start_at")
		places = append(places, "?")
		values = append(values, *role.StartAt)
	}
	if fn(role.NsId) {
		fields = append(fields, "ns_id")
		places = append(places, "?")
		values = append(values, *role.NsId)
	}
	return strings.Join(fields, ", "), strings.Join(places, ", "), values
}

// calcUpdateField 计算待更新的字段
func calcUpdateField(role *entity.Role, fn func(f any) bool) (string, []any) {
	var fields []string
	var values []any
	if fn(role.ID) {
		fields = append(fields, "SET id = ?")
		values = append(values, *role.ID)
	}
	if fn(role.Title) {
		fields = append(fields, "SET title = ?")
		values = append(values, *role.Title)
	}
	if fn(role.StartAt) {
		fields = append(fields, "SET start_at = ?")
		values = append(values, *role.StartAt)
	}
	if fn(role.NsId) {
		fields = append(fields, "SET ns_id = ?")
		values = append(values, *role.NsId)
	}
	return strings.Join(fields, ", "), values
}

// internalSelectNodeByIDs 根据 ID 列表插入节点
func (ag *autoGen) internalSelectNodeByIDs(tx *sql.Tx, db *sql.DB, ids []int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID 列表: %+v 的数据", ids)
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString("SELECT id, title, start_at, ns_id FROM role WHERE id ")
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

// internalDirectInsert 直接插入树节点, 需要提前计算好树相关信息
func (ag *autoGen) internalDirectInsert(tx *sql.Tx, role *entity.Role, fn func(f any) bool) int64 {
	recorder := logger.AccessLogger(ag.ctx)
	if role.TreeNo == nil {
		panic("需要填充树号")
	}
	if role.Left == nil {
		panic("需要填充左值")
	}
	if role.Right == nil {
		panic("需要填充右值")
	}
	if role.Level == nil {
		panic("需要填充层级")
	}
	fields, places, values := calcInsertField(role, fn)
	var sqlBuilder strings.Builder
	sqlBuilder.WriteString("INSERT INTO role(")
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

// internalUpdateNodeInBothWhenInsert 在两个节点间插入时更新
func (ag *autoGen) internalUpdateNodeInBothWhenInsert(tx *sql.Tx, left, right, treeNo int) {
	recorder := logger.AccessLogger(ag.ctx)
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := tx.Prepare("UPDATE role SET left = left + 2 WHERE left > ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err := stmt.ExecContext(ag.getDbCtx(), right, treeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	_, err = result.RowsAffected()
	errorHandler(err)
	stmt, err = tx.Prepare("UPDATE role SET right = right + 2 WHERE right > ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err = stmt.ExecContext(ag.getDbCtx(), right, treeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	_, err = result.RowsAffected()
	errorHandler(err)
	stmt, err = tx.Prepare("UPDATE role SET left = left + 1, right = right + 1, level = level + 1 WHERE level >= ? AND right <= ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err = stmt.ExecContext(ag.getDbCtx(), left, right, treeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	_, err = result.RowsAffected()
	errorHandler(err)
}

// internalUpdateNodeInOnlyPrecursorWhenInsert 插入至前驱节点时更新
func (ag *autoGen) internalUpdateNodeInOnlyPrecursorWhenInsert(tx *sql.Tx, right, treeNo int) {
	recorder := logger.AccessLogger(ag.ctx)
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := tx.Prepare("UPDATE role SET left = left + 2 WHERE left > ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err := stmt.ExecContext(ag.getDbCtx(), right, treeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	_, err = result.RowsAffected()
	errorHandler(err)
	stmt, err = tx.Prepare("UPDATE role SET right = right + 2 WHERE right >= ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err = stmt.ExecContext(ag.getDbCtx(), right, treeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	_, err = result.RowsAffected()
	errorHandler(err)
}

// internalInsertWithFunc 根据函数进行插入
func (ag *autoGen) internalInsertWithFunc(tx *sql.Tx, role *entity.Role, pid, sid int64, fn func(f any) bool) int64 {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("插入节点, 前驱: %+v, 后继: %+v, 节点: %+v", pid, sid, *role)
	if pid == 0 {
		return ag.internalDirectInsert(tx, role, fn)
	}
	precursorNodes := ag.internalSelectNodeByIDs(tx, nil, []int64{pid})
	nodeLen := len(precursorNodes)
	if nodeLen == 1 {
		precursor := precursorNodes[0]
		role.TreeNo = precursor.TreeNo
		level := *precursor.Level + 1
		role.Level = &level
		if sid == 0 {
			right := *precursor.Right + 1
			role.Left = precursor.Right
			role.Right = &right
			ag.internalUpdateNodeInOnlyPrecursorWhenInsert(tx, *precursor.Right, *precursor.TreeNo)
			return ag.internalDirectInsert(tx, role, fn)
		} else {
			successorNodes := ag.internalSelectNodeByIDs(tx, nil, []int64{pid})
			if len(successorNodes) == 1 {
				successor := successorNodes[0]
				right := *successor.Right + 2
				role.Left = successor.Left
				role.Right = &right
				ag.internalUpdateNodeInBothWhenInsert(tx, *successor.Left, *successor.Right, *successor.TreeNo)
				return ag.internalDirectInsert(tx, role, fn)
			}
			panic("存在多个或不存在后继节点")
		}
	}
	panic("存在多个或不存在前驱节点")
}

// internalDirectDelete 直接删除(逻辑 or 物理)
func (ag *autoGen) internalDirectDelete(tx *sql.Tx, id int64) bool {
	recorder := logger.AccessLogger(ag.ctx)
	nodes := ag.internalSelectNodeByIDs(tx, nil, []int64{id})
	if len(nodes) == 1 {
		node := nodes[0]
		right := *node.Right
		left := *node.Left
		treeNo := *node.TreeNo
		delta := right - left + 1
		ag.internalUpdateNodeWhenDelete(tx, delta, right, treeNo)
		errorHandler := util.ErrToLogAndPanic(recorder)
		stmt, err := tx.Prepare("DELETE FROM role WHERE left >= ? AND right <= ? AND tree_no = ? AND ns_id = ?;")
		defer util.DeferClose(stmt, errorHandler)
		errorHandler(err)
		result, err := stmt.ExecContext(ag.getDbCtx(), left, right, treeNo, util.GetNsID(ag.ctx))
		errorHandler(err)
		af, err := result.RowsAffected()
		errorHandler(err)
		if af == 1 {
			return true
		}
		panic("删除错误")
	}
	panic("节点数错误")
}

// internalUpdateNodeWhenDelete 删除时更新节点
func (ag *autoGen) internalUpdateNodeWhenDelete(tx *sql.Tx, delta, right, treeNo int) {
	recorder := logger.AccessLogger(ag.ctx)
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := tx.Prepare("UPDATE role SET left = left - ? WHERE left > ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err := stmt.ExecContext(ag.getDbCtx(), delta, right, treeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	_, err = result.RowsAffected()
	errorHandler(err)
	stmt, err = tx.Prepare("UPDATE role SET right = right - ? WHERE right > ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	result, err = stmt.ExecContext(ag.getDbCtx(), delta, right, treeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	_, err = result.RowsAffected()
	errorHandler(err)
}

// SelectByID 根据 ID 查询
func (ag *autoGen) SelectByID(id int64) *entity.Role {
	ds := ag.BatchSelectByID([]int64{id})
	if len(ds) == 1 {
		return ds[0]
	}
	return nil
}

// SelectByIDs 根据 ID 列表查询
func (ag *autoGen) SelectByIDs(ids ...int64) []*entity.Role {
	ds := ag.BatchSelectByID(ids)
	return ds
}

// BatchSelectByID 根据 ID 批量查询
func (ag *autoGen) BatchSelectByID(ids []int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID 列表: %+v 的数据", ids)
	db := database.FetchDB()
	return ag.internalSelectNodeByIDs(nil, db, ids)
}

// SelectByName 根据名称查询
func (ag *autoGen) SelectByName(name string) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 NAME: %+v 的数据", name)
	db := database.FetchDB()
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := db.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE name like ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	rows, err := stmt.QueryContext(ag.getDbCtx(), name, util.GetNsID(ag.ctx))
	errorHandler(err)
	defer util.DeferClose(rows, errorHandler)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectMaxLevel 查询最大层级
func (ag *autoGen) SelectMaxLevel(treeNo int) int {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 TN: %+v 的最大层级", treeNo)
	db := database.FetchDB()
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := db.Prepare("SELECT MAX(level) FROM role WHERE tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	row := stmt.QueryRowContext(ag.getDbCtx(), treeNo, util.GetNsID(ag.ctx))
	ds := util.Row(row, mapperNumeric[int])
	return *ds
}

// SelectMaxRight 查询最大右值
func (ag *autoGen) SelectMaxRight(treeNo int) int {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 TN: %+v 的最大层级", treeNo)
	db := database.FetchDB()
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := db.Prepare("SELECT MAX(right) FROM role WHERE tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	row := stmt.QueryRowContext(ag.getDbCtx(), treeNo, util.GetNsID(ag.ctx))
	ds := util.Row(row, mapperNumeric[int])
	return *ds
}

// SelectMaxLeft 查询最大左值
func (ag *autoGen) SelectMaxLeft(treeNo int) int {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 TN: %+v 的最大层级", treeNo)
	db := database.FetchDB()
	errorHandler := util.ErrToLogAndPanic(recorder)
	stmt, err := db.Prepare("SELECT MAX(left) FROM role WHERE tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	row := stmt.QueryRowContext(ag.getDbCtx(), treeNo, util.GetNsID(ag.ctx))
	ds := util.Row(row, mapperNumeric[int])
	return *ds
}

// SelectMaxTreeNo 查询最大树号
func (ag *autoGen) SelectMaxTreeNo() int {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Info("查询最大TN")
	db := database.FetchDB()
	row := db.QueryRowContext(ag.getDbCtx(), "SELECT MAX(tree_no) FROM role WHERE ns_id = ?;", util.GetNsID(ag.ctx))
	ds := util.Row(row, mapperNumeric[int])
	return *ds
}

// SelectAllPosterity 查询所有子代
func (ag *autoGen) SelectAllPosterity(id int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的所有子代(含自身)数据", id)
	// recorder.Warn("不建议查询全部子代, 如果树比较大, 数据量将会非常大")
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE left > ? AND right < ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	rows, err := secondStmt.QueryContext(ag.getDbCtx(), *currentNode.Left, *currentNode.Right, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectDirectPosterity 查询直系子代
func (ag *autoGen) SelectDirectPosterity(id int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的直系子代数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE level = ? AND left > ? AND right < ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	rows, err := secondStmt.QueryContext(ag.getDbCtx(), *currentNode.Level+1, *currentNode.Left, *currentNode.Right, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	errorHandler(err)
	defer util.DeferClose(rows, errorHandler)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectBrother 查询兄弟(不含自身)
func (ag *autoGen) SelectBrother(id int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的兄弟数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE level = ? AND tree_no = ? AND id != ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	rows, err := secondStmt.QueryContext(ag.getDbCtx(), *currentNode.Level+1, *currentNode.TreeNo, id, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectBrotherAndSelf 查询兄弟和自身
func (ag *autoGen) SelectBrotherAndSelf(id int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的兄弟以及自身数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE level = ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	rows, err := secondStmt.QueryContext(ag.getDbCtx(), *currentNode.Level+1, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectAncestorChain 查询祖链
func (ag *autoGen) SelectAncestorChain(id int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的祖链数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE left < ? AND right > ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	rows, err := secondStmt.QueryContext(ag.getDbCtx(), *currentNode.Left, *currentNode.Right, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectAncestor 查询祖节点
func (ag *autoGen) SelectAncestor(id int64, level int) *entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的祖代(%+v)数据", id, level)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE left < ? AND right > ? AND level = ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	row = secondStmt.QueryRowContext(ag.getDbCtx(), *currentNode.Left, *currentNode.Right, level, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	ds := util.Row(row, mapperAll)
	return ds
}

// SelectParent 查询父节点
func (ag *autoGen) SelectParent(id int64) *entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的父节点数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE left < ? AND right > ? AND level = ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	row = secondStmt.QueryRowContext(ag.getDbCtx(), *currentNode.Left, *currentNode.Right, *currentNode.Level-1, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	ds := util.Row(row, mapperAll)
	return ds
}

// SelectByTreeNoAndLevel 根据树号和层级查询
func (ag *autoGen) SelectByTreeNoAndLevel(treeNo, level int) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 TN: %+v LL: %+v 的同代数据", treeNo, level)
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	stmt, err := db.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE level = ? AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	rows, err := stmt.QueryContext(ag.getDbCtx(), treeNo, level, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectByLevel 根据层级查询
func (ag *autoGen) SelectByLevel(level int) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 LL: %+v 的同代(跨树)数据", level)
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	stmt, err := db.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE level = ? AND ns_id = ?;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	rows, err := stmt.QueryContext(ag.getDbCtx(), level, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectRoot 查询根节点
func (ag *autoGen) SelectRoot(id int64) *entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的根节点数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE level = 1 AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	row = secondStmt.QueryRowContext(ag.getDbCtx(), *currentNode.TreeNo, util.GetNsID(ag.ctx))
	ds := util.Row(row, mapperAll)
	return ds
}

// SelectLeafOfNodeWithPage 查询对应节点叶子节点(分页)
func (ag *autoGen) SelectLeafOfNodeWithPage(id int64, page, size uint) ([]*entity.Role, int64) {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("分页查询 ID: %+v 的叶子节点数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE left >= ? AND right <= ? AND left + 1 = right AND tree_no = ? AND ns_id = ? ORDER BY left LIMIT ? OFFSET ?;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	rows, err := secondStmt.QueryContext(ag.getDbCtx(), *currentNode.Left, *currentNode.Right, *currentNode.TreeNo, size, (page-1)*size, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	thirdStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE left >= ? AND right <= ? AND left + 1 = right AND tree_no = ? AND ns_id = ?;")
	defer util.DeferClose(thirdStmt, errorHandler)
	errorHandler(err)
	row = thirdStmt.QueryRowContext(ag.getDbCtx(), *currentNode.Left, *currentNode.Right, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	total := util.Row(row, mapperNumeric[int64])
	return ds, *total
}

// SelectAllLeafOfNode 查询对应节点的所有叶子节点
func (ag *autoGen) SelectAllLeafOfNode(id int64) []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("查询 ID: %+v 的所有叶子节点数据", id)
	treeInfoSql := treeInfoSelectSql()
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	tx, err := db.Begin()
	defer util.HandleTx(tx, errorHandler)
	errorHandler(err)
	firstStmt, err := tx.Prepare(treeInfoSql)
	defer util.DeferClose(firstStmt, errorHandler)
	errorHandler(err)
	row := firstStmt.QueryRowContext(ag.getDbCtx(), id, util.GetNsID(ag.ctx))
	currentNode := util.Row(row, mapperAll)
	secondStmt, err := tx.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE left >= ? AND right <= ? AND left + 1 = right AND tree_no = ? AND ns_id = ? ORDER BY left;")
	defer util.DeferClose(secondStmt, errorHandler)
	errorHandler(err)
	rows, err := secondStmt.QueryContext(ag.getDbCtx(), *currentNode.Left, *currentNode.Right, *currentNode.TreeNo, util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}

// SelectAllRoot 查询所有的根节点
func (ag *autoGen) SelectAllRoot() []*entity.Role {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Info("查询的所有根节点数据")
	errorHandler := util.ErrToLogAndPanic(recorder)
	db := database.FetchDB()
	stmt, err := db.Prepare("SELECT id, title, start_at, ns_id FROM role WHERE level = 1 AND ns_id = ? ORDER BY tree_no;")
	defer util.DeferClose(stmt, errorHandler)
	errorHandler(err)
	rows, err := stmt.QueryContext(ag.getDbCtx(), util.GetNsID(ag.ctx))
	defer util.DeferClose(rows, errorHandler)
	errorHandler(err)
	ds := util.Rows(rows, mapperAll)
	return ds
}
func (ag *autoGen) Insert(tx *sql.Tx, role *entity.Role) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, 0, 0, func(f any) bool {
		return true
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertUnderNode 插入至节点下方(做叶子节点)
func (ag *autoGen) InsertUnderNode(tx *sql.Tx, role *entity.Role, pid int64) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, pid, 0, func(f any) bool {
		return true
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertBetweenNode 插入至两节点间
func (ag *autoGen) InsertBetweenNode(tx *sql.Tx, role *entity.Role, pid, sid int64) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, pid, sid, func(f any) bool {
		return true
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// BatchInsert 批量插入
func (ag *autoGen) BatchInsert(tx *sql.Tx, roles []*entity.Role) []int64 {
	ids := ag.BatchInsertWithFunc(tx, roles, 0, 0, func(f any) bool {
		return true
	})
	if len(ids) == len(roles) {
		return ids
	}
	panic("插入失败, 仅返回记录数等于插入记录数时成功")
}

// BatchInsertUnderNode 批量插入至节点下方
func (ag *autoGen) BatchInsertUnderNode(tx *sql.Tx, roles []*entity.Role, pid int64) []int64 {
	ids := ag.BatchInsertWithFunc(tx, roles, pid, 0, func(f any) bool {
		return true
	})
	if len(ids) == len(roles) {
		return ids
	}
	panic("插入失败, 仅返回记录数等于插入记录数时成功")
}

// BatchInsertBetweenNode 批量插入至两节点间(谨慎使用)
func (ag *autoGen) BatchInsertBetweenNode(tx *sql.Tx, roles []*entity.Role, pid, sid int64) []int64 {
	ids := ag.BatchInsertWithFunc(tx, roles, pid, sid, func(f any) bool {
		return true
	})
	if len(ids) == len(roles) {
		return ids
	}
	panic("插入失败, 仅返回记录数等于插入记录数时成功")
}

// InsertNonNil 插入非空字段
func (ag *autoGen) InsertNonNil(tx *sql.Tx, role *entity.Role) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, 0, 0, func(f any) bool {
		return f != nil
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertNonNilUnderNode 插入非空字段并挂载到某节点下方
func (ag *autoGen) InsertNonNilUnderNode(tx *sql.Tx, role *entity.Role, pid int64) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, pid, 0, func(f any) bool {
		return f != nil
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertNonNilBetweenNode 插入非空字段并挂载到两节点之间
func (ag *autoGen) InsertNonNilBetweenNode(tx *sql.Tx, role *entity.Role, pid, sid int64) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, pid, sid, func(f any) bool {
		return f != nil
	})
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertWithFunc 根据函数插入字段
func (ag *autoGen) InsertWithFunc(tx *sql.Tx, role *entity.Role, fn func(f any) bool) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, 0, 0, fn)
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertWithFuncUnderNode 根据函数插入并挂载到某节点下方
func (ag *autoGen) InsertWithFuncUnderNode(tx *sql.Tx, role *entity.Role, pid int64, fn func(f any) bool) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, pid, 0, fn)
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// InsertWithFuncBetweenNode 根据函数插入并挂载到两节点之间
func (ag *autoGen) InsertWithFuncBetweenNode(tx *sql.Tx, role *entity.Role, pid, sid int64, fn func(f any) bool) int64 {
	ids := ag.BatchInsertWithFunc(tx, []*entity.Role{role}, pid, sid, fn)
	if len(ids) == 1 {
		return ids[0]
	}
	panic("插入失败, 仅返回一条记录时成功")
}

// BatchInsertWithFunc 根据函数批量插入
func (ag *autoGen) BatchInsertWithFunc(tx *sql.Tx, roles []*entity.Role, pid, sid int64, fn func(f any) bool) []int64 {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("插入至 PID: %+v SID: %+v 的同代数据", pid, sid)
	ids := make([]int64, len(roles))
	for i, role := range roles {
		ids[i] = ag.internalInsertWithFunc(tx, role, pid, sid, fn)
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
func (ag *autoGen) UpdateByID(tx *sql.Tx, role *entity.Role) bool {
	return ag.BatchUpdateWithFuncByID(tx, []*entity.Role{role}, func(f any) bool {
		return true
	})
}

// UpdateNonNilByID 根据 ID 更新非空字段
func (ag *autoGen) UpdateNonNilByID(tx *sql.Tx, role *entity.Role) bool {
	return ag.BatchUpdateWithFuncByID(tx, []*entity.Role{role}, func(f any) bool {
		return f != nil
	})
}

// UpdateWithFuncByID 根据 ID 更新满足函数的字段
func (ag *autoGen) UpdateWithFuncByID(tx *sql.Tx, role *entity.Role, fn func(f any) bool) bool {
	return ag.BatchUpdateWithFuncByID(tx, []*entity.Role{role}, fn)
}

// BatchUpdateWithFuncByID 根据 ID 批量更新满足函数的字段
func (ag *autoGen) BatchUpdateWithFuncByID(tx *sql.Tx, roles []*entity.Role, fn func(f any) bool) bool {
	recorder := logger.AccessLogger(ag.ctx)
	recorder.Sugar().Infof("批量更新列表数据")
	for _, role := range roles {
		if role.ID == nil {
			panic("ID 字段不能为空")
		}
		id := *role.ID
		fields, values := calcUpdateField(role, fn)
		var sqlBuilder strings.Builder
		sqlBuilder.WriteString("UPDATE role ")
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
