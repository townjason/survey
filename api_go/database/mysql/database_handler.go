package mysql

import (
	"api/config"
	"api/util/log"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type databaseHandler struct {
	db *sql.DB
	tx *sql.Tx
	isBeginTransaction bool
}

var (
	db *sql.DB
	err error
)

func DatabaseOpen() {
	db, err = sql.Open("mysql",
		""+config.Database.Username+
			":"+config.Database.Password+
			"@tcp("+config.Database.Host+
			":"+config.Database.Port+
			")/"+config.Database.Name+"?parseTime=true&loc=Local")

	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func (_db *databaseHandler) getDb() *sql.DB {
	return _db.db
}

func (_db *databaseHandler) IsBeginTransaction() bool {
	return _db.isBeginTransaction
}

func (_db *databaseHandler) connect() error {
	if db == nil {
		return errors.New("database is nil")
	}
	
	_db.db = db
	return nil
}

func (_db *databaseHandler) Model(s interface{}) Table {
	return newTable("", s, _db, false)
}

func (_db *databaseHandler) getTx() *sql.Tx {
	return _db.tx
}

func (_db *databaseHandler) Transaction(closure func(Database)) {
	if err := _db.connect(); err != nil {
		log.Error(err)
	} else {
		log.Error(_db.beginTransaction())
		_db.isBeginTransaction = true
		closure(_db)
	}
}

func (_db *databaseHandler) beginTransaction() error {
	if _db.tx == nil {
		tx, err := _db.db.Begin()
		if err != nil {
			return err
		}
		_db.tx = tx
	}
	return nil
}

func (_db *databaseHandler) Commit() error {
	return _db.tx.Commit()
}

func (_db *databaseHandler) Rollback() error {
	return _db.tx.Rollback()
}

func (_db *databaseHandler) Option(closure func(Database)) {
	if err := _db.connect(); err != nil {
		log.Error(err)
	} else {
		closure(_db)
	}
}

func (_db *databaseHandler) CloseDb() {
	if _db.db != nil {
		log.Error(_db.db.Close())
	}
}

func (_db *databaseHandler) CloseRows(rows *sql.Rows) {
	if rows != nil {
		log.Error(rows.Close())
	}
}

func (_db *databaseHandler) CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		log.Error(stmt.Close())
	}
}

func (_db *databaseHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return _db.db.Query(query, args...)
}

func (_db *databaseHandler) QueryRow(query string, args ...interface{}) *sql.Row {
	return _db.db.QueryRow(query, args...)
}
