package mappers

import (
	"github.com/gin-gonic/gin"
	"mt/controllers"
)

func FriendMapper(re *gin.Engine) {
	user := re.Group("/friend")
	{
		user.POST("/getFriend", controllers.GetFriendById)
		user.POST("/addFriend", controllers.AddFriendById)
	}
}
