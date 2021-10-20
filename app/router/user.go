package router

import (
	"github.com/gin-gonic/gin"
	"go-fast/app/handler"
)

func RegisterUser(router *gin.Engine) {
	g := router.Group("/v1/user")
	{
		g.GET(":id", handler.GetUser)
	}
}
