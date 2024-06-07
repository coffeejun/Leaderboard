package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "111")
	})
	r.Run() // 在 0.0.0.0:8080 上启动服务

}
