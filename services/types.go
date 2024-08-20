package services

import (
	"database/sql"
)

type AppServices struct {
	Drawings    DrawingSvc
	Collections CollectionsSvc
	Users       UserSvc
}

type UserSvc interface {
	CreateUser(firstName, lastName, email, passwordHash string) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByID(email string) (User, error)
	DeleteUser(email string) error
	UpdateUser(column, data string) error
}

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

type AppUsers struct {
	DB *sql.DB
}

type User struct {
	FirstName    string
	LastName     string
	ID           string
	Email        string
	PasswordHash string
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
