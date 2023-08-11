package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mt/models"
	"mt/routers"
	"mt/utils"
)

func main() {
	utils.Db = models.Init()
	err := models.InitMongoDB()
	if err != nil {
		fmt.Println("redis err:", err)
		return
	}
	re := gin.Default()
	routers.CreateRoute(re)
	_ = re.Run(":8080")
}
