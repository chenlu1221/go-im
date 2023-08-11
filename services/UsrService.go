package services

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"mt/models"
	"mt/utils"
	"strconv"
	"time"
)

// LoginService 登录
func LoginService(mod string, pas string, msg string) (models.User, int) {
	var u models.User
	var tx *gorm.DB
	if msg == "1" {
		tx = utils.Db.Where("mobile = ?",
			mod).Find(&u)
	} else {
		tx = utils.Db.Where("email = ?",
			mod).Find(&u)
	}
	if tx.RowsAffected != 0 {
		if utils.ValidatePasswd(pas, u.Salt, u.Passwd) {
			str := fmt.Sprintf("%d", time.Now().Unix())
			token := utils.MD5Encode(str)
			u.Tocken = token
			utils.Db.Model(&u).Where("id=?", u.Id).Update("tocken", u.Tocken)
			//登录成功，将用户加入列表
			models.UserMap[u.Id] = u
			return u, 1 //用户存在，密码正确
		} else {
			return u, 2 //用户存在，密码错误
		}
	} else {
		return u, 3 //用户不存在
	}

}

// RegisterService 注册
func RegisterService(usr models.User) int {
	var u, e models.User
	tx := utils.Db.Where("mobile = ?",
		usr.Mobile).Find(&u)
	te := utils.Db.Where("email = ?",
		usr.Email).Find(&e)
	if tx.RowsAffected == 0 && te.RowsAffected == 0 {
		utils.Db.Create(&usr)
		utils.Db.Where("mobile =?", usr.Mobile).Find(&u)
		utils.Db.Model(&u).Update("nickname", u.Nickname+strconv.FormatInt(u.Id, 10))
		return 0
	} else if tx.RowsAffected != 0 { //手机号已经注册
		return 1
	} else { //邮箱被注册
		return 2
	}
}

// ChangeService 修改密码
func ChangeService(ema string, pswd string) int {
	var e, u models.User
	te := utils.Db.Where("email = ?",
		ema).Find(&e)
	if te.RowsAffected == 0 {
		return 1
	} else if te.RowsAffected == 1 {
		utils.Db.Where("email =?", ema).Find(&u)
		utils.Db.Model(&u).Update("passwd", pswd)
		return 0
	} else {
		return 2
	}
}

// ChangeNService 修改昵称
func ChangeNService(nickname string, id string) {
	var e models.User
	utils.Db.Model(&e).Where("id=?", id).Update("nickname", nickname)
}

func GetUser(id string) (string, string, string) {
	var user models.User
	tx := utils.Db.Where("id=?", id).Find(&user)
	if tx.RowsAffected != 0 {
		return user.Nickname, user.Avatar, strconv.FormatInt(user.Id, 10)
	} else {
	}
	return "", "", ""
}

// Exit 退出登录
func Exit(id string) {
	fmt.Println("链接" + id + "关闭")
	if err := models.FWebRequestMap[id].WriteMessage(websocket.PingMessage, []byte{}); err != nil {
		//发送心跳检测失败，关闭链接
		models.FWebRequestMap[id].Close()
		return
	}
	delete(models.FWebRequestMap, id) //删除发送好友请求的websocket链接
	iid, _ := strconv.ParseInt(id, 10, 64)
	delete(models.UserMap, iid) //删除id对应的用户列表
	models.CreateChanMap(id)    //删除对应的消息通道
}
