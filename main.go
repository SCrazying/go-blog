package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/config"
	_ "go-blog/model"

)

func main(){

	router := gin.Default()

	//配置静态文件路径
	router.Static("/public","./public")



	router.Run(config.Cfg.GetServer())
}
