package main

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/michaeltelford/echo_reference_project/src/api"
	"github.com/spf13/viper"
)

var (
	host string
	port string
	e    *echo.Echo
)

func init() {
	viper.SetDefault("VERSION", "1")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8000")
	viper.SetDefault("DEBUG", false)

	viper.AutomaticEnv()

	host = viper.GetString("HOST")
	port = addPortColonPrefix(viper.GetString("PORT"))

	e = echo.New()
	e.Debug = viper.GetBool("DEBUG")

	db := sqlx.MustConnect("postgres", buildConnectionString())

	api.DB = db
}

func main() {
	// Ensure you declare your routes without a trailing slash.
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	version := fmt.Sprint("/v", viper.GetString("VERSION"))
	app := e.Group(version)

	api.NewAuthor().InitRoutes(app)
	// TODO: Other resource routes go here...

	e.Logger.Fatal(e.Start(host + port))
}

func addPortColonPrefix(port string) string {
	if strings.HasPrefix(port, ":") {
		return port
	}
	return ":" + port
}

func buildConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
}
