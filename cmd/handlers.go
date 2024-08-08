package main

import (
	"fmt"
	"net/http"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/actanonv/excalidraw-local/ui"
	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
)

func (a *Application) Index(c echo.Context) error {
	drawingID := c.QueryParam("drawing")
	selectedCollection := c.QueryParam("select-collection")

	var activeDrawing services.Drawing
	var activeCollection services.Collection
	var err error

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

	return c.Render(http.StatusOK, "home", pageData)
}

func (a *Application) NewDrawing(c echo.Context) error {
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

	return c.Render(http.StatusOK, "home/drawing-list", updatedDrawingList)
}

func (a *Application) UpdateDrawingData(c echo.Context) error {
	type DrawingData struct {
		ID   string `form:"drawing"`
		Data string `form:"payload"`
	}

	scData := DrawingData{}
	err := c.Bind(&scData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	err = services.Drawings().UpdateDrawingData(scData.ID, scData.Data)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return nil
}

func (a *Application) NewCollection(c echo.Context) error {
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

func (a *Application) DrawingList(c echo.Context) error {
	collectionID := c.QueryParam("collection-id")
	if collectionID == "" {
		return c.String(http.StatusBadRequest, "collection id not found")
	}

	drawingList, err := services.Drawings().GetDrawings(collectionID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return c.Render(http.StatusOK, "home/drawing-list", drawingList)
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
