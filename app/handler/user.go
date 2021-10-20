package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println(name)
}
