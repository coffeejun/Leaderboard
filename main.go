package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个实例
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	r.Run(":9999")
}
