package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	// Author resource.
	Author struct {
		ID     int64  `json:"id" db:"id"`
		Name   string `json:"name" db:"name"`
		Age    int    `json:"age,omitempty" db:"age"`
		Salary int    `json:"-" db:"salary"`
	}
)

// NewAuthor returns an empty Author pointer.
func NewAuthor() *Author {
	return new(Author)
}

// InitRoutes sets up any author based endpoints.
func (u *Author) InitRoutes(group *echo.Group) {
	group.GET("/authors", u.list)
	group.GET("/authors/:id", u.get)
	group.POST("/authors", u.create)
}

func (u *Author) list(c echo.Context) error {
	authors := make([]Author, 0)
	sql := "SELECT * FROM authors"

	if err := DB.Select(&authors, sql); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &authors)
}

func (u *Author) get(c echo.Context) error {
	author := new(Author)
	sql := "SELECT * FROM authors WHERE id = $1 LIMIT 1"

	if err := DB.Get(author, sql, c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, author)
}

func (u *Author) create(c echo.Context) error {
	author := new(Author)
	var err error
	if err = c.Bind(author); err != nil {
		return err
	}

	sql := `INSERT INTO authors (name, age, salary) VALUES ($1, $2, $3)`
	result, err := DB.Exec(sql, author.Name, author.Age, author.Salary)
	if err != nil {
		return err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return err
	}
	author.ID = id

	return c.JSON(http.StatusOK, author)
}
