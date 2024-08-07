package ui

import "github.com/actanonv/excalidraw-local/services"

type IndexCollections struct {
	CollectionsList  []services.Collection
	ActiveCollection services.Collection
}

type IndexPageData struct {
	ActiveScene     services.Scene
	SceneList       []services.Scene
	CollectionsData IndexCollections
}
