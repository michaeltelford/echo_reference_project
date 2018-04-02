package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type (
	// Greet resource containing attributes e.g. JSON, DB table columns etc.
	Greet struct {
		Message string `json:"message"`
	}
)

// NewGreet returns a Greet pointer
func NewGreet() *Greet {
	return new(Greet)
}

// InitRoutes sets up routes for the Greet resource.
func (g *Greet) InitRoutes(group *echo.Group) {
	group.GET("/greet", g.list)
	group.GET("/greet/:name", g.get)
	group.POST("/greet", g.create)
}

// list takes a query param or default value and builds a message.
func (g *Greet) list(c echo.Context) error {
	var name string
	if name = c.QueryParam("name"); name == "" {
		name = "World!"
	}

	g.Message = fmt.Sprintf("Hello, %s", name)
	return c.JSON(http.StatusOK, g)
}

// get takes a path param and builds a message.
func (g *Greet) get(c echo.Context) error {
	g.Message = fmt.Sprintf("Hello, %s", c.Param("name"))
	return c.JSON(http.StatusOK, g)
}

// create takes a json req body and builds a message.
func (g *Greet) create(c echo.Context) error {
	if err := c.Bind(g); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, g)
}
