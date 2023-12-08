package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wtran29/go-bookstore/auth/src/domain/token"
	"github.com/wtran29/go-bookstore/auth/src/utils/errors"
)

type TokenHandler interface {
	GetTokenByID(ctx *gin.Context)
	CreateToken(ctx *gin.Context)
	UpdateTokenExpiry(ctx *gin.Context)
}

type tokenHandler struct {
	service token.Service
}

// GetTokenByID implements TokenHandler.
func (h *tokenHandler) GetTokenByID(ctx *gin.Context) {
	token, err := h.service.GetTokenByID(ctx.Param("token_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func (h *tokenHandler) CreateToken(ctx *gin.Context) {
	var token token.Token
	if err := ctx.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}
	if err := h.service.CreateToken(token); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, token)
}

func (h *tokenHandler) UpdateTokenExpiry(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "unimplemented")
}

func NewHandler(s token.Service) TokenHandler {
	return &tokenHandler{
		service: s,
	}
}
