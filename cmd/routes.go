package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (a *Application) SetupRouting() {
	routeProtector := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				a.Server.Logger.Error(err)
				return c.Redirect(http.StatusMovedPermanently, "/login")
			}

			log.Println(sess.Values["userEmail"])
			if sess.Values["userEmail"] == nil {
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

	appRoutes := a.Server.Group("/app")
	appRoutes.Use(routeProtector)
	appRoutes.GET("", a.Index)
	appRoutes.POST("/new-drawing", a.NewDrawing)
	appRoutes.POST("update-drawing-data", a.UpdateDrawingData)
	appRoutes.POST("/new-collection", a.NewCollection)
	appRoutes.GET("/drawing-list", a.DrawingList)
	appRoutes.DELETE("/", a.DeleteDrawing)
	//TODO: replace with fig config
	a.Server.Static("/public", "./public")
}
