package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_createDefaultCollection(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Collection{})

	a := AppCollections{DB: db}

	tests := []struct {
		testName string
		wantErr  error
		expect   Collection
	}{
		{
			testName: "create default collection successfully",
			wantErr:  nil,
			expect:   Collection{ID: "default", Name: "default"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			collection, err := a.createDefaultCollection()

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, but got %v", tt.wantErr, err)
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assert.Equal(t, tt.expect.ID, collection.ID)
			assert.Equal(t, tt.expect.Name, collection.Name)
		})
	}

	if err = db.Exec("DROP TABLE collections;").Error; err != nil {
		t.Fatal(err)
	}
}

func Test_CreateCollection(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Collection{})

	a := AppCollections{DB: db}

	tests := []struct {
		testName string
		name     string
		wantErr  error
		expect   Collection
	}{
		{
			testName: "valid collection creation",
			name:     "MyCollection",
			wantErr:  nil,
			expect:   Collection{Name: "MyCollection"},
		},
		{
			testName: "empty name",
			name:     "",
			wantErr:  errors.New("no name passed"),
			expect:   Collection{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			collection, err := a.CreateCollection(tt.name)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, but got %v", tt.wantErr, err)
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assert.Equal(t, tt.expect.Name, collection.Name)
			assert.NotEmpty(t, collection.ID)
		})
	}

	if err = db.Exec("DROP TABLE collections;").Error; err != nil {
		t.Fatal(err)
	}
}

func Test_GetCollection(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Collection{})

	a := AppCollections{DB: db}

	// Prepare a collection for testing
	expectedCollection := Collection{ID: "testID", Name: "Test Collection"}
	db.Create(&expectedCollection)

	tests := []struct {
		testName     string
		collectionID string
		wantErr      error
		expect       Collection
	}{
		{
			testName:     "valid ID",
			collectionID: "testID",
			wantErr:      nil,
			expect:       expectedCollection,
		},
		{
			testName:     "empty ID",
			collectionID: "",
			wantErr:      errors.New("no id passed"),
			expect:       Collection{},
		},
		{
			testName:     "collection not found",
			collectionID: "nonexistentID",
			wantErr:      errors.New("collection not found"),
			expect:       Collection{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			collection, err := a.GetCollection(tt.collectionID)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, but got %v", tt.wantErr, err)
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assert.Equal(t, tt.expect.ID, collection.ID)
			assert.Equal(t, tt.expect.Name, collection.Name)
		})
	}

	if err = db.Exec("DROP TABLE collections;").Error; err != nil {
		t.Fatal(err)
	}
}

func Test_GetCollections(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Collection{})

	a := AppCollections{DB: db}

	collection1 := Collection{ID: "testID1", Name: "Collection 1"}
	collection2 := Collection{ID: "testID2", Name: "Collection 2"}
	db.Create(&[]Collection{collection1, collection2})

	tests := []struct {
		testName string
		wantErr  error
		expect   []Collection
	}{
		{
			testName: "get all collections",
			wantErr:  nil,
			expect:   []Collection{collection1, collection2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			collections, err := a.GetCollections()

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, but got %v", tt.wantErr, err)
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assert.Equal(t, len(tt.expect), len(collections))
			for i := range collections {
				assert.Equal(t, tt.expect[i].ID, collections[i].ID)
				assert.Equal(t, tt.expect[i].Name, collections[i].Name)
			}
		})
	}

	if err = db.Exec("DROP TABLE collections;").Error; err != nil {
		t.Fatal(err)
	}
}
