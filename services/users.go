package services

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (u *AppUsers) CreateUser(firstName, lastName, email, passwordHash string) (User, error) {
	if email == "" {
		return User{}, errors.New("no email passed")
	}

	if passwordHash == "" {
		return User{}, errors.New("no password hash passed")
	}

	ID := generateID()

	user := User{
		FirstName:    firstName,
		LastName:     lastName,
		ID:           ID,
		Email:        email,
		PasswordHash: passwordHash,
	}

	if err := u.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		return User{}, err
	}

	return user, nil
}

func (u *AppUsers) GetUserByEmail(email string) (User, error) {
	if email == "" {
		return User{}, errors.New("no email passed")
	}

	var user User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}

	return user, nil
}

func (u *AppUsers) GetUserByID(ID string) (User, error) {
	if ID == "" {
		return User{}, errors.New("no ID passed")
	}

	var user User
	if err := u.DB.First(&user, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}

	return user, nil
}

func (u *AppUsers) DeleteUser(email string) error {
	return nil
}

func (u *AppUsers) UpdateUser(column, data string) error {
	return nil
}
