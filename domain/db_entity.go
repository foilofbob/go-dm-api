package domain

import (
	"database/sql"
	"fmt"
	"go_dm_api/config"

	"github.com/go-sql-driver/mysql"
)

func GetByID(tableName string, id int) *sql.Row {
	db := DBConnection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?", tableName)
	return db.QueryRow(query, id)
}

func GetAll(tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	return DBQuery(query)
}

func DBQuery(queryStr string, args ...any) (*sql.Rows, error) {
	db := DBConnection()
	defer db.Close()
	return db.Query(queryStr, args...)
}

func DBExec(execStr string, args ...any) (sql.Result, error) {
	db := DBConnection()
	defer db.Close()
	return db.Exec(execStr, args...)
}

func DBConnection() *sql.DB {
	cfg := config.Cfg()

	dbcfg := mysql.NewConfig()
	dbcfg.User = cfg.Database.Username
	dbcfg.Passwd = cfg.Database.Password
	dbcfg.Net = "tcp"
	dbcfg.Addr = fmt.Sprintf("%s:%s", cfg.Database.Host, cfg.Database.Port)
	dbcfg.DBName = cfg.Database.DB
	dbcfg.MultiStatements = true // Required for golang-migrate

	db, err := sql.Open(config.DBDriver(), dbcfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	return db
}
