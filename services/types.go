package services

import "database/sql"

type DrawingSvc interface {
	CreateDrawing(collectionID, name string) (Drawing, error)
	GetDrawings(collectionID string) ([]Drawing, error)
	GetDrawing(drawingID string) (Drawing, error)
	UpdateDrawingData(drawingID, data string) error
	DeleteDrawing(drawingID string) error
}

type CollectionsSvc interface {
	createDefaultCollection() (Collection, error)
	CreateCollection(name string) (Collection, error)
	GetCollection(collectionID string) (Collection, error)
	GetCollections() ([]Collection, error)
}

type AppDrawings struct {
	DB *sql.DB
}

type AppCollections struct {
	DB *sql.DB
}

type Collection struct {
	Name string
	ID   string
}

type Drawing struct {
	Name       string
	Collection string
	ID         string
	Data       string
}
