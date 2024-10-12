package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/routes/index"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	indexHandler := index.IndexHandler{}
	app.GET("/", indexHandler.HandleIndexShow)

	app.Logger.Fatal(app.Start(":1234"))
}
