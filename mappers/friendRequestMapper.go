package mappers

import (
	"github.com/gin-gonic/gin"
	"mt/controllers"
)

func FriendRequestMapper(re *gin.Engine) {
	user := re.Group("/friendR")
	{
		user.POST("/getList", controllers.GetList)
		user.POST("/agree", controllers.Agree)
		user.POST("/friendSql", controllers.StoreFriend)
	}
}
