package mappers

import (
	"github.com/gin-gonic/gin"
	"mt/controllers"
)

func ImgMapper(re *gin.Engine) {
	user := re.Group("/img")
	{
		user.POST("/getHeadImg", controllers.GetHeadImg)
		user.POST("/unloadImg", controllers.LoadHeadImg)
	}
}
