package ui

import "github.com/actanonv/excalidraw-local/services"

type IndexCollections struct {
	CollectionsList  []services.Collection
	ActiveCollection services.Collection
}

type PresentUser struct {
	Initials string
	Name     string
}

type DrawingListData struct {
	DrawingList []services.Drawing
	PresenceMap map[string][]PresentUser
}

type IndexPageData struct {
	ActiveDrawing   services.Drawing
	CollectionsData IndexCollections
	DrawingListData DrawingListData
	Editing         bool
	Toast           string
}

type AuthPageData struct {
	Error string
}
