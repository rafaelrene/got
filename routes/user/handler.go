package user

import (
	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/routes"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return routes.Render(c, Show())
}
