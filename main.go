package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/config"
	"go-blog/controller"
	_ "go-blog/model"
	"log"
)

func setTemplate(route *gin.Engine) {
	route.LoadHTMLGlob("./view/*")
}

type testinferface interface {
	rrr() string
}

type aaaa struct {
}

var aa = new(aaaa)

func main() {

	router := gin.Default()

	// 定义模板相关设置
	setTemplate(router)
	//配置静态文件路径
	router.Static("/public/assert", "./public/assert")
	router.Static("/public/klorofil/", "./public/klorofil/")

	//用户的路由绑定
	router.POST("/user/", controller.DoLogin)
	//router.POST("/user/id", controller.Register)
	//router.GET("/user", controller.AllUser)
	//router.DELETE("/user/:id", controller.DelUser)

	//添加后台入口的登陆
	router.GET("/login", controller.ShowLogin)
	//router.GET("/login", controller.AdminLogin)
	//router.GET("/register", controller.AdminRegister)

	err := router.Run(config.Cfg.GetServer())
	if err != nil {
		log.Fatalln(err)
	}

}
