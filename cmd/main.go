package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/routes/index"
	"github.com/rafaelrene/got/routes/user"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	indexHandler := index.IndexHandler{}
	app.GET("/", indexHandler.HandleIndexShow)

	userHandler := user.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Logger.Fatal(app.Start(":1234"))
}
