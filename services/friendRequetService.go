package services

import (
	"mt/models"
	"mt/utils"
)

func GetFriendRequestList(id string) []models.FriendQuest {
	var fr []models.FriendQuest
	utils.Db.Where("friend_id=?", id).Find(&fr)
	return fr
}
func AgreeFriendRequest(id string, friendId string) error {
	tx := utils.Db.Model(&models.FriendQuest{}).Where("user_id=? AND friend_id=?", friendId, id).Update("request_status", "2")
	if tx.Error != nil {
		return tx.Error
	}
	err := AddFriendForSql(id, friendId) //添加好友记录
	if err != nil {
		return err
	}
	return nil
}
func StoreFriendRequest(id string, friendId string, code string) error {
	var re = models.FriendQuest{
		Id:            0,
		UserId:        id,
		FriendId:      friendId,
		RequestStatus: code,
	}
	tx := utils.Db.Create(&re)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
