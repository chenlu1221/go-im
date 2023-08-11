package models

import (
	"github.com/gorilla/websocket"
)

type FriendQuest struct {
	Id            int64
	UserId        string
	FriendId      string
	RequestStatus string
}

func (FriendQuest) TableName() string {
	return "frrequest"
}

var FWebRequestMap = make(map[string]*websocket.Conn)

type FriendR struct {
	Id            string
	RequestStatus string
}
