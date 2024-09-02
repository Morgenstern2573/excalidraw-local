package services

import (
	"time"

	"gorm.io/gorm"
)

type AppServices struct {
	Drawings    DrawingSvc
	Collections CollectionsSvc
	Users       UserSvc
	// DrawingAccessLogs DrawingAccessLogSvc
}

type DrawingAccessLogSvc interface {
	RecordAccess(userID string, drawingID string) error
	GetUserLogs(userID string, count int) ([]AccessLog, error)
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
	DB *gorm.DB
}

type AppCollections struct {
	DB *gorm.DB
}

type AppUsers struct {
	DB *gorm.DB
}

type AppDrawingAccessLogs struct {
	DB *gorm.DB
}

type AccessLog struct {
	ID         string `gorm:"primaryKey"`
	UserID     string
	DrawingID  string
	AccessedAt time.Time
}

type User struct {
	ID           string `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
}

type Collection struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

type Drawing struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	CollectionID string
	Data         string
}
