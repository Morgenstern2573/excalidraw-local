package main

import (
	"fmt"
	"net/http"

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
		return c.Redirect(http.StatusMovedPermanently, "/login")
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
		ActiveDrawing: activeDrawing,
		DrawingList:   drawingList,
		CollectionsData: ui.IndexCollections{
			ActiveCollection: activeCollection,
			CollectionsList:  appCollections,
		},
	}

	if htmx.IsHxRequest(c.Request()) {
		if selectedCollection != "" {
			c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/?select-collection=%s", activeCollection.ID))
		} else {
			c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/?drawing=%s", activeDrawing.ID))
		}

		c.Response().Header().Add("HX-Trigger-After-Swap", "initExcalidraw")
	}

	err = a.Presence.UserAtDrawing(userID, drawingID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return c.Render(http.StatusOK, "home", pageData)
}
