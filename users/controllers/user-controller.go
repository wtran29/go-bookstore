package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser() {

}

func CreateUser() {

}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
