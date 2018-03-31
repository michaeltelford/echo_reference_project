package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	greetJSON = `{"message":"Hello, Joe Bloggs"}`
)

func TestNewGreet(t *testing.T) {
	assert.NotNil(t, NewGreet())
}

func TestListGreeting(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/greet", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	g := new(greet)

	if assert.NoError(t, g.listGreeting(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"message":"Hello, World!"}`, rec.Body.String())
	}
}

func TestListGreetingWithQueryParam(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("name", "Joe Bloggs")
	req, _ := http.NewRequest(echo.GET, "/greet?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	g := new(greet)

	if assert.NoError(t, g.listGreeting(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, greetJSON, rec.Body.String())
	}
}

func TestGetGreeting(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/greet/:name")
	c.SetParamNames("name")
	c.SetParamValues("Joe Bloggs")
	g := new(greet)

	if assert.NoError(t, g.getGreeting(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, greetJSON, rec.Body.String())
	}
}

func TestCreateGreeting(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/greet", strings.NewReader(greetJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	g := new(greet)

	if assert.NoError(t, g.createGreeting(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, greetJSON, rec.Body.String())
	}
}
