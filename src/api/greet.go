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

func NewGreet() *greet {
	return new(greet)
}

func (g *greet) InitRoutes(group *echo.Group) {
	group.GET("/greet", g.listGreeting)
	group.GET("/greet/:name", g.getGreeting)
	group.POST("/greet", g.createGreeting)
}

// listGreeting takes a query param or default value and builds a message.
func (g *greet) listGreeting(c echo.Context) error {
	var name string
	if name = c.QueryParam("name"); name == "" {
		name = "World!"
	}

	g.Message = fmt.Sprintf("Hello, %s", name)
	return c.JSON(http.StatusOK, g)
}

// getGreeting takes a path param and builds a message.
func (g *greet) getGreeting(c echo.Context) error {
	g.Message = fmt.Sprintf("Hello, %s", c.Param("name"))
	return c.JSON(http.StatusOK, g)
}

// createGreeting takes a json req body and builds a message.
func (g *greet) createGreeting(c echo.Context) error {
	if err := c.Bind(g); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, g)
}
