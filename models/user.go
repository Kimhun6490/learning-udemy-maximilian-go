package models

import (
	"errors"

	"example.com/gin/db"
	"example.com/gin/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id

	return nil
}

func (u *User) ValidateCredentials() error {
	var storedPassword string
	query := `
	SELECT id, password FROM users WHERE email = ?
	`
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &storedPassword)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(u.Password, storedPassword) {
		return errors.New("invalid credentials")
	}

	return nil
}
