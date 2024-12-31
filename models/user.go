package models

import (
	"HumptyDumpy01/go-restful-api/db"
	"HumptyDumpy01/go-restful-api/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	query := `
	INSERT INTO users (email, password) VALUES (?, ?)
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	u.Password = hashedPassword

	return err
}

func (u User) ValidateCredentials() error {
	query := `
SELECT password FROM users WHERE email = ?
`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("invalid email or password")
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid email or password")
	}
	return nil
}
