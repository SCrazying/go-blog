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

func Register(username, password, email string) (int, string) {

	user := &User{}
	Db.Where("username=?", username).First(user)
	if user.Username == username {
		//用户已经存在
		return 2, "用户已经存在"
	} else {
		user := User{
			Id:       111,
			Username: username,
			Password: password,
			Email:    email,
		}
		Db.Create(&user)
		return 0, "用户创建成功"
	}
	return -1, "系统错误"
}

func Login(username, password string) *User {
	user := &User{}
	Db.Where("username=? and password=?", username, password).First(user)
	return user
}
