package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mt/models"
	"mt/services"
)

// GetList 获取好友请求列表
func GetList(c *gin.Context) {
	id := c.PostForm("id")
	Flist := services.GetFriendRequestList(id)
	if len(Flist) == 0 {
		c.JSON(200, gin.H{
			"msg": 1,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  0,
			"data": Flist,
		})
	}
}

// Agree 同意好友申请
func Agree(c *gin.Context) {
	id := c.PostForm("id")
	friendId := c.PostForm("friendId")
	err := services.AgreeFriendRequest(id, friendId)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"err":  err,
		})
		return
	}
	var user models.FJson
	user.Id = id
	name, avatar, _ := services.GetUser(friendId)
	user.Nickname = name
	if avatar != "1" {
		data, err := ioutil.ReadFile("C:/Users/Administrator/Desktop/go/img/" + avatar)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		base64Str := base64.StdEncoding.EncodeToString(data)
		user.Avatar = "data:image/jpeg;base64," + base64Str
	} else {
		user.Avatar = "1"
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": user,
	})
}

// StoreFriend 存储新发来的好友申请
func StoreFriend(c *gin.Context) {
	id := c.PostForm("id")
	friendId := c.PostForm("friendId")
	code := c.PostForm("code")
	err := services.StoreFriendRequest(id, friendId, code)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"err":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
	})
}
