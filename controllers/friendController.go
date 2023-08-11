package controllers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mt/models"
	"mt/services"
)

// GetFriendById 获取好友列表
func GetFriendById(c *gin.Context) {
	id := c.PostForm("id")
	friend, i := services.GetFriend(id)
	if i == 1 {
		c.JSON(200, gin.H{
			"msg": 1,
			"err": "好友列表为空",
		})
	} else if i == 0 {
		var users []models.FJson
		for _, m := range friend {
			name, avatar, id := services.GetUser(m.FriendId)
			if avatar != "1" {
				data, err := ioutil.ReadFile("C:/Users/Administrator/Desktop/go/img/" + avatar)
				base64Str := base64.StdEncoding.EncodeToString(data)
				if err != nil {
					log.Fatal(err)
				}
				users = append(users, models.FJson{Nickname: name, Avatar: "data:image/jpeg;base64," + base64Str, Id: id})
			} else {
				users = append(users, models.FJson{Nickname: name, Avatar: avatar, Id: id})
			}
		}
		c.JSON(200, gin.H{
			"msg":  0,
			"data": users,
		})
	}
}

// AddFriendById 发送添加好友请求
func AddFriendById(c *gin.Context) {
	mineId := c.PostForm("mineId")
	friendId := c.PostForm("friendId")
	f := services.AddFriend(mineId, friendId)
	if f == 0 {
		c.JSON(200, gin.H{
			"msg":  0,
			"data": "发送成功",
		})
	}
	if f == 1 {
		c.JSON(200, gin.H{
			"msg":  1,
			"data": "请勿重复发送",
		})
	}
	if f == 2 {
		c.JSON(200, gin.H{
			"msg":  1,
			"data": "你们已经是好友",
		})
	}
}
