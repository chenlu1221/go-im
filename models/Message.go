package models

type Message struct {
	//谁发的
	UserId int64 `bson:"userId"`
	//对端用户ID/群ID
	DstId string `bson:"dstId"`
	//群聊还是私聊
	Cmd string `bson:"cmd"`
	//消息按照什么样式展示
	Media string `bson:"media"`
	//消息的内容
	Content string `bson:"content"`
	//预览图片
	Pic string `bson:"pic"`
	//服务的URL
	Url string `bson:"url"`
	//简单描述
	Memo string `bson:"memo"`
	//其他和数字相关的,语音长度/红包金额
	Amount int `bson:"amount"`
}

const (
	//点对点单聊,dstid是用户ID
	CMD_SINGLE_MSG = "10"
	//群聊消息,dstid是群id
	CMD_ROOM_MSG = "11"
)
const (
	MEDIA_TYPE_TEXT = "1" //文本样式

	MEDIA_TYPE_VOICE = "3" //语音样式

	MEDIA_TYPE_IMG = "4" //图片样式

	MEDIA_TYPE_EMOJ = "6" //emoj表情样式

	MEDIA_TYPE_LINK = "7" //超链接样式

	MEDIA_TYPE_VIDEO = "8" //视频样式
)

func (Message) TableName() string {
	return "message"
}
