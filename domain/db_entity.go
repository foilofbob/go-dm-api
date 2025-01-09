package domain

import (
	"database/sql"
	"fmt"
	"go_dm_api/config"
)

func GetByID(tableName string, id int) *sql.Row {
	db := DBConnection()
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)
	return db.QueryRow(query, id)
}

func GetAll(tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
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
	db, err := sql.Open(config.DBDriver(), config.DBConnectString())
	if err != nil {
		panic(err.Error())
	}
	return db
}
