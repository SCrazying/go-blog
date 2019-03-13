package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminPage(ctx *gin.Context) {

	//需要加入session校验
	ctx.HTML(http.StatusOK, "adminindex.html", gin.H{})
}

func BlogIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func AdminLogin(ctx *gin.Context) {
	//需要加入session校验
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func AdminRegister(ctx *gin.Context) {
	//需要加入session校验
	ctx.HTML(http.StatusOK, "register.html", gin.H{})
}
