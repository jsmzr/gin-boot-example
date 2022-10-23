package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot-gin"
	log "github.com/jsmzr/boot-log"
)

type DemoMiddleware struct{}

func (d *DemoMiddleware) Load(e *gin.Engine) error {
	e.Use(func(ctx *gin.Context) {
		log.Info("DemoMiddleware start")
		ctx.Next()
		log.Info("DemoMiddleware end")
	})
	return nil
}

func (d *DemoMiddleware) Order() int {
	return 0
}

func init() {
	boot.RegisterMiddleware("demo", &DemoMiddleware{})
}
