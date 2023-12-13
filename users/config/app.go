package config

import (
	"github.com/gin-gonic/gin"
	"github.com/wtran29/go-bookstore/users/logger"
)

type Config struct {
}

var (
	router = gin.Default()
)

func StartApplication() {

	routes()
	logger.Info("starting the application")
	router.Run(":8081")
}
