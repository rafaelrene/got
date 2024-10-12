package index

import (
	"github.com/labstack/echo/v4"
	"github.com/rafaelrene/got/routes"
)

type IndexHandler struct{}

func (h IndexHandler) HandleIndexShow(c echo.Context) error {
	return routes.Render(c, Show())
}
