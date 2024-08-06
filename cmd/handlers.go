package main

import (
	"net/http"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/actanonv/excalidraw-local/ui"
	"github.com/labstack/echo/v4"
)

func (a *Application) Index(c echo.Context) error {
	sceneID := c.Param("sceneID")

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
		ActiveScene:      activeScene,
		ActiveCollection: activeCollection,
		Collections:      appCollections,
		SceneList:        sceneList,
	}

	return c.Render(http.StatusOK, "home", pageData)
}
