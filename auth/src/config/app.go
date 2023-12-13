package config

import (
	"github.com/gin-gonic/gin"
	"github.com/wtran29/go-bookstore/auth/src/handlers"
	"github.com/wtran29/go-bookstore/auth/src/repository/database"
	"github.com/wtran29/go-bookstore/auth/src/repository/rest"
	"github.com/wtran29/go-bookstore/auth/src/services"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// session, dbErr := cassandra.GetSession()
	// if dbErr != nil {
	// 	panic(dbErr)
	// }
	// session.Close()

	tokenHandler := handlers.NewHandler(services.NewService(rest.NewRepo(), database.NewRepo()))

	router.GET("/oauth/access_token/:token_id", tokenHandler.GetTokenByID)
	router.POST("/oauth/access_token", tokenHandler.CreateToken)
	router.PUT("/oauth/token_expiry", tokenHandler.UpdateTokenExpiry)
	router.Run(":8080")
}
