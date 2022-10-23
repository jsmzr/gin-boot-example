package demo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot-gin"
)

func InitDemoRouter(e *gin.Engine) {
	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "pong"})
	})
}

func init() {
	boot.RegisterRouter(InitDemoRouter)
}
