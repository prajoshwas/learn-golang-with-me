package models

import (
	"errors"
	"example/rest-api/db"
	"example/rest-api/utils"
	"fmt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) SaveUser() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID

	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var storedPassword string

	err := row.Scan(&u.ID, &storedPassword)
	fmt.Println("Stored password", storedPassword)
	fmt.Println("Input password", u.Password)
	if err != nil {
		return err
	}

	isValid := utils.UnhashPassword(storedPassword, u.Password)
	fmt.Println("Is valid", isValid)
	if !isValid {
		return errors.New("Invalid credentials")
	}

	return nil
}
