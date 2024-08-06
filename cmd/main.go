package main

import (
	"database/sql"
	"log"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	Server *echo.Echo
}

func main() {
	db, err := sql.Open("sqlite3", "./data/xdraw-local.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	services.Init(db)

	app := &Application{Server: echo.New()}
	app.Server.HideBanner = true
	app.SetupRenderer()
	app.SetupRouting()
	app.Server.Use(middleware.Logger())

	app.Server.Logger.Fatal(app.Server.Start(":4000"))
}
