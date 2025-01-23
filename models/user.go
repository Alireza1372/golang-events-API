package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)
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
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievesPassword string
	err := row.Scan(&u.ID, &retrievesPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievesPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil

}
