package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mt/services"
	"strings"
)

func GetHeadImg(c *gin.Context) {

	if n := services.ImgService(c.PostForm("id")); n != "" {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", n))
		c.File("C:/Users/Administrator/Desktop/go/img/" + n)
	} else {
		c.JSON(200, gin.H{
			"msg": 1,
		})
	}
}

// LoadHeadImg 修改头像
func LoadHeadImg(c *gin.Context) {
	id := c.PostForm("id")
	img := c.PostForm("data")
	data := img[strings.IndexByte(img, ',')+1:]
	imageBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": 1,
			"err": err,
		})
		fmt.Println("解码错误:", err)
		return
	}
	dst := fmt.Sprintf("C:/Users/Administrator/Desktop/go/img/%s", id+".png")
	err = ioutil.WriteFile(dst, imageBytes, 0644)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": 1,
			"err": err,
		})
		fmt.Println("保存文件错误:", err)
		return
	}
	services.LoadingImg(id, id+".png")
	c.JSON(200, gin.H{
		"msg": 0,
	})
}
