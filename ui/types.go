package ui

import "github.com/actanonv/excalidraw-local/services"

type IndexCollections struct {
	CollectionsList  []services.Collection
	ActiveCollection services.Collection
}

type DrawingListData struct {
	DrawingList []services.Drawing
	PresenceMap map[string][]string
}

type IndexPageData struct {
	ActiveDrawing   services.Drawing
	CollectionsData IndexCollections
	DrawingListData DrawingListData
}

type AuthPageData struct {
	Error string
}
