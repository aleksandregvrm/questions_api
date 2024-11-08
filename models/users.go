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

func AddUser(userData *User) error {
	body := User{}
	byteData, err := json.Marshal(userData)
	if err != nil {
		return utils.NewCustomError("Internal server error", 500)
	}
	err = json.Unmarshal(byteData, body)
	if err != nil {
		return utils.NewCustomError("Invalid data", 400)
	}
	_, err = database.Db.Exec("insert into users(username,password) values ($1,$2)", body.Username, body.Password)
	if err != nil {
		fmt.Println(err)
		return utils.NewCustomError("couldn't create the new user", 500)
	} else {
		return nil
	}

}

func LoginUser(ctx *gin.Context) {
	body := User{}
	data, err := ctx.GetRawData()

	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}
	err = json.Unmarshal(data, &body)
	if len(body.Password) < 8 {
		ctx.AbortWithStatusJSON(400, "Invalid credentials, password should be longer")
		return
	}

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
