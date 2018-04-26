package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	// User resource.
	User struct {
		ID     int64  `json:"id" db:"id"`
		Name   string `json:"name" db:"name"`
		Age    int    `json:"age,omitempty" db:"age"`
		Salary int    `json:"-" db:"salary"`
	}
)

// NewUser returns an empty User pointer.
func NewUser() *User {
	return new(User)
}

// InitRoutes sets up any user based endpoints.
func (u *User) InitRoutes(group *echo.Group) {
	group.GET("/users", u.list)
	group.GET("/users/:id", u.get)
	group.POST("/users", u.create)
}

func (u *User) list(c echo.Context) error {
	users := make([]User, 0)
	sql := "SELECT * FROM user"

	if err := DB.Select(&users, sql); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &users)
}

func (u *User) get(c echo.Context) error {
	user := new(User)
	sql := "SELECT * FROM user WHERE id = $1 LIMIT 1"

	if err := DB.Get(user, sql, c.Param("id")); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (u *User) create(c echo.Context) error {
	user := new(User)
	var err error
	if err = c.Bind(user); err != nil {
		return err
	}

	sql := `INSERT INTO user (name, age, salary) VALUES ($1, $2, $3)`
	result, err := DB.Exec(sql, user.Name, user.Age, user.Salary)
	if err != nil {
		return err
	}

	var id int64
	if id, err = result.LastInsertId(); err != nil {
		return err
	}
	user.ID = id

	return c.JSON(http.StatusOK, user)
}
