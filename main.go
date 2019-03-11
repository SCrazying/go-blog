package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/config"
)

func main(){

	router := gin.Default()

	//配置静态文件路径
	router.Static("/public","./public")
	cfg:= &config.Config{}
	fmt.Println(cfg)

	router.Run(":2333")
}
