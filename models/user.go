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

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer st.Close()
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := st.Exec(user.Email, hashPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	user.ID = userId
	return err

}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password from users where email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var password string
	err := row.Scan(&u.ID, &password)
	if err != nil {
		return err
	}
	validPassword := utils.ComparePassword(u.Password, password)
	if !validPassword {
		return errors.New("Wrong password")
	}
	return nil
}
