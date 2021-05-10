package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shayue/ginessential/common"
	"shayue/ginessential/model"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		tokens, claim, err := common.ParseToken(tokenString)

		if err != nil || !tokens.Valid {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 通过验证后，获取claim中的userID
		userID := claim.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)

		// 验证用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 将user的信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
