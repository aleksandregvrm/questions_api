package models

import (
	"encoding/json"
	"fmt"

	database "example.com/questions/db"
	"example.com/questions/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64
	Email    string
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) AddUser() error {
	query := "INSERT INTO users(email, username, password) VALUES ($1, $2, $3)"
	stmt, err := database.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hPsw, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, u.Username, hPsw)

	if err != nil {
		return err
	}
	return err
}

func (u *User) LoginUser() error {
	if len(u.Password) < 8 {
		return utils.NewCustomError("Invalid credentials", 400)
	}
	dbUser, err := u.getSingleUser()
	if err != nil {
		fmt.Println(err)
		return utils.NewCustomError("Invalid Credentials", 400)
	}
	fmt.Println(dbUser)
	comparedPassBoolean := utils.ComparePasswords(u.Password, dbUser.Password)
	if !comparedPassBoolean {
		return utils.NewCustomError("Invalid Credentials", 400)
	}
	return nil
}

func (u *User) getSingleUser() (*User, error) {
	var user User
	query := "SELECT username,email,password,id FROM users WHERE username = $1"
	row := database.Db.QueryRow(query, u.Username)
	err := row.Scan(&user.Username, &user.Email, &user.Password, &user.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

func UpdateUserInfo(ctx *gin.Context) {
	body := User{}
	data, err := ctx.GetRawData()

	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}

	err = json.Unmarshal(data, &body)
}
