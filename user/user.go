package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot-gin"
)

func InitUserRouter(e *gin.Engine) {
	e.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "user1"})
	})
}

func init() {
	boot.RegisterRouter(InitUserRouter)
}
