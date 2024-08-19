package services

import (
	"database/sql"
	"fmt"
	"time"

	"math/rand"
)

var serviceLayer AppServices

func Init(db *sql.DB) {

	serviceLayer = AppServices{
		Users:       &AppUsers{DB: db},
		Drawings:    &AppDrawings{DB: db},
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

func Drawings() DrawingSvc {
	return serviceLayer.Drawings
}

func Collections() CollectionsSvc {
	return serviceLayer.Collections
}

func Users() UserSvc {
	return serviceLayer.Users
}
