package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/handler"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Logger.Fatal(app.Start(":1234"))

	fmt.Println("It works!")
}
