package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Application) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
