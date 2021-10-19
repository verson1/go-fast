package app

import (
	"github.com/gin-gonic/gin"
	"go-fast/app/router"
	"go-fast/config"
)

func NewApiHttpServer() {
	r := gin.Default()
	r.Use()

	router.RegisterUser(r)

	r.Run(config.Conf.ListenPort)
}
