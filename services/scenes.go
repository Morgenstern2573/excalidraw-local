package services

import (
	"database/sql"
	"errors"
	"log"
)

func (a *AppScenes) CreateScene(collectionID, name string) (Scene, error) {
	if name == "" {
		return Scene{}, errors.New("no name passed")
	}

	if collectionID == "" {
		return Scene{}, errors.New("no collection passed")
	}

	id := generateID()

	query := "INSERT INTO Scenes (Name, CollectionID, ID, Data) VALUES (?, ?, ?, ?)"
	_, err := a.DB.Exec(query, name, collectionID, id, "")
	if err != nil {
		log.Printf("Error creating scene: %v", err)
		return Scene{}, err
	}
	return Scene{ID: id}, nil
}

func (a *AppScenes) GetScene(sceneID string) (Scene, error) {
	if sceneID == "" {
		return Scene{}, errors.New("no id passed")
	}

	retv := Scene{}
	retv.ID = sceneID
	name, err := a.GetSceneName(sceneID)

	if err != nil {
		return Scene{}, err
	}

	data, err := a.GetSceneData(sceneID)
	if err != nil {
		return Scene{}, err
	}

	collection, err := a.GetParentCollectionID(sceneID)
	if err != nil {
		return Scene{}, err
	}

	retv.Name = name
	retv.Data = data
	retv.Collection = collection
	return retv, nil
}

func (a *AppScenes) GetScenes(collectionID string) ([]Scene, error) {
	if collectionID == "" {
		return nil, errors.New("no collection passed")
	}

	retv := make([]Scene, 0)
	query := "SELECT ID FROM Scenes WHERE CollectionID = ?"
	rows, err := a.DB.Query(query, collectionID)
	if err != nil {
		log.Printf("Error getting scenes: %v", err)
		return nil, err
	}
	defer rows.Close()

	var sceneIDs []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Printf("Error scanning scene ID: %v", err)
			return nil, err
		}
		sceneIDs = append(sceneIDs, id)
	}

	for _, id := range sceneIDs {
		scene, err := a.GetScene(id)
		if err != nil {
			return nil, nil
		}

		retv = append(retv, scene)
	}

	return retv, nil
}

func (a *AppScenes) GetSceneData(sceneID string) (string, error) {
	if sceneID == "" {
		return "", errors.New("no id passed")
	}

	query := "SELECT Data FROM Scenes WHERE ID = ?"
	row := a.DB.QueryRow(query, sceneID)

	var data string
	if err := row.Scan(&data); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("scene not found")
		}
		log.Printf("Error getting scene data: %v", err)
		return "", err
	}
	return data, nil
}

func (a *AppScenes) GetSceneName(sceneID string) (string, error) {
	if sceneID == "" {
		return "", errors.New("no id passed")
	}

	query := "SELECT Name FROM Scenes WHERE ID = ?"
	row := a.DB.QueryRow(query, sceneID)

	var name string
	if err := row.Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("scene not found")
		}
		log.Printf("Error getting scene name: %v", err)
		return "", err
	}
	return name, nil
}

func (a *AppScenes) GetParentCollectionID(sceneID string) (string, error) {
	if sceneID == "" {
		return "", errors.New("no id passed")
	}

	query := "SELECT CollectionID FROM Scenes WHERE ID = ?"
	row := a.DB.QueryRow(query, sceneID)

	var id string
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("scene not found")
		}
		log.Printf("Error getting parent collection id: %v", err)
		return "", err
	}
	return id, nil
}
