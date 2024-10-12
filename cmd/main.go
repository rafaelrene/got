package main

import (
	"database/sql"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/db"
	"github.com/rafaelrene/got/routes/index"
)

func prepareRoutes(app *echo.Echo, db *sql.DB) {
	indexHandler := index.IndexHandler{Db: db}

	app.GET("/", indexHandler.HandleIndexShow)

	app.POST("/add-todo", indexHandler.HandleAddTodo)
}

func main() {
	db, connector, dir := db.GetDbConnection()
	defer os.RemoveAll(dir)
	defer connector.Close()
	defer db.Close()

	app := echo.New()
	app.Static("/static", "static")

	prepareRoutes(app, db)

	app.Logger.Fatal(app.Start(":1234"))
}
