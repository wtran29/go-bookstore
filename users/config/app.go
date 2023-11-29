package config

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
}

var (
	router = gin.Default()
)

func StartApplication() {

	routes()
	router.Run(":8080")
}
