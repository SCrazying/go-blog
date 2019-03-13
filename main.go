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
func main() {

	router := gin.Default()

	// 定义模板相关设置
	setTemplate(router)
	//配置静态文件路径
	router.Static("/public/assert", "./public/assert")

	//用户的路由绑定
	router.GET("/user", controller.Login)
	router.POST("/user", controller.Register)
	router.GET("/users", controller.AllUser)
	router.DELETE("/user", controller.DelUser)

	//添加后台入口的登陆
	router.GET("/admin", controller.AdminPage)
	router.GET("/login", controller.AdminLogin)
	router.GET("/register", controller.AdminRegister)

	err := router.Run(config.Cfg.GetServer())
	if err != nil {
		log.Fatalln(err)
	}

}
