package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wtran29/go-bookstore/users/domain/users"
	"github.com/wtran29/go-bookstore/users/services"
	"github.com/wtran29/go-bookstore/users/utils/errors"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	fmt.Println(user)
	// bytes, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	// TODO: Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	// TODO: handle json error
	// 	return
	// }
	if err := ctx.ShouldBindJSON(&user); err != nil {
		jsonErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(jsonErr.Status, jsonErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	fmt.Println(user)
	ctx.JSON(http.StatusCreated, result)

}
func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "to be implemented")
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "to be implemented")

}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
