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

	users := make([]User, 0)
	Db.Find(&users)
	fmt.Println("dada", users)
	var t test
	db.First(&t)
	fmt.Println("dada", t)

	var t1 test
	db.Last(&t1)
	fmt.Println("dada", t1)

	tt := make([]test, 0)
	db.Find(&tt)
	fmt.Println("dada", tt)

	uu := make([]User, 0)
	db.Find(&uu)
	fmt.Println("dada", uu)
}
