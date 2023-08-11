package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"mt/models"
	"mt/utils"
	"sync"
)

// FriendWebSRequest 判断web socket请求id token是否正确
func FriendWebSRequest(mineId string, token string) bool {
	var user models.User
	utils.Db.Where("id= ?", mineId).Find(&user)
	if user.Tocken == token {
		return true
	} else {
		return false
	}
}

// SendFriendWebSRequest 发送好友请求
func SendFriendWebSRequest(mineId string, friendId string) bool {
	var (
		mu sync.Mutex //f互斥锁
		f  models.User
	)
	if a, b := models.FWebRequestMap[friendId]; b {
		mu.Lock()
		utils.Db.Where("id=?", mineId).Find(&f)
		mu.Unlock()
		if err := a.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			// 发送心跳消息失败，关闭连接
			Exit(friendId)
			return false
		}
		err := a.WriteJSON(gin.H{
			"msg":  "1",
			"data": f,
		})
		if err != nil {
			fmt.Println("writeJSON err:", err)
			return false
		}
		return true
	}
	return false
}
