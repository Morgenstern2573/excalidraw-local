package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Test_CreateDrawing(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}

	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	tests := []struct {
		testName     string
		collectionID string
		drawingName  string
		wantErr      error
		expect       Drawing
	}{
		{
			testName:     "default",
			collectionID: "test",
			drawingName:  "testDrawing",
			wantErr:      nil,
			expect:       Drawing{Name: "testDrawing", CollectionID: "test"},
		},
		{
			testName:     "empty drawing name",
			collectionID: "test",
			drawingName:  "",
			wantErr:      errors.New("no name passed"),
			expect:       Drawing{},
		},
		{
			testName:     "empty collection ID",
			collectionID: "",
			drawingName:  "testDrawing",
			wantErr:      errors.New("no collection passed"),
			expect:       Drawing{},
		},
		{
			testName:     "empty collection ID and name",
			collectionID: "",
			drawingName:  "",
			wantErr:      errors.New("no name passed"),
			expect:       Drawing{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			drawing, err := a.CreateDrawing(tt.collectionID, tt.drawingName)

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

			assert.Equal(t, tt.drawingName, drawing.Name)
			assert.Equal(t, tt.collectionID, drawing.CollectionID)
			assert.Empty(t, drawing.Data)
			assert.NotEmpty(t, drawing.ID)
		})
	}

	if err = db.Exec("DROP TABLE drawings;").Error; err != nil {
		t.Fatal(err)
	}
}

