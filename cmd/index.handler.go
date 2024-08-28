package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/actanonv/excalidraw-local/ui"
	"github.com/donseba/go-htmx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (a *Application) Index(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	var userID string
	userID, ok := sess.Values["userID"].(string)

	if !ok {
		err = errors.New("error with userID")
		a.Server.Logger.Error(err)
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}

	user, err := services.Users().GetUserByID(userID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	drawingID := c.QueryParam("drawing")
	selectedCollection := c.QueryParam("select-collection")

	var activeDrawing services.Drawing
	var activeCollection services.Collection

	if drawingID == "" {
		activeDrawing = services.Drawing{}
	} else {
		activeDrawing, err = services.Drawings().GetDrawing(drawingID)

		if err != nil {
			a.Server.Logger.Error(err)

			if err.Error() == "drawing not found" {
				return c.String(http.StatusNotFound, "404 not found")
			}

			return err
		}
	}

	if selectedCollection != "" {
		activeCollection, err = services.Collections().GetCollection(selectedCollection)
		if err != nil && err.Error() == "collection not found" {
			return c.String(http.StatusNotFound, "Collection not found")
		} else if err != nil {
			a.Server.Logger.Error(err)
			return err
		}
	} else if drawingID != "" {
		activeCollection, err = services.Collections().GetCollection(activeDrawing.Collection)

		if err != nil && err.Error() == "collection not found" {
			return c.String(http.StatusNotFound, "Collection not found")
		} else if err != nil {
			a.Server.Logger.Error(err)
			return err
		}
	} else {
		activeCollection, err = services.Collections().GetCollection("default")

		if err != nil {
			a.Server.Logger.Error(err)
			return err
		}
	}

	appCollections, err := services.Collections().GetCollections()

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	drawingList, err := services.Drawings().GetDrawings(activeCollection.ID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	pageData := ui.IndexPageData{
		ActiveDrawing:   activeDrawing,
		DrawingListData: ui.DrawingListData{DrawingList: drawingList, PresenceMap: nil},
		CollectionsData: ui.IndexCollections{
			ActiveCollection: activeCollection,
			CollectionsList:  appCollections,
		},
		Editing: false,
	}

	if htmx.IsHxRequest(c.Request()) {
		if selectedCollection != "" {
			c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/app?select-collection=%s", activeCollection.ID))
		} else {
			c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/app?drawing=%s", activeDrawing.ID))
		}

		c.Response().Header().Add("HX-Trigger-After-Swap", "initExcalidraw")
	}

	if a.Presence.IsUserPresent(userID) {
		err = a.Presence.UserAtDrawing(userID, drawingID)
	} else {
		d := PresenceDetails{
			UserID:     userID,
			Name:       fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			login:      time.Now(),
			lastUpdate: time.Now(),
		}

		if drawingID != "" {
			d.LastDrawing = drawingID
		}

		a.Presence.AddUser(&d)
	}
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	presenceMap, err := makePresenceMap(drawingList, a.Presence.Users, userID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	pageData.DrawingListData.PresenceMap = presenceMap

	editRequest := c.QueryParam("edit")
	if editRequest != "" {
		ownerID, locked := a.Lock.IsDrawingLocked(drawingID)
		if !locked {
			pageData.Editing = true
			a.Lock.LockDrawing(userID, drawingID)
		} else if ownerID != userID {
			pageData.Editing = false
			pageData.Toast = "Sorry, someone else is editing this drawing."
			c.Response().Header().Add("HX-Trigger-After-Settle", "showToast")
		} else {
			pageData.Editing = true
		}
	}

	err = c.Render(http.StatusOK, "home", pageData)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return nil
}
