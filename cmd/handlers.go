package main

import (
	"fmt"
	"net/http"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/actanonv/excalidraw-local/ui"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (a *Application) CreateDrawing(c echo.Context) error {
	type FormData struct {
		CollectionID string `form:"collection-ID"`
		DrawingName  string `form:"drawing-name"`
	}

	fData := new(FormData)
	err := c.Bind(fData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	drawing, err := services.Drawings().CreateDrawing(
		fData.CollectionID, fData.DrawingName)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	updatedDrawingList, err := services.Drawings().GetDrawings(drawing.Collection)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return c.Render(http.StatusOK, "home/drawing-list", ui.DrawingListData{DrawingList: updatedDrawingList})
}

func (a *Application) UpdateDrawingData(c echo.Context) error {
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

	type DrawingData struct {
		ID   string `form:"drawing"`
		Data string `form:"payload"`
	}

	drawingData := DrawingData{}
	err = c.Bind(&drawingData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	err = services.Drawings().UpdateDrawingData(drawingData.ID, drawingData.Data)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	err = a.Presence.UserAtDrawing(userID, drawingData.ID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return nil
}

func (a *Application) CreateCollection(c echo.Context) error {
	type FormData struct {
		CollectionName string `form:"collection-name"`
	}

	fData := new(FormData)
	err := c.Bind(fData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	newCollection, err := services.Collections().CreateCollection(
		fData.CollectionName)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	updatedCollections, err := services.Collections().GetCollections()

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return c.Render(http.StatusOK, "home/select-collection", ui.IndexCollections{
		CollectionsList:  updatedCollections,
		ActiveCollection: newCollection,
	})
}

func (a *Application) UpdatePresenceIndicators(c echo.Context) error {
	collectionID := c.QueryParam("collection-id")
	if collectionID == "" {
		return c.String(http.StatusBadRequest, "collection id not found")
	}

	drawingList, err := services.Drawings().GetDrawings(collectionID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	presenceMap := make(map[string][]string)

	for _, details := range a.Presence.Users {
		for _, drawing := range drawingList {
			if details.LastDrawing == drawing.ID {
				_, found := presenceMap[details.LastDrawing]
				if found {
					presenceMap[details.LastDrawing] = append(presenceMap[details.LastDrawing], details.UserID)
				} else {
					presenceMap[details.LastDrawing] = []string{details.UserID}
				}

				break
			}
		}
	}

	return c.Render(http.StatusOK, "home/drawing-list", ui.DrawingListData{DrawingList: drawingList, PresenceMap: presenceMap})
}

func (a *Application) DeleteDrawing(c echo.Context) error {
	drawingID := c.QueryParam("drawing-ID")

	if drawingID == "" {
		return c.String(http.StatusBadRequest, "no drawing id found")
	}

	err := services.Drawings().DeleteDrawing(drawingID)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	c.Response().Header().Add("HX-Trigger", fmt.Sprintf(`{"deleteDrawing":%q}`, drawingID))

	return nil
}
