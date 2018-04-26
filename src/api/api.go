//go:generate charlatan Database SQLResult
package api

import "database/sql"

// DB is a DB handle set at the api package level.
var DB Database

// Database handle interface sqlx package.
type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// SQLResult interface matching database/sql package.
type SQLResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
