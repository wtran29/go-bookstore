package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wtran29/go-bookstore/users/domain/users"
	"github.com/wtran29/go-bookstore/users/services"
	"github.com/wtran29/go-bookstore/users/utils/errors"
)

func getUserId(id string) (int64, *errors.JsonError) {
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")

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
		jsonErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(jsonErr.Status, jsonErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, result)

}
func GetUser(ctx *gin.Context) {
	userId, err := getUserId(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func UpdateUser(ctx *gin.Context) {
	userId, err := getUserId(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		jsonErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(jsonErr.Status, jsonErr)
		return
	}
	user.ID = userId

	isPartial := ctx.Request.Method == http.MethodPatch

	updatedUser, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		ctx.JSON(updateErr.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(ctx *gin.Context) {
	userId, err := getUserId(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "to be implemented")

}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
