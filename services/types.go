package services

import "database/sql"

type SceneSvc interface {
	CreateScene(collectionID, name string) (Scene, error)
	GetScenes(collectionID string) ([]Scene, error)
	GetScene(sceneID string) (Scene, error)
	UpdateSceneData(sceneID, data string) error
	DeleteScene(sceneID string) error
}

type CollectionsSvc interface {
	createDefaultCollection() (Collection, error)
	CreateCollection(name string) (Collection, error)
	GetCollection(collectionID string) (Collection, error)
	GetCollections() ([]Collection, error)
}

type AppScenes struct {
	DB *sql.DB
}

type AppCollections struct {
	DB *sql.DB
}

type Collection struct {
	Name string
	ID   string
}

type Scene struct {
	Name       string
	Collection string
	ID         string
	Data       string
}
