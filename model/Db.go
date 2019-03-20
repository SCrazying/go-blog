package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-blog/config"
)

var Db *gorm.DB

func init() {
	db, err := gorm.Open(config.Cfg.GetUrl())
	if err != nil {
		fmt.Println("数据库连接失败 ", err)
		panic(err)
	} else {
		fmt.Println("连接db成功")
		Db = db
	}
}
