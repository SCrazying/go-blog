package model

import (
	"time"
)

type User struct {
	Id         int `gorm:"primary_key:true"`
	Username   string
	Password   string
	Email      string
	CreateTime time.Time
	Isactive   int
}

func (User) TableName() string {
	return "user"
}

//使用用户名和密码查询用户
func FindUser(id int, password string) *User {
	user := &User{}
	flag := Db.Where("id=? AND password=?", id, password).Find(&user).RecordNotFound()
	if flag {
		return nil
	} else {
		return user
	}
}
