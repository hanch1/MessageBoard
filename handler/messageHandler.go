package handler

import (
	"app1/dao"
	"app1/model"
	"app1/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// /msg/getAll
func GetMsgs(c *gin.Context) {
	// 分页参数
	pageIndex := c.Query("pageIndex")
	pageSize := c.Query("pageSize")
	if len(pageIndex) == 0 || len(pageSize) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"mag": "参数为空",
		})
		return
	}
	index, err1 := strconv.Atoi(pageIndex)
	size, err2 := strconv.Atoi(pageSize)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	msgs, _ := dao.GetMessages(index, size)
	c.JSON(http.StatusOK, gin.H{
		"msg": msgs,
	})
}

func AddMsg(c *gin.Context) {

	cookie, err := c.Request.Cookie("user_cookie")
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "请登录",
		})
		return
	}

	content := c.Request.FormValue("content")
	// 是否开启关键词屏蔽
	flag := c.Request.FormValue("isCheck")
	// 开启关键词屏蔽
	if flag == "true" {
		t := util.NewTrie()
		t.Insert("警察")
		t.Insert("枪支")
		t.Insert("垃圾")
		content = t.Replace(content)
	}
	if len(content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数不正确",
		})
		return
	}
	// 从cookie中获取用户id   %07
	id := getUidFromCookie(cookie.Value)
	username, _ := dao.GetUsernameByUid(id)
	msg := &model.Message{
		UserName: username,
		Content:  content,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
	}

	err2 := dao.AddMessage(msg)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "添加成功",
	})
}

func getUidFromCookie(cookie string) int {
	str := cookie[1:]
	id, _ := strconv.Atoi(string(str))
	return id
}
