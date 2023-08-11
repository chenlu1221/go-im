package models

type Friend struct {
	Id       int64  `gorm:"column:Id;bigint(20) NOT NULL AUTO_INCREMENT"`
	UserId   string `gorm:"column:UserId;varchar(255) DEFAULT NULL COMMENT '用户id'"`
	FriendId string `gorm:"column:FriendId;varchar(255) DEFAULT NULL COMMENT '好友id'"`
}

func (Friend) TableName() string {
	return "friend"
}

type FJson struct {
	Nickname string
	Avatar   string
	Id       string
}
