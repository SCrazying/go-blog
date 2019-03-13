package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"net/http"
)

//login
func Login(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	user := model.Login(username, password)
	c.JSON(http.StatusOK, user)
	if user.Username == username && user.Password == password {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录失败",
		})
	}
}

//注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	code, msg := model.Register(username, password, email)
	switch code {

	case 0:
		c.JSON(http.StatusCreated, gin.H{
			"message": msg,
		})
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})

		break
	}

}

//所有用户
func AllUser(c *gin.Context) {

	users := make([]model.User, 0)

	c.JSON(http.StatusOK, users)

}

//删除用户
func DelUser(c *gin.Context) {
	//userid := c.Query("userid")

	c.JSON(http.StatusNoContent, gin.H{

		"message": "用户删除成功",
	})
}
