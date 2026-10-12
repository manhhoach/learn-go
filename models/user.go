package models

import (
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
