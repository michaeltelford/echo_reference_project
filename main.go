package main

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/michaeltelford/echo_reference_project/src/api"
)

var (
	host string
	port string
	e    *echo.Echo
)

func init() {
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8000")
	viper.SetDefault("DEBUG", false)

	viper.AutomaticEnv()

	host = viper.GetString("HOST")
	port = addPortColonPrefix(viper.GetString("PORT"))

	e = echo.New()
	e.Debug = viper.GetBool("DEBUG")
}

func main() {
	// Ensure you declare your routes without a trailing Slash
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	app := e.Group("/v1")

	api.NewGreet().InitRoutes(app)
	// TODO: Other resource routes go here...

	e.Logger.Fatal(e.Start(host + port))
}

func addPortColonPrefix(port string) string {
	if strings.HasPrefix(port, ":") {
		return port
	}
	return ":" + port
}
