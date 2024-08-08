package services

import (
	"database/sql"
	"errors"
	"log"
)

func (a *AppDrawings) CreateDrawing(collectionID, name string) (Drawing, error) {
	if name == "" {
		return Drawing{}, errors.New("no name passed")
	}

	if collectionID == "" {
		return Drawing{}, errors.New("no collection passed")
	}

	id := generateID()

	query := "INSERT INTO Drawings (Name, CollectionID, ID, Data) VALUES (?, ?, ?, ?)"
	_, err := a.DB.Exec(query, name, collectionID, id, "")
	if err != nil {
		log.Printf("Error creating drawing: %v", err)
		return Drawing{}, err
	}
	return Drawing{ID: id, Collection: collectionID, Name: name}, nil
}

func (a *AppDrawings) GetDrawing(drawingID string) (Drawing, error) {
	if drawingID == "" {
		return Drawing{}, errors.New("no id passed")
	}

	retv := Drawing{}
	retv.ID = drawingID
	name, err := a.GetDrawingName(drawingID)

	if err != nil {
		return Drawing{}, err
	}

	data, err := a.GetDrawingData(drawingID)
	if err != nil {
		return Drawing{}, err
	}

	collection, err := a.GetParentCollectionID(drawingID)
	if err != nil {
		return Drawing{}, err
	}

	retv.Name = name
	retv.Data = data
	retv.Collection = collection
	return retv, nil
}

func (a *AppDrawings) GetDrawings(collectionID string) ([]Drawing, error) {
	if collectionID == "" {
		return nil, errors.New("no collection passed")
	}

	retv := make([]Drawing, 0)
	query := "SELECT ID FROM Drawings WHERE CollectionID = ?"
	rows, err := a.DB.Query(query, collectionID)
	if err != nil {
		log.Printf("Error getting drawings: %v", err)
		return nil, err
	}
	defer rows.Close()

	var drawingIDs []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Printf("Error scanning drawing ID: %v", err)
			return nil, err
		}
		drawingIDs = append(drawingIDs, id)
	}

	for _, id := range drawingIDs {
		drawing, err := a.GetDrawing(id)
		if err != nil {
			return nil, nil
		}

		retv = append(retv, drawing)
	}

	return retv, nil
}

func (a *AppDrawings) GetDrawingData(drawingID string) (string, error) {
	if drawingID == "" {
		return "", errors.New("no id passed")
	}

	query := "SELECT Data FROM Drawings WHERE ID = ?"
	row := a.DB.QueryRow(query, drawingID)

	var data string
	if err := row.Scan(&data); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("drawing not found")
		}
		log.Printf("Error getting drawing data: %v", err)
		return "", err
	}
	return data, nil
}

func (a *AppDrawings) UpdateDrawingData(drawingID, data string) error {
	if drawingID == "" {
		return errors.New("no id passed")
	}

	query := "UPDATE Drawings SET Data = ? WHERE ID = ?"
	_, err := a.DB.Exec(query, data, drawingID)

	if err != nil {
		return nil
	}

	return nil
}

func (a *AppDrawings) GetDrawingName(drawingID string) (string, error) {
	if drawingID == "" {
		return "", errors.New("no id passed")
	}

	query := "SELECT Name FROM Drawings WHERE ID = ?"
	row := a.DB.QueryRow(query, drawingID)

	var name string
	if err := row.Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("drawing not found")
		}
		log.Printf("Error getting drawing name: %v", err)
		return "", err
	}
	return name, nil
}

func (a *AppDrawings) GetParentCollectionID(drawingID string) (string, error) {
	if drawingID == "" {
		return "", errors.New("no id passed")
	}

	query := "SELECT CollectionID FROM Drawings WHERE ID = ?"
	row := a.DB.QueryRow(query, drawingID)

	var id string
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("drawing not found")
		}
		log.Printf("Error getting parent collection id: %v", err)
		return "", err
	}
	return id, nil
}

func (a *AppDrawings) DeleteDrawing(drawingID string) error {
	if drawingID == "" {
		return errors.New("no drawing passed")
	}

	query := "DELETE FROM Drawings WHERE ID = ?"
	_, err := a.DB.Exec(query, drawingID)
	if err != nil {
		log.Printf("Error deleting drawing: %v", err)
		return err
	}
	return nil
}
