package services

import (
	"database/sql"

	"github.com/rs/xid"
)

var serviceLayer AppServices

func Init(db *sql.DB) {

	serviceLayer = AppServices{
		Users:             &AppUsers{DB: db},
		Drawings:          &AppDrawings{DB: db},
		Collections:       &AppCollections{DB: db},
		DrawingAccessLogs: &AppDrawingAccessLogs{DB: db},
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
	guid := xid.New()

	return guid.String()
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

func AccessLogs() DrawingAccessLogSvc {
	return serviceLayer.DrawingAccessLogs
}
