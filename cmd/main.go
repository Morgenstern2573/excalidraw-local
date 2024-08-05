package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	Server *echo.Echo
}

func main() {
	app := &Application{Server: echo.New()}
	app.Server.HideBanner = true
	app.SetupRenderer()
	app.SetupRouting()
	app.Server.Use(middleware.Logger())
	app.Server.Logger.Fatal(app.Server.Start(":4000"))
}
