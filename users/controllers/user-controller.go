package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	authLib "github.com/wtran29/auth-lib/auth"
	"github.com/wtran29/go-bookstore/resterr"
	"github.com/wtran29/go-bookstore/users/domain/users"
	"github.com/wtran29/go-bookstore/users/services"
)

func getUserId(id string) (int64, *resterr.JsonError) {
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, resterr.NewBadRequestError("user id should be a number", err)

	}
	return userId, nil
}

func CreateUser(ctx *gin.Context) {
	var user users.User

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
		jsonErr := resterr.NewBadRequestError("invalid json body", err)
		ctx.JSON(jsonErr.Status, jsonErr)
		return
	}

	result, err := services.UsersService.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, result.ReadJson(ctx.GetHeader("X-Public") == "true"))

}
func GetUser(ctx *gin.Context) {
	if err := authLib.AuthenticateRequest(ctx.Request); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	if callerId := authLib.GetCallerId(ctx.Request); callerId == 0 {
		err := resterr.JsonError{
			Status:  http.StatusUnauthorized,
			Message: "resource not available",
		}
		ctx.JSON(err.Status, err)
		return
	}
	userId, err := getUserId(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	if authLib.GetCallerId(ctx.Request) == user.ID {
		ctx.JSON(http.StatusOK, user.ReadJson(false))
		return
	}
	ctx.JSON(http.StatusOK, user.ReadJson(authLib.IsPublic(ctx.Request)))

}

func UpdateUser(ctx *gin.Context) {
	userId, err := getUserId(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		jsonErr := resterr.NewBadRequestError("invalid json body", err)
		ctx.JSON(jsonErr.Status, jsonErr)
		return
	}
	user.ID = userId

	isPartial := ctx.Request.Method == http.MethodPatch

	updatedUser, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		ctx.JSON(updateErr.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, updatedUser.ReadJson(ctx.GetHeader("X-Public") == "true"))
}

func DeleteUser(ctx *gin.Context) {
	userId, err := getUserId(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	if err := services.UsersService.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("user %d deleted", userId), "status": "204", "errors": "false"})
}

func SearchUser(ctx *gin.Context) {
	status := ctx.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	result := users.ReadJson(ctx.GetHeader("X-Public") == "true")
	ctx.JSON(http.StatusOK, result)
}

func Login(ctx *gin.Context) {
	var login users.Login
	// fmt.Println(login)
	if err := ctx.ShouldBindJSON(&login); err != nil {
		jsonErr := resterr.NewBadRequestError("invalid json body", err)
		ctx.JSON(jsonErr.Status, jsonErr)
		return
	}

	user, err := services.UsersService.LoginUser(login)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user.ReadJson(ctx.GetHeader("X-Public") == "true"))

}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
