package services

import (
	"errors"
	"log"
)

func (u *AppUsers) CreateUser(firstName, lastName, email, passwordHash string) (User, error) {
	if email == "" {
		return User{}, errors.New("no email passed")
	}

	if passwordHash == "" {
		return User{}, errors.New("no password hash passed")
	}

	ID := generateID()

	query := "INSERT INTO Users (FirstName, LastName, ID, Email, PasswordHash) VALUES (?, ?, ?, ?, ?)"
	_, err := u.DB.Exec(query, firstName, lastName, ID, email, passwordHash)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return User{}, err
	}
	return User{FirstName: firstName, LastName: lastName, ID: ID, Email: email, PasswordHash: passwordHash}, nil
}

func (u *AppUsers) GetUserByEmail(email string) (User, error) {
	if email == "" {
		return User{}, errors.New("no email passed")
	}

	query := "SELECT FirstName, LastName, ID, Email, PasswordHash FROM Users WHERE Email = ?"
	row := u.DB.QueryRow(query, email)
	var user User
	err := row.Scan(&user.FirstName, &user.LastName, &user.ID, &user.Email, &user.PasswordHash)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *AppUsers) GetUserByID(ID string) (User, error) {
	if ID == "" {
		return User{}, errors.New("no email passed")
	}

	query := "SELECT FirstName, LastName, ID, Email, PasswordHash FROM Users WHERE ID = ?"
	row := u.DB.QueryRow(query, ID)
	var user User
	err := row.Scan(&user.FirstName, &user.LastName, &user.ID, &user.Email, &user.PasswordHash)

	if err != nil {
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
