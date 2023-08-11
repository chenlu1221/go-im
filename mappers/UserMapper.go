package mappers

import (
	"github.com/gin-gonic/gin"
	"mt/controllers"
)

func UserMapper(re *gin.Engine) {
	user := re.Group("/user")
	{
		user.POST("/login", controllers.Login)
		user.POST("/loginE", controllers.Login)
		user.POST("/register", controllers.Register)
		user.POST("/emailVer", controllers.EmailVer)
		user.POST("/changeU", controllers.ChangeU)
		user.POST("/changeN", controllers.ChangeN)
		user.POST("/findUser", controllers.FindUser)
		user.POST("/exit", controllers.ExitUser)
	}
}
