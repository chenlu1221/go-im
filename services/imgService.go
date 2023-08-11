package services

import (
	"mt/models"
	"mt/utils"
)

func ImgService(id string) string {
	var u models.User
	utils.Db.Where("Id =?", id).Find(&u)
	if u.Avatar != "1" {
		return u.Avatar
	} else {
		return ""
	}
}
func LoadingImg(id string, name string) {
	utils.Db.Model(&models.User{}).Where("id=?", id).Update("avatar", name)
}
