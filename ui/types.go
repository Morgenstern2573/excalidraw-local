package ui

import "github.com/actanonv/excalidraw-local/services"

type IndexPageData struct {
	ActiveScene      services.Scene
	ActiveCollection services.Collection
	Collections      []services.Collection
	SceneList        []services.Scene
}
