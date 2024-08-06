package services

import (
	"database/sql"
	"errors"
	"log"
)

func (a *AppCollections) createDefaultCollection() (Collection, error) {
	id := "default"
	name := "default"

	query := "INSERT INTO Collections (Name, ID) VALUES (?, ?)"
	_, err := a.DB.Exec(query, name, id)
	if err != nil {
		log.Printf("Error creating collection: %v", err)
		return Collection{}, err
	}

	return Collection{ID: id, Name: name}, nil

}

func (a *AppCollections) CreateCollection(name string) (Collection, error) {
	id := generateID()

	query := "INSERT INTO Collections (Name, ID) VALUES (?, ?)"
	_, err := a.DB.Exec(query, name, id)
	if err != nil {
		log.Printf("Error creating collection: %v", err)
		return Collection{}, err
	}
	return Collection{ID: id, Name: name}, nil
}

func (a *AppCollections) GetCollection(collectionID string) (Collection, error) {
	query := "SELECT Name FROM Collections WHERE ID = ?"
	row := a.DB.QueryRow(query, collectionID)

	var name string
	if err := row.Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return Collection{}, errors.New("collection not found")
		}
		log.Printf("Error getting collection name: %v", err)
		return Collection{}, err
	}

	return Collection{Name: name, ID: collectionID}, nil
}

func (a *AppCollections) GetCollections() ([]Collection, error) {
	retv := make([]Collection, 0)
	query := "SELECT ID FROM Collections;"

	rows, err := a.DB.Query(query)
	if err != nil {
		log.Printf("Error getting Collections: %v", err)
		return nil, err
	}
	defer rows.Close()

	var CollectionIDs []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Printf("Error scanning Collection ID: %v", err)
			return nil, err
		}
		CollectionIDs = append(CollectionIDs, id)
	}

	for _, id := range CollectionIDs {
		Collection, err := a.GetCollection(id)
		if err != nil {
			return nil, nil
		}

		retv = append(retv, Collection)
	}

	return retv, nil
}
