package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	Server   *echo.Echo
	Presence *Presence
	Lock     *DrawingLock
}

func main() {

	db := connectToDB()
	defer db.Close()

	services.Init(db)

	app := &Application{Server: echo.New(), Presence: newPresence(), Lock: NewLock()}
	go func() {
		for {
			for userID, details := range app.Presence.Users {
				if time.Since(details.lastUpdate) > (30 * time.Minute) {
					app.Presence.RemoveUser(userID)
					app.Lock.RemoveUser(userID)
				}
			}

			time.Sleep(30 * time.Minute)
		}
	}()

	app.Server.HideBanner = true
	// replace with env variable
	app.Server.Use(session.Middleware(sessions.NewCookieStore([]byte("internal_secret_super_secret"))))

	app.SetupRenderer()
	app.SetupRouting()

	app.Server.Logger.Fatal(app.Server.Start(":4000"))
}

func connectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/xdraw-local.db")
	if err != nil {
		log.Fatal(err)
	}

	// Enable foreign key checks
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
