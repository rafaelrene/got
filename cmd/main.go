package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/routes/index"
)

func prepareRoutes(app *echo.Echo) {
	indexHandler := index.IndexHandler{}
	app.GET("/", indexHandler.HandleIndexShow)
}

func main() {
	app := echo.New()
	app.Static("/static", "static")

	prepareRoutes(app)

	app.Logger.Fatal(app.Start(":1234"))
}
