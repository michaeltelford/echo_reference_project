package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthor(t *testing.T) {
	assert.IsType(t, &Author{}, NewAuthor())
}

func TestListAuthors(t *testing.T) {
	expectedListAuthorsJSON := `[{"id":1,"name":"Joe Bloggs","age":23},{"id":2,"name":"John Smith","age":37}]`

	// Mock the Database func Select which is called by api's list func.
	f := &FakeDatabase{
		SelectHook: func(dest interface{}, query string, args ...interface{}) (ident1 error) {
			v := reflect.ValueOf(dest).Elem()
			v.Set(reflect.Append(v, reflect.ValueOf(Author{
				ID:     1,
				Name:   "Joe Bloggs",
				Age:    23,
				Salary: 35000,
			}), reflect.ValueOf(Author{
				ID:     2,
				Name:   "John Smith",
				Age:    37,
				Salary: 150000,
			})))
			return
		},
	}
	DB = f

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/authors", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := new(Author)

	if assert.NoError(t, u.list(c)) {
		f.AssertSelectCalledOnce(t)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.JSONEq(t, expectedListAuthorsJSON, rec.Body.String())
	}
}

func TestGetAuthor(t *testing.T) {
	expectedAuthorJSON := `{"id":1,"name":"Joe Bloggs","age":23}`

	// Mock the Database func Get which is called by api's get func.
	f := &FakeDatabase{
		GetHook: func(dest interface{}, query string, args ...interface{}) (ident1 error) {
			s := reflect.ValueOf(dest).Elem()
			s.FieldByName("ID").SetInt(1)
			s.FieldByName("Name").SetString("Joe Bloggs")
			s.FieldByName("Age").SetInt(23)
			s.FieldByName("Salary").SetInt(35000)
			return
		},
	}
	DB = f

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/authors/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := new(Author)

	if assert.NoError(t, u.get(c)) {
		f.AssertGetCalledOnce(t)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.JSONEq(t, expectedAuthorJSON, rec.Body.String())
	}
}

func TestCreateAuthor(t *testing.T) {
	createAuthorJSON := `{"name":"Joe Bloggs","age":23,"salary":25000}`
	expectedCreatedAuthorJSON := `{"id":1,"name":"Joe Bloggs","age":23}`

	// Mock the Database func Exec which is called by api's create func.
	f := &FakeDatabase{
		ExecHook: func(query string, args ...interface{}) (ident1 sql.Result, ident2 error) {
			ident1 = &FakeSQLResult{
				LastInsertIdHook: func() (ident1 int64, ident2 error) {
					ident1 = 1
					return
				},
				RowsAffectedHook: func() (ident1 int64, ident2 error) {
					ident1 = 1
					return
				},
			}
			return
		},
	}
	DB = f

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/authors", strings.NewReader(createAuthorJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := new(Author)

	if assert.NoError(t, u.create(c)) {
		f.AssertExecCalledOnce(t)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.JSONEq(t, expectedCreatedAuthorJSON, rec.Body.String())
	}
}
