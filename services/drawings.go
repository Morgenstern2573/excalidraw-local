package services

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (a *AppDrawings) CreateDrawing(collectionID, name string) (Drawing, error) {
	var err error
	if name == "" {
		return Drawing{}, errors.New("no name passed")
	}

	if collectionID == "" {
		return Drawing{}, errors.New("no collection passed")
	}

	id := generateID()

	drawing := Drawing{ID: id, Name: name, CollectionID: collectionID, Data: ""}
	a.DB.Create(&drawing)

	if err = a.DB.Error; err != nil {
		log.Printf("Error creating drawing: %v", err)
		return Drawing{}, err
	}
	return drawing, nil
}

func (a *AppDrawings) GetDrawing(drawingID string) (Drawing, error) {
	if drawingID == "" {
		return Drawing{}, errors.New("no id passed")
	}

	var drawing Drawing
	if err := a.DB.First(&drawing, "id = ?", drawingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Drawing{}, errors.New("drawing not found")
		}
		return Drawing{}, err
	}
	return drawing, nil
}

func (a *AppDrawings) GetDrawings(collectionID string) ([]Drawing, error) {
	var err error
	if collectionID == "" {
		return nil, errors.New("no collection passed")
	}

	var drawings []Drawing
	if err = a.DB.Where("collectionID = ?", collectionID).Find(&drawings).Error; err != nil {
		return nil, err
	}
	return drawings, nil
}

func (a *AppDrawings) GetDrawingData(drawingID string) (string, error) {
	if drawingID == "" {
		return "", errors.New("no id passed")
	}

	var drawing Drawing
	if err := a.DB.Select("Data").First(&drawing, "id = ?", drawingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("drawing not found")
		}
		return "", err
	}
	return drawing.Data, nil
}

func (a *AppDrawings) UpdateDrawingData(drawingID, data string) error {
	if drawingID == "" {
		return errors.New("no id passed")
	}

	tx := a.DB.Model(&Drawing{}).Where("id = ?", drawingID).Update("Data", data)

	if err := tx.Error; err != nil {
		return err
	}

	if tx.RowsAffected == 0 {
		return errors.New("id matches no records")
	}

	return nil
}

func (a *AppDrawings) GetDrawingName(drawingID string) (string, error) {
	var err error
	if drawingID == "" {
		return "", errors.New("no id passed")
	}

	var drawing Drawing
	if err = a.DB.Select("Name").First(&drawing, "id = ?", drawingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("drawing not found")
		}
		return "", err
	}

	return drawing.Name, nil
}

func (a *AppDrawings) GetParentCollectionID(drawingID string) (string, error) {
	if drawingID == "" {
		return "", errors.New("no id passed")
	}

	var drawing Drawing
	if err := a.DB.Select("Collection").First(&drawing, "id = ?", drawingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("drawing not found")
		}
		return "", err
	}
	return drawing.CollectionID, nil
}

func (a *AppDrawings) DeleteDrawing(drawingID string) error {
	if drawingID == "" {
		return errors.New("no drawing passed")
	}

	if err := a.DB.Delete(&Drawing{}, "id = ?", drawingID).Error; err != nil {
		return err
	}
	return nil
}
