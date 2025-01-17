package models

import (
	"HumptyDumpy01/go-restful-api/db"
	"HumptyDumpy01/go-restful-api/utils"
	"errors"
	"fmt"
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
	fmt.Println(`Printing id`, id)
	if err != nil {
		return err
	}
	u.ID = id
	u.Password = hashedPassword

	return err
}

func (u *User) ValidateCredentials() (int64, error) {
	query := `
	SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return 0, errors.New("invalid email or password")
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return 0, errors.New("invalid email or password")
	}
	return u.ID, nil
}
