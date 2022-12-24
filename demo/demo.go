package demo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot-gin"
)

// @Summary swagger-demo
// @Tags demo
// @Router /ping [get]
// @Success 200 {string} pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func InitDemoRouter(e *gin.Engine) {
	e.GET("/ping", Ping)
}

func init() {
	boot.RegisterRouter(InitDemoRouter)
}
