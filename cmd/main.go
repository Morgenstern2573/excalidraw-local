package main

import (
	"log"
	"time"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type Application struct {
	Server   *echo.Echo
	Presence *Presence
	Lock     *DrawingLock
}

func main() {

	db := connectToDB()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}

		err = sqlDB.Close()
		if err != nil {
			log.Println("Failed to close database connection:", err)
		}
	}()

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

func connectToDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./data/xdraw-local.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Enable foreign key checks
	tx := db.Exec("PRAGMA foreign_keys = ON;")
	if err = tx.Error; err != nil {
		log.Fatal(err)
	}

	return db
}
