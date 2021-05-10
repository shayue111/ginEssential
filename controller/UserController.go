package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"shayue/ginessential/common"
	"shayue/ginessential/model"
	"shayue/ginessential/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	/*
		数据验证
	*/
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位！"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位！"})
		return
	}

	// 如果name未传，随机初始化字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在！"})
		return
	}

	// 创建用户
	newUser := &model.User{
		Model:     gorm.Model{},
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}

	DB.Create(newUser)

	// 返回结果
	ctx.JSON(200, gin.H{"msg": "注册成功"})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	// 这里将符合条件的数据的第一条导入赋值给user
	db.Where("telephone = ?", telephone).First(&user)
	// 由于user.ID初始化为0，若此时user.ID != 0，则说明确实存在满足该条件的数据
	if user.ID != 0 {
		return true
	}

	return false
}
