package models

import (
	"fmt"
	"time"
)

type User struct {
	Id       int64
	Mobile   string `gorm:"varchar(11) NOT NULL COMMENT '手机号';primary_key"`
	Passwd   string `gorm:"varchar(255) DEFAULT NULL COMMENT '密码'"`
	Avatar   string //头像
	Nickname string //昵称
	Tocken   string //鉴权因子
	Create   time.Time
	Salt     string //加密因子，跟在Passwd后面的随机数
	Email    string
}

func (User) TableName() string {
	return "user"
}

var UserMap = make(map[int64]User)
var TextChanMap = make(map[string]chan any)
var VoiceChanMap = make(map[string]chan any) //语音
var ImgChanMap = make(map[string]chan any)
var EmojChanMap = make(map[string]chan any)
var LinkChanMap = make(map[string]chan any)
var VideoChanMap = make(map[string]chan any)

func InitChanMap(id string) {
	TextChanMap[id] = make(chan any, 8)
	VoiceChanMap[id] = make(chan any, 8)
	ImgChanMap[id] = make(chan any, 8)
	EmojChanMap[id] = make(chan any, 8)
	LinkChanMap[id] = make(chan any, 8)
	VideoChanMap[id] = make(chan any, 8)
	fmt.Println(id + "消息通道已经初始化")
}
func CreateChanMap(id string) {
	delete(TextChanMap, id)
	delete(VoiceChanMap, id)
	delete(ImgChanMap, id)
	delete(EmojChanMap, id)
	delete(LinkChanMap, id)
	delete(VideoChanMap, id)
}
