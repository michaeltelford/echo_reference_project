package main

import (
	// https://echo.labstack.com/guide
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// TODO: Find & replace to use your $GOPATH
	"github.com/michaeltelford/echo_reference_project/src/api"
)

func main() {
	e := echo.New()

	// Ensure you declare your routes without a trailing Slash
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	app := e.Group("/v1")

	api.NewGreet().InitRoutes(app)
	// TODO: Other resource routes go here...

	e.Logger.Fatal(e.Start(":8000"))
}
