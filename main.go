package main

import (
	"github.com/gin-gonic/gin"
	"shayue/ginessential/common"
)

func main() {
	// 初始化数据库
	common.InitDB()

	r := gin.Default()

	r = CollectRoute(r)

	// 监听并在 0.0.0.0:8080 上启动服务
	panic(r.Run())
}
