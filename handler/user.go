package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return Render(c, user.Show())
}
