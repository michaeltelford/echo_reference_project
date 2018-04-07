package api

import (
	"database/sql"
	"reflect"
)

type mockResult struct {
	ID   int64
	Rows int64
}

func (result mockResult) LastInsertId() (int64, error) {
	return result.ID, nil
}

func (result mockResult) RowsAffected() (int64, error) {
	return result.Rows, nil
}

// mockDatabase implements the api.Database interface.
type mockDatabase struct{}

func (db *mockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return mockResult{
		ID:   1,
		Rows: 1,
	}, nil
}

func (db *mockDatabase) Get(dest interface{}, query string, args ...interface{}) error {
	s := reflect.ValueOf(dest).Elem()

	s.FieldByName("ID").SetInt(1)
	s.FieldByName("Name").SetString("Joe Bloggs")
	s.FieldByName("Age").SetInt(23)
	s.FieldByName("Salary").SetInt(35000)

	return nil
}

func (db *mockDatabase) Select(dest interface{}, query string, args ...interface{}) error {
	v := reflect.ValueOf(dest).Elem()

	v.Set(reflect.Append(v, reflect.ValueOf(User{
		ID:     1,
		Name:   "Joe Bloggs",
		Age:    23,
		Salary: 35000,
	}), reflect.ValueOf(User{
		ID:     2,
		Name:   "John Smith",
		Age:    37,
		Salary: 150000,
	})))

	return nil
}

func newMockDatabase() *mockDatabase {
	return &mockDatabase{}
}
