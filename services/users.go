package services

import (
	"errors"
	"log"
)

func (u *AppUsers) CreateUser(email, passwordHash string) (User, error) {
	if email == "" {
		return User{}, errors.New("no email passed")
	}

	if passwordHash == "" {
		return User{}, errors.New("no password hash passed")
	}

	query := "INSERT INTO Users (Email, PasswordHash) VALUES (?, ?)"
	_, err := u.DB.Exec(query, email, passwordHash)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return User{}, err
	}
	return User{Email: email, PasswordHash: passwordHash}, nil
}

func (u *AppUsers) GetUser(email string) (User, error) {
	if email == "" {
		return User{}, errors.New("no email passed")
	}

	query := "SELECT PasswordHash FROM Users WHERE Email = ?"
	row := u.DB.QueryRow(query, email)
	var pwdHash string
	err := row.Scan(&pwdHash)

	if err != nil {
		return User{}, err
	}

	return User{Email: email, PasswordHash: pwdHash}, nil
}

func (u *AppUsers) DeleteUser(email string) error {
	return nil
}

func (u *AppUsers) UpdateUser(column, data string) error {
	return nil
}
