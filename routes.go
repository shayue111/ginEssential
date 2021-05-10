package main

import (
	"github.com/gin-gonic/gin"
	"shayue/ginessential/controller"
	"shayue/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 用户注册
	r.POST("/api/auth/register", controller.Register)

	// 用户登入
	r.POST("/api/auth/login", controller.Login)

	// 用户信息验证
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
