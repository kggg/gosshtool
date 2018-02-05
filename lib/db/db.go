package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func checkerr(err error) {
	if err != nil {
		log.Panicln(err.Error())
	}
}

func New(host string, user string, pass string, port int, dbname string) (*sql.DB, error) {
	consql := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	db, err := sql.Open("mysql", consql)
	checkerr(err)
	return db, err
}

func CheckExists(db *sql.DB, query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}

//fetch data from database
func FetchRows(db *sql.DB, sqlstr string, args ...interface{}) ([]map[string]string, error) {
	stmtOut, err := db.Prepare(sqlstr)
	checkerr(err)
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	checkerr(err)
	columns, err := rows.Columns()
	checkerr(err)
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		checkerr(err)
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return ret, nil
}

// for data update and delete
func Execute(db *sql.DB, sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := db.Prepare(sqlstr)
	checkerr(err)
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	checkerr(err)
	return result.RowsAffected()
}

// for insert data to database
func Insert(db *sql.DB, sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := db.Prepare(sqlstr)
	checkerr(err)
	defer stmtIns.Close()
	result, err := stmtIns.Exec(args...)
	checkerr(err)
	return result.LastInsertId()
}
