package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/config"
)

func main() {

	router := gin.Default()

	//配置静态文件路径
	router.Static("/public", "./public")
	cfg := &config.Config{}
	fmt.Println(cfg)

	//获取配置文件里的服务器连接地址
	addr := config.Cfg.GetHttpAddr()
	fmt.Println(addr)
	//
	router.Run(addr)
}
