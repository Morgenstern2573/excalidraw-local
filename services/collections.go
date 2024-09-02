package services

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (a *AppCollections) createDefaultCollection() (Collection, error) {
	collection := Collection{ID: "default", Name: "default"}

	if err := a.DB.Create(&collection).Error; err != nil {
		log.Printf("Error creating collection: %v", err)
		return Collection{}, err
	}

	return collection, nil
}

func (a *AppCollections) CreateCollection(name string) (Collection, error) {
	if name == "" {
		return Collection{}, errors.New("no name passed")
	}

	id := generateID()

	collection := Collection{ID: id, Name: name}

	if err := a.DB.Create(&collection).Error; err != nil {
		log.Printf("Error creating collection: %v", err)
		return Collection{}, err
	}

	return collection, nil
}

func (a *AppCollections) GetCollection(collectionID string) (Collection, error) {
	if collectionID == "" {
		return Collection{}, errors.New("no id passed")
	}
	var collection Collection

	if err := a.DB.First(&collection, "id = ?", collectionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Collection{}, errors.New("collection not found")
		}

		return Collection{}, err
	}

	return collection, nil
}

func (a *AppCollections) GetCollections() ([]Collection, error) {
	var retv []Collection

	result := a.DB.Find(&retv)
	if err := result.Error; err != nil {
		return nil, err
	}

	return retv, nil
}
