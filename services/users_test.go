package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Test_CreateUser(t *testing.T) {
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

	db.AutoMigrate(&User{})

	a := AppUsers{DB: db}

	tests := []struct {
		testName     string
		firstName    string
		lastName     string
		email        string
		passwordHash string
		wantErr      error
		expect       User
	}{
		{
			testName:     "valid input",
			firstName:    "John",
			lastName:     "Doe",
			email:        "john.doe@example.com",
			passwordHash: "hashedpassword",
			wantErr:      nil,
			expect:       User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", PasswordHash: "hashedpassword"},
		},
		{
			testName:     "missing email",
			firstName:    "John",
			lastName:     "Doe",
			email:        "",
			passwordHash: "hashedpassword",
			wantErr:      errors.New("no email passed"),
			expect:       User{},
		},
		{
			testName:     "missing password hash",
			firstName:    "John",
			lastName:     "Doe",
			email:        "john.doe@example.com",
			passwordHash: "",
			wantErr:      errors.New("no password hash passed"),
			expect:       User{},
		},
		{
			testName:     "missing both email and password hash",
			firstName:    "John",
			lastName:     "Doe",
			email:        "",
			passwordHash: "",
			wantErr:      errors.New("no email passed"),
			expect:       User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user, err := a.CreateUser(tt.firstName, tt.lastName, tt.email, tt.passwordHash)

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

			// Assert that the returned user matches the expected user
			assert.Equal(t, tt.expect.FirstName, user.FirstName)
			assert.Equal(t, tt.expect.LastName, user.LastName)
			assert.Equal(t, tt.expect.Email, user.Email)
			assert.Equal(t, tt.expect.PasswordHash, user.PasswordHash)

			// Additional assertion to verify the user is actually stored in the database
			var dbUser User
			if err := db.First(&dbUser, "id = ?", user.ID).Error; err != nil {
				t.Fatalf("user not found in the database: %v", err)
			}
			assert.Equal(t, user.FirstName, dbUser.FirstName)
			assert.Equal(t, user.LastName, dbUser.LastName)
			assert.Equal(t, user.Email, dbUser.Email)
			assert.Equal(t, user.PasswordHash, dbUser.PasswordHash)
		})
	}

	if err = db.Exec("DROP TABLE users;").Error; err != nil {
		t.Fatal(err)
	}
}

func Test_GetUserByEmail(t *testing.T) {
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

	db.AutoMigrate(&User{})

	a := AppUsers{DB: db}

	expectedUser := User{ID: "testID", FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", PasswordHash: "hashedpassword"}
	db.Create(&expectedUser)

	tests := []struct {
		testName string
		email    string
		wantErr  error
		expect   User
	}{
		{
			testName: "valid email",
			email:    "john.doe@example.com",
			wantErr:  nil,
			expect:   expectedUser,
		},
		{
			testName: "empty email",
			email:    "",
			wantErr:  errors.New("no email passed"),
			expect:   User{},
		},
		{
			testName: "user not found",
			email:    "nonexistent@example.com",
			wantErr:  errors.New("user not found"),
			expect:   User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user, err := a.GetUserByEmail(tt.email)

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

			assert.Equal(t, tt.expect.FirstName, user.FirstName)
			assert.Equal(t, tt.expect.LastName, user.LastName)
			assert.Equal(t, tt.expect.Email, user.Email)
			assert.Equal(t, tt.expect.PasswordHash, user.PasswordHash)
		})
	}

	if err = db.Exec("DROP TABLE users;").Error; err != nil {
		t.Fatal(err)
	}
}

func Test_GetUserByID(t *testing.T) {
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

	db.AutoMigrate(&User{})

	a := AppUsers{DB: db}

	// Prepare a user for testing
	expectedUser := User{ID: "testID", FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", PasswordHash: "hashedpassword"}
	db.Create(&expectedUser)

	tests := []struct {
		testName string
		ID       string
		wantErr  error
		expect   User
	}{
		{
			testName: "valid ID",
			ID:       "testID",
			wantErr:  nil,
			expect:   expectedUser,
		},
		{
			testName: "empty ID",
			ID:       "",
			wantErr:  errors.New("no ID passed"),
			expect:   User{},
		},
		{
			testName: "user not found",
			ID:       "nonexistentID",
			wantErr:  errors.New("user not found"),
			expect:   User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user, err := a.GetUserByID(tt.ID)

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

			assert.Equal(t, tt.expect.FirstName, user.FirstName)
			assert.Equal(t, tt.expect.LastName, user.LastName)
			assert.Equal(t, tt.expect.Email, user.Email)
			assert.Equal(t, tt.expect.PasswordHash, user.PasswordHash)
		})
	}

	if err = db.Exec("DROP TABLE users;").Error; err != nil {
		t.Fatal(err)
	}
}
