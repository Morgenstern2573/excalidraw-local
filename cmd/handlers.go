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
	sceneID := c.QueryParam("scene")

	var activeScene services.Scene
	var activeCollection services.Collection
	var err error

	if sceneID == "" {
		activeScene = services.Scene{}
		activeCollection, err = services.Collections().GetCollection("default")

		if err != nil {
			a.Server.Logger.Error(err)
			return err
		}
	} else {
		activeScene, err = services.Scenes().GetScene(sceneID)

		if err != nil {
			a.Server.Logger.Error(err)

			if err.Error() == "scene not found" {
				return c.String(http.StatusNotFound, "404 not found")
			}

			return err
		}

		activeCollection, err = services.Collections().GetCollection(activeScene.Collection)

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

	sceneList, err := services.Scenes().GetScenes(activeCollection.ID)
	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	pageData := ui.IndexPageData{
		ActiveScene: activeScene,
		SceneList:   sceneList,
		CollectionsData: ui.IndexCollections{
			ActiveCollection: activeCollection,
			CollectionsList:  appCollections,
		},
	}

	if htmx.IsHxRequest(c.Request()) {
		c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/?scene=%s", activeScene.ID))
		c.Response().Header().Add("HX-Trigger-After-Swap", "initExcalidraw")
	}

	return c.Render(http.StatusOK, "home", pageData)
}

func (a *Application) NewScene(c echo.Context) error {
	type FormData struct {
		CollectionID string `form:"collection-ID"`
		SceneName    string `form:"scene-name"`
	}

	fData := new(FormData)
	err := c.Bind(fData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	scene, err := services.Scenes().CreateScene(
		fData.CollectionID, fData.SceneName)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	updatedSceneList, err := services.Scenes().GetScenes(scene.Collection)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	return c.Render(http.StatusOK, "home/scene-list", updatedSceneList)
}

func (a *Application) UpdateSceneData(c echo.Context) error {
	type SceneData struct {
		ID   string `form:"scene"`
		Data string `form:"payload"`
	}

	scData := SceneData{}
	err := c.Bind(&scData)

	if err != nil {
		a.Server.Logger.Error(err)
		return err
	}

	err = services.Scenes().UpdateSceneData(scData.ID, scData.Data)

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

func (a *Application) SceneList(c echo.Context) error {
	return nil
}
