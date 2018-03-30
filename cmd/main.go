package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger()) // Response logging only
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())

	e.GET("/greet", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"data": "Hello, World!",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
