package api

import (
	"database/sql"
)

var (
	// DB is a DB handle set at the api package level.
	DB Database
)

type (
	// Database handle interface for easy mocking in tests.
	Database interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		Get(dest interface{}, query string, args ...interface{}) error
		Select(dest interface{}, query string, args ...interface{}) error
	}
)
