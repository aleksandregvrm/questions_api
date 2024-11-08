package models

import (
	"encoding/json"
	"fmt"

	database "example.com/questions/db"
	"example.com/questions/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) AddUser() error {
	_, err := database.Db.Exec("insert into users(username,password) values ($1,$2)", u.Username, u.Password)
	if err != nil {
		fmt.Println(err)
		return utils.NewCustomError("couldn't create the new user", 500)
	} else {
		return nil
	}

}

func (u *User) LoginUser() error {
	if len(u.Password) < 8 {
		return utils.NewCustomError("Invalid credentials", 400)
	}
	dbUser, err := getSingleUser(u.Username)
	if err != nil {
		return utils.NewCustomError("Invalid Credentials", 400)
	}
	if !utils.ComparePasswords(u.Password, dbUser.Password) {
		return utils.NewCustomError("Invalid Credentials", 400)
	}
	return nil
}

func getSingleUser(username string) (*User, error) {
	var user User
	query := "SELECT * FROM users WHERE username = ?"
	row := database.Db.QueryRow(query, username)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
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
