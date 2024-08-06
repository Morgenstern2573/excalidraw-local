package services

import (
	"database/sql"
	"fmt"
	"time"

	"math/rand"
)

type AppServices struct {
	Scenes      SceneSvc
	Collections CollectionsSvc
}

var ServiceLayer AppServices

func Init(db *sql.DB) {

	ServiceLayer = AppServices{
		Scenes:      &AppScenes{DB: db},
		Collections: &AppCollections{DB: db},
	}

	_, err := Collections().GetCollection("default")
	if err != nil && err.Error() == "collection not found" {
		_, err = Collections().createDefaultCollection()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
}

func generateID() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	retv := 1000000000 + r.Int63n(9000000000)

	return fmt.Sprintf("%d", retv)
}

func Scenes() SceneSvc {
	return ServiceLayer.Scenes
}

func Collections() CollectionsSvc {
	return ServiceLayer.Collections
}
