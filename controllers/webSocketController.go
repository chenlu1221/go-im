package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"mt/models"
	"mt/services"
	"mt/utils"
	"net/http"
	"strconv"
	"time"
)

var C = make(map[string]chan int)

func WebSocketLogin(context *gin.Context) {
	id := context.Query("id")
	var stop = make(chan int)
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return services.FriendWebSRequest(id, context.Query("token"))
		},
	}
	C[id] = make(chan int)
	//初始化id对应的消息通道
	models.InitChanMap(id)
	conn, err := upGrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Println(err)
		http.NotFound(context.Writer, context.Request)
		return
	}
	if a, b := models.FWebRequestMap[id]; b {
		fmt.Println("关闭链接")
		go func() {
			stop <- 1
		}()
		go send(conn, id, a)
	} else {
		models.FWebRequestMap[id] = conn
	}
	log.Println("连接建立成功", conn.RemoteAddr())
	// 启动心跳检测
	go startHeartbeat(id, stop)
	//启动消息侦听
	go clickMessage(id)
	//启动发送消息的多路复用
	go sendMessage(id)
}
func send(conn *websocket.Conn, id string, a *websocket.Conn) {
	select {
	case <-C[id]:
		a.Close()
		models.FWebRequestMap[id] = conn
	}
}

//心跳检测
func startHeartbeat(id string, stop chan int) {
	fmt.Println(id + "心跳检测")
	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			if _, b := models.FWebRequestMap[id]; b {
				if err := models.FWebRequestMap[id].WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					//发送心跳检测失败，关闭链接
					services.Exit(id)
					return
				}
			} else {
				return
			}
		case <-stop:
			C[id] <- 1
			return
		}
	}
}

//消息侦听
func clickMessage(id string) {
	for {
		_, p, err := models.FWebRequestMap[id].ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		msg := models.Message{}
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Println(err)
			return
		}

		cendMessage(msg)
	}
}

//消息处理
func cendMessage(msg models.Message) {
	//消息存入数据库
	fmt.Println(msg)
	_, err := utils.MongoDB.Database("go-im").Collection("message").InsertOne(context.Background(), msg)
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}
	//好友之间发送消息
	if msg.Cmd == models.CMD_SINGLE_MSG {
		parseInt, err := strconv.ParseInt(msg.DstId, 10, 64)
		if err != nil {
			fmt.Println("ParseInt err:", err)
			return
		}
		fmt.Println(1)
		fmt.Println(parseInt)
		//接收消息的用户是否在线
		_, b := models.UserMap[parseInt]
		if b {
			//判断消息类型，并把消息丢入通道
			fmt.Println(b)
			fmt.Println(msg.DstId + "在线")
			switch msg.Media {
			case "1":
				go func() {
					fmt.Println(msg.DstId + "在线1")
					models.TextChanMap[msg.DstId] <- msg
					println(len(models.TextChanMap[msg.DstId]))
				}()
			}
		}
	}
	//群聊消息
	if msg.Cmd == models.CMD_ROOM_MSG {
		//处理消息
	}
}

//发送消息的多路复用
func sendMessage(id string) {
	fmt.Println(2)
	fmt.Println(id)
	for {
		fmt.Println("for")
		select {
		case a := <-models.TextChanMap[id]: //文字消息
			fmt.Println(3)
			err := models.FWebRequestMap[id].WriteJSON(gin.H{
				"msg":  "0",
				"data": a,
			})
			if err != nil {
				fmt.Println("writeJSON err:", err)
				return
			}
		}
	}
}
