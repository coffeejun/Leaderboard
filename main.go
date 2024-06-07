package main

import "gorank/router"

func main() {
	//创建一个实例
	r := router.Router()
	r.Run(":9999")
}
