package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (a *Application) SetupRouting() {
	routeProtector := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				err  error
				sess *sessions.Session
			)

			if sess, err = session.Get("session", c); err != nil {
				a.Server.Logger.Error(err)
				return c.Redirect(http.StatusMovedPermanently, "/login")
			}

			if sess.Values["userID"] == nil {
				a.Server.Logger.Info("no user id found in session")
				return c.Redirect(http.StatusMovedPermanently, "/login")
			}

			err = next(c)
			return err
		}
	}

	a.Server.GET("/login", a.RenderLogin)
	a.Server.POST("/login", a.LoginUser)

	a.Server.GET("/register", a.RenderRegister)
	a.Server.POST("/register", a.RegisterUser)

	appRoutes := a.Server.Group("/app", routeProtector)
	appRoutes.GET("", a.Index)
	appRoutes.POST("/new-drawing", a.CreateDrawing)
	appRoutes.POST("update-drawing-data", a.UpdateDrawingData)
	appRoutes.POST("/create-collection", a.CreateCollection)
	appRoutes.GET("/presence-update", a.UpdatePresenceIndicators)
	appRoutes.DELETE("", a.DeleteDrawing)
	//TODO: replace with fig config
	a.Server.Static("/public", "./public")
}
