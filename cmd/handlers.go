package main

import (
	"fmt"
	"net/http"

	"github.com/actanonv/excalidraw-local/services"
	"github.com/actanonv/excalidraw-local/ui"
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
		ActiveScene:      activeScene,
		ActiveCollection: activeCollection,
		Collections:      appCollections,
		SceneList:        sceneList,
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

	return c.Redirect(http.StatusFound, fmt.Sprintf("/?scene=%s", scene.ID))
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
