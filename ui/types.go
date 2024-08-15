package ui

import "github.com/actanonv/excalidraw-local/services"

type IndexCollections struct {
	CollectionsList  []services.Collection
	ActiveCollection services.Collection
}

type IndexPageData struct {
	ActiveDrawing   services.Drawing
	DrawingList     []services.Drawing
	CollectionsData IndexCollections
}

type AuthPageData struct {
	Error string
}
