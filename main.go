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
	r.POST("/user/list", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "user list")
	})
	r.PUT("/user/add", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "user add")
	})
	r.DELETE("/user/delete", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "user delete")
	})
	r.Run(":9999")
}
