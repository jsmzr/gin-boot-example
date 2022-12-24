package main

import (
	"fmt"

	boot "github.com/jsmzr/boot-gin"
	_ "github.com/jsmzr/boot-plugin-gin-swagger"
	_ "github.com/jsmzr/boot-plugin-logrus"
	_ "github.com/jsmzr/boot-plugin-prometheus"
	_ "github.com/jsmzr/boot-plugin-skywalking"
	_ "github.com/jsmzr/gin-boot-example/demo"
	_ "github.com/jsmzr/gin-boot-example/docs"
	_ "github.com/jsmzr/gin-boot-example/middlewares"
	_ "github.com/jsmzr/gin-boot-example/user"
)

func main() {
	if err := boot.Run(); err != nil {
		fmt.Println(err)
	}
}
