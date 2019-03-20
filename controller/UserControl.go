package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"

	"net/http"
	"strconv"
)

type UserController struct {
}

func DoLogin(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户id错误"})

		return
	}
	password := ctx.PostForm("password")
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码不能为空"})
		return
	}
	user := model.FindUser(id, password)
	if user != nil {

		ctx.JSON(http.StatusOK, gin.H{"message": "用户登录成功"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名或者密码错误"})
	}
}