func Test_GetDrawing(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	// Prepare a drawing for testing
	expectedDrawing := Drawing{ID: "testID", Name: "testName", CollectionID: "testCollection"}
	db.Create(&expectedDrawing)

	tests := []struct {
		testName  string
		drawingID string
		wantErr   error
		expect    Drawing
	}{
		{
			testName:  "valid drawing",
			drawingID: "testID",
			wantErr:   nil,
			expect:    expectedDrawing,
		},
		{
			testName:  "empty drawing ID",
			drawingID: "",
			wantErr:   errors.New("no id passed"),
			expect:    Drawing{},
		},
		{
			testName:  "drawing not found",
			drawingID: "nonExistentID",
			wantErr:   errors.New("drawing not found"),
			expect:    Drawing{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			drawing, err := a.GetDrawing(tt.drawingID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			assert.Equal(t, tt.expect.ID, drawing.ID)
			assert.Equal(t, tt.expect.Name, drawing.Name)
			assert.Equal(t, tt.expect.CollectionID, drawing.CollectionID)
		})
	}

	db.Exec("DROP TABLE drawings;")
}

func Test_GetDrawings(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	// Prepare drawings for testing
	drawings := []Drawing{
		{ID: "testID1", Name: "testName1", CollectionID: "testCollection"},
		{ID: "testID2", Name: "testName2", CollectionID: "testCollection"},
	}
	db.Create(&drawings)

	tests := []struct {
		testName     string
		collectionID string
		wantErr      error
		expect       []Drawing
	}{
		{
			testName:     "valid collection",
			collectionID: "testCollection",
			wantErr:      nil,
			expect:       drawings,
		},
		{
			testName:     "empty collection ID",
			collectionID: "",
			wantErr:      errors.New("no collection passed"),
			expect:       nil,
		},
		{
			testName:     "collection not found",
			collectionID: "nonExistentCollection",
			wantErr:      nil, // no error, just no results
			expect:       []Drawing{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			drawings, err := a.GetDrawings(tt.collectionID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			assert.Equal(t, len(tt.expect), len(drawings))
		})
	}

	db.Exec("DROP TABLE drawings;")
}

func Test_GetDrawingData(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	// Prepare a drawing for testing
	expectedDrawing := Drawing{ID: "testID", Data: "testData"}
	db.Create(&expectedDrawing)

	tests := []struct {
		testName  string
		drawingID string
		wantErr   error
		expect    string
	}{
		{
			testName:  "valid drawing ID",
			drawingID: "testID",
			wantErr:   nil,
			expect:    "testData",
		},
		{
			testName:  "empty drawing ID",
			drawingID: "",
			wantErr:   errors.New("no id passed"),
			expect:    "",
		},
		{
			testName:  "drawing not found",
			drawingID: "nonExistentID",
			wantErr:   errors.New("drawing not found"),
			expect:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			data, err := a.GetDrawingData(tt.drawingID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			assert.Equal(t, tt.expect, data)
		})
	}

	db.Exec("DROP TABLE drawings;")
}

func Test_UpdateDrawingData(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	// Prepare a drawing for testing
	expectedDrawing := Drawing{ID: "testID", Data: "oldData"}
	db.Create(&expectedDrawing)

	tests := []struct {
		testName  string
		drawingID string
		data      string
		wantErr   error
	}{
		{
			testName:  "valid update",
			drawingID: "testID",
			data:      "newData",
			wantErr:   nil,
		},
		{
			testName:  "empty drawing ID",
			drawingID: "",
			data:      "newData",
			wantErr:   errors.New("no id passed"),
		},
		{
			testName:  "drawing not found",
			drawingID: "nonExistentID",
			data:      "newData",
			wantErr:   errors.New("id matches no records"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := a.UpdateDrawingData(tt.drawingID, tt.data)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			var updatedDrawing Drawing
			db.First(&updatedDrawing, "id = ?", tt.drawingID)
			assert.Equal(t, tt.data, updatedDrawing.Data)
		})
	}

	db.Exec("DROP TABLE drawings;")
}

func Test_GetDrawingName(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	// Prepare a drawing for testing
	expectedDrawing := Drawing{ID: "testID", Name: "testName"}
	db.Create(&expectedDrawing)

	tests := []struct {
		testName  string
		drawingID string
		wantErr   error
		expect    string
	}{
		{
			testName:  "valid drawing ID",
			drawingID: "testID",
			wantErr:   nil,
			expect:    "testName",
		},
		{
			testName:  "empty drawing ID",
			drawingID: "",
			wantErr:   errors.New("no id passed"),
			expect:    "",
		},
		{
			testName:  "drawing not found",
			drawingID: "nonExistentID",
			wantErr:   errors.New("drawing not found"),
			expect:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			name, err := a.GetDrawingName(tt.drawingID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			assert.Equal(t, tt.expect, name)
		})
	}

	db.Exec("DROP TABLE drawings;")
}
func Test_GetParentCollectionID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	// Prepare a drawing for testing
	expectedDrawing := Drawing{ID: "testID", CollectionID: "testCollection"}
	db.Create(&expectedDrawing)

	tests := []struct {
		testName  string
		drawingID string
		wantErr   error
		expect    string
	}{
		{
			testName:  "valid drawing ID",
			drawingID: "testID",
			wantErr:   nil,
			expect:    "testCollection",
		},
		{
			testName:  "empty drawing ID",
			drawingID: "",
			wantErr:   errors.New("no id passed"),
			expect:    "",
		},
		{
			testName:  "drawing not found",
			drawingID: "nonExistentID",
			wantErr:   errors.New("drawing not found"),
			expect:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			collectionID, err := a.GetParentCollectionID(tt.drawingID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			assert.Equal(t, tt.expect, collectionID)
		})
	}

	db.Exec("DROP TABLE drawings;")
}

func Test_DeleteDrawing(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	a := AppDrawings{DB: db}

	// Prepare a drawing for testing
	drawing := Drawing{ID: "testID"}
	db.Create(&drawing)

	tests := []struct {
		testName  string
		drawingID string
		wantErr   error
	}{
		{
			testName:  "valid delete",
			drawingID: "testID",
			wantErr:   nil,
		},
		{
			testName:  "empty drawing ID",
			drawingID: "",
			wantErr:   errors.New("no drawing passed"),
		},
		{
			testName:  "drawing not found",
			drawingID: "nonExistentID",
			wantErr:   nil, // gorm does not return an error if no rows affected
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := a.DeleteDrawing(tt.drawingID)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			var deletedDrawing Drawing
			result := db.First(&deletedDrawing, "id = ?", tt.drawingID)
			assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
		})
	}

	db.Exec("DROP TABLE drawings;")
}

/*
	TEST SETUP CODE
	db, err := gorm.Open(sqlite.Open("../testData/xdraw.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}

	defer sqlDB.Close()

	db.AutoMigrate(&Drawing{})

	if err = db.Exec("DROP TABLE drawings").Error; err != nil {
		t.Fatal(err)
	}
*/
