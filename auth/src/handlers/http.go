package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wtran29/go-bookstore/auth/src/domain/token"
	"github.com/wtran29/go-bookstore/auth/src/services"
	"github.com/wtran29/go-bookstore/resterr"
)

type TokenHandler interface {
	GetTokenByID(ctx *gin.Context)
	CreateToken(ctx *gin.Context)
	UpdateTokenExpiry(ctx *gin.Context)
}

type tokenHandler struct {
	service services.TokenService
}

func NewHandler(s services.TokenService) TokenHandler {
	return &tokenHandler{
		service: s,
	}
}

// GetTokenByID implements TokenHandler.
func (h *tokenHandler) GetTokenByID(ctx *gin.Context) {
	token, err := h.service.GetTokenByID(ctx.Param("token_id"))
	fmt.Println(token)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func (h *tokenHandler) CreateToken(ctx *gin.Context) {
	var token token.TokenRequest
	if err := ctx.ShouldBindJSON(&token); err != nil {
		restErr := resterr.NewBadRequestError("bad request", errors.New("invalid json body"))
		ctx.JSON(restErr.Status, restErr)
		return
	}
	at, err := h.service.CreateToken(token)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, at)
}

func (h *tokenHandler) UpdateTokenExpiry(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "unimplemented")
}
