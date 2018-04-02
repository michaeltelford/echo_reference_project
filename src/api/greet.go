package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	// TODO: List any attributes here e.g. JSON, DB table columns etc.
	greet struct {
		Message string `json:"message"`
	}
)

// NewGreet returns a greet pointer
func NewGreet() *greet {
	return new(greet)
}

// Initialise any greeting routes.
func (g *greet) InitRoutes(group *echo.Group) {
	group.GET("/greet", g.list)
	group.GET("/greet/:name", g.get)
	group.POST("/greet", g.create)
}

// list takes a query param or default value and builds a message.
func (g *greet) list(c echo.Context) error {
	var name string
	if name = c.QueryParam("name"); name == "" {
		name = "World!"
	}

	g.Message = fmt.Sprintf("Hello, %s", name)
	return c.JSON(http.StatusOK, g)
}

// get takes a path param and builds a message.
func (g *greet) get(c echo.Context) error {
	g.Message = fmt.Sprintf("Hello, %s", c.Param("name"))
	return c.JSON(http.StatusOK, g)
}

// create takes a json req body and builds a message.
func (g *greet) create(c echo.Context) error {
	if err := c.Bind(g); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, g)
}
