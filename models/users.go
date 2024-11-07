package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "example.com/questions/db"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func AddUser(ctx *gin.Context) {
	body := User{}
	data, err := ctx.GetRawData()

	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	_, err = database.Db.Exec("insert into users(username,password) values ($1,$2)", body.Username, body.Password)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, "Couldn't create the new user.")
	} else {
		ctx.JSON(http.StatusOK, "User is successfully created.")
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
