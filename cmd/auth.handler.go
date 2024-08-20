package main

import (
	"net/http"
	"time"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/actanonv/excalidraw-local/ui"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (a *Application) RenderLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func (a *Application) RenderRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register", nil)
}

func (a *Application) RegisterUser(c echo.Context) error {
	type FormData struct {
		FirstName         string `form:"first-name"`
		LastName          string `form:"last-name"`
		Email             string `form:"email"`
		Password          string `form:"password"`
		ConfirmedPassword string `form:"confirm-password"`
	}

	var formData FormData

	err := c.Bind(&formData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	if formData.Email == "" ||
		formData.Password == "" ||
		formData.ConfirmedPassword == "" {
		return c.String(http.StatusBadRequest, "required param not found")
	}

	pageData := ui.AuthPageData{}

	if formData.Password != formData.ConfirmedPassword {
		pageData.Error = "Password and confirm password don't match"
		return c.Render(http.StatusOK, "register", pageData)
	}

	user, err := services.Users().CreateUser(formData.FirstName, formData.LastName, formData.Email, formData.Password)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	d := PresenceDetails{
		UserID:     user.Email,
		Name:       user.Email,
		login:      time.Now(),
		lastUpdate: time.Now(),
	}
	a.Presence.AddUser(&d)

	return c.Redirect(http.StatusFound, "/app")
}

func (a *Application) LoginUser(c echo.Context) error {
	type FormData struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var formData FormData

	err := c.Bind(&formData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	if formData.Email == "" ||
		formData.Password == "" {
		return c.String(http.StatusBadRequest, "required param not found")
	}

	user, err := services.Users().GetUser(formData.Email)

	pageData := ui.AuthPageData{}

	if err != nil {
		pageData.Error = "incorrect username or password"
		return c.Render(http.StatusOK, "login", pageData)
	}

	if formData.Password != user.PasswordHash {
		pageData.Error = "incorrect username or password"
		return c.Render(http.StatusOK, "login", pageData)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	sess.Values["userID"] = user.ID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	d := PresenceDetails{
		UserID:     user.ID,
		Name:       user.Email,
		login:      time.Now(),
		lastUpdate: time.Now(),
	}
	a.Presence.AddUser(&d)
	return c.Redirect(http.StatusFound, "/app")
}
