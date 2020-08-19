package handler

import (
	dao "app1/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

//  /user/register  用户注册
func Register(c *gin.Context) {
	//获取用户名和密码
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	passwordAgain := c.Request.FormValue("password-again")

	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名或密码为空!",
		})
		return
	}
	if password != passwordAgain {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "两次密码不一致!",
		})
		return
	}
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.GetUserByName(username)
	if user.Id > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名已存在",
		})
		return
	}
	//用户名可用，将用户信息保存到数据库中
	err := dao.SaveUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
// /user/login
func Login(c *gin.Context){
	//获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0{
		c.JSON(http.StatusForbidden, gin.H{
			"msg":"用户名或密码为空",
		})
		return
	}
	user, _ := dao.GetUserByNameAndPassword(username, password)
	if user.Id > 0{
		// 设置cookie，标识用户是否登录
		c.SetCookie("user_cookie", string(user.Id), 10000, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"msg":"登陆成功",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":"用户名或密码错误",
	})
}

// /user/getAll
func GetAllUser(c *gin.Context) {
	users, _ := dao.GetAllUser()
	c.JSON(http.StatusOK, gin.H{
		"users":users,
	})
}