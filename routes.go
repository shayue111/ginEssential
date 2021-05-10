package main

import (
	"github.com/gin-gonic/gin"
	"shayue/ginessential/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
