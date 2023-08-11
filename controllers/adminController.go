package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"mt/models"
	"mt/services"
	"mt/utils"
	"time"
)

func Login(c *gin.Context) {
	println(c.ClientIP())
	b, a := services.LoginService(c.PostForm("tel"), c.PostForm("password"), c.PostForm("msg"))
	if a == 1 {
		c.JSON(200, gin.H{
			"code": 0,
			"data": b,
		})
	}
	if a == 2 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "密码错误",
		})
	}
	if a == 3 {
		c.JSON(200, gin.H{
			"code": 2,
			"msg":  "用户不存在",
		})
	}
}

// Register 注册
func Register(c *gin.Context) {
	var usr models.User
	Passwd1 := c.PostForm("password")
	usr.Mobile = c.PostForm("tel")
	usr.Email = c.PostForm("email")
	usr.Avatar = "1"
	usr.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	usr.Nickname = "用户"
	usr.Passwd = utils.MakePasswd(Passwd1, usr.Salt)
	usr.Create = time.Now()
	usr.Tocken = fmt.Sprintf("%08d", rand.Int31())
	a := services.RegisterService(usr)
	if a == 0 {
		c.JSON(200, gin.H{
			"code": 0,
		})
	} else if a == 1 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "该手机号已注册",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2,
			"msg":  "该邮箱已注册",
		})
	}
}

// EmailVer 邮箱验证
func EmailVer(ctx *gin.Context) {
	em := ctx.PostForm("email")
	if eTime, err := utils.EmailCode[em]; err == true {
		if time.Now().Sub(eTime).Minutes() < 1 {
			ctx.JSON(200, gin.H{
				"code": 1,
				"err":  "1分钟内请不要重复获取",
			})
			return
		}
	}
	validate, err := utils.SendEmailValidate([]string{em})
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"err":  "无效邮箱账号",
		})
		fmt.Println("email err:", err)
		return
	}
	utils.EmailCode[em] = time.Now()
	ctx.JSON(200, gin.H{
		"code":    0,
		"reEmail": validate,
	})
}

// ChangeU 修改密码
func ChangeU(c *gin.Context) {
	em := c.PostForm("email")
	Passwd1 := c.PostForm("password")
	salt := fmt.Sprintf("%06d", rand.Int31n(10000))
	if a := services.ChangeService(em, utils.MakePasswd(Passwd1, salt)); a == 0 {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改成功",
		})
	} else if a == 1 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "用户不存在",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2,
			"msg":  "邮箱账户过多",
		})
	}
}

// ChangeN 修改昵称
func ChangeN(c *gin.Context) {
	nk := c.PostForm("nickname")
	id := c.PostForm("id")
	services.ChangeNService(nk, id)
	c.JSON(200, gin.H{
		"msg": 0,
	})
}

// FindUser 查找用户
func FindUser(c *gin.Context) {
	id := c.PostForm("id")
	name, avatar, _ := services.GetUser(id)
	if name != "" {
		if avatar != "1" {
			data, err := ioutil.ReadFile("C:/Users/Administrator/Desktop/go/img/" + avatar)
			base64Str := base64.StdEncoding.EncodeToString(data)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(200, gin.H{
				"msg":  utils.YesCode,
				"data": models.FJson{Nickname: name, Avatar: "data:image/jpeg;base64," + base64Str, Id: id},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  utils.YesCode,
				"data": models.FJson{Nickname: name, Avatar: "1", Id: id},
			})
		}
	} else {
		c.JSON(200, gin.H{
			"msg": utils.NoCode,
			"err": "用户不存在",
		})
	}
}
func ExitUser(c *gin.Context) {
	id := c.PostForm("id")
	services.Exit(id)
	c.JSON(200, nil)
}
