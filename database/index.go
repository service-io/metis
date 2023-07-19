// Package database
// @author tabuyos
// @since 2023/6/30
// @description database
package database

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/qustavo/sqlhooks/v2"
	"go.uber.org/zap"
	"metis/config"
	"metis/util/logger"
	"time"
)

var db *sql.DB

type Hooks struct{}

// Before hook will print the query with it's args and return the context with the timestamp
func (h *Hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return context.WithValue(ctx, "begin", time.Now()), nil
}

// After hook will get the timestamp registered on the Before hook and print the elapsed time
func (h *Hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	begin := ctx.Value("begin").(time.Time)
	useLogger := logger.UseLogger()
	useLogger.Sugar().Infof("sql: %s, bvs: %v, took: %s", query, args, time.Since(begin))
	return ctx, nil
}

func (h *Hooks) OnError(ctx context.Context, err error, query string, args ...interface{}) error {
	begin := ctx.Value("begin").(time.Time)
	useLogger := logger.UseLogger()
	useLogger.Sugar().Infof("sql: %s, bvs: %v, took: %s", query, args, time.Since(begin))
	return err
}

func init() {
	var driverName = "mysqlWithHooks"

	sql.Register(driverName, sqlhooks.Wrap(&mysql.MySQLDriver{}, &Hooks{}))

	tomlConfig := config.TomlConfig()
	single := tomlConfig.MySQL.Single

	var params = make(map[string]string)
	params["parseTime"] = "true"

	cfg := mysql.Config{
		User:   single.User,
		Passwd: single.Pass,
		Net:    "tcp",
		Addr:   single.Addr,
		DBName: single.Name,
		Params: params,
	}

	var err error
	db, err = sql.Open(driverName, cfg.FormatDSN())
	if err != nil {
		useLogger := logger.UseLogger()
		useLogger.Error(err.Error(), zap.Error(err))
	}
}

func FetchDB() *sql.DB {
	return db
}
