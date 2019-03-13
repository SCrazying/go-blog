package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/config"
	"go-blog/controller"
	_ "go-blog/model"
)

func main() {

	router := gin.Default()

	//配置静态文件路径
	router.Static("/public", "./public")

	//用户的路由绑定
	router.GET("/user", controller.Login)
	router.POST("/user", controller.Register)
	router.GET("/users", controller.AllUser)
	router.DELETE("/user", controller.DelUser)

	router.Run(config.Cfg.GetServer())
}
