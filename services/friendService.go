package services

import (
	"mt/models"
	"mt/utils"
)

//好友请求1未处理，2同意

func GetFriend(id string) ([]models.Friend, int) {
	var f []models.Friend
	s := utils.Db.Where("UserId = ?", id).Find(&f)
	if s.RowsAffected == 0 {
		return nil, utils.NoCode
	}
	return f, utils.YesCode
}
func AddFriend(mineId string, friendId string) int {
	var fq models.FriendQuest
	var fq1 models.FriendQuest
	tx := utils.Db.Where("user_id = ? AND friend_id=?", mineId, friendId).Find(&fq)
	tx1 := utils.Db.Where("user_id = ? AND friend_id=?", friendId, mineId).Find(&fq1)
	if tx.RowsAffected == 0 && (tx1.RowsAffected == 0 || fq1.RequestStatus == "1") {
		if request := SendFriendWebSRequest(mineId, friendId); !request { //发送websocket失败，再存入数据库，防止申请丢失
			newF := models.FriendQuest{
				UserId:        mineId,
				FriendId:      friendId,
				RequestStatus: "1",
			}
			utils.Db.Create(&newF)
		}
		return utils.YesCode
	} else if fq.RequestStatus == "2" || fq1.RequestStatus == "2" { //已经同意
		return 2
	} else {
		return utils.NoCode //已经发送，未处理
	}
}
func AddFriendForSql(id string, friendId string) error {
	var fq, f models.Friend
	fq.FriendId = id
	f.UserId = id
	fq.UserId = friendId
	f.FriendId = friendId
	tx := utils.Db.Create(&f)
	utils.Db.Create(&fq)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
