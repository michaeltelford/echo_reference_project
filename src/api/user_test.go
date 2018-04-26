package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	DB = newMockDatabase()
	os.Exit(m.Run())
}

func TestNewUser(t *testing.T) {
	assert.IsType(t, &User{}, NewUser())
}

func TestListUsers(t *testing.T) {
	expectedListUsersJSON := `[{"id":1,"name":"Joe Bloggs","age":23},{"id":2,"name":"John Smith","age":37}]`

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := new(User)

	if assert.NoError(t, u.list(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.JSONEq(t, expectedListUsersJSON, rec.Body.String())
	}
}

func TestGetUser(t *testing.T) {
	expectedUserJSON := `{"id":1,"name":"Joe Bloggs","age":23}`

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := new(User)

	if assert.NoError(t, u.get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.JSONEq(t, expectedUserJSON, rec.Body.String())
	}
}

func TestCreateUser(t *testing.T) {
	createUserJSON := `{"name":"Joe Bloggs","age":23,"salary":25000}`
	expectedCreatedUserJSON := `{"id":1,"name":"Joe Bloggs","age":23}`

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/users", strings.NewReader(createUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	u := new(User)

	if assert.NoError(t, u.create(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get(echo.HeaderContentType))
		assert.JSONEq(t, expectedCreatedUserJSON, rec.Body.String())
	}
}
