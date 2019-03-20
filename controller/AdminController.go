package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
}

func ShowLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
