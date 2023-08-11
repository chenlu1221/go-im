package mappers

import (
	"github.com/gin-gonic/gin"
	"mt/controllers"
)

func MessageMapper(re *gin.Engine) {
	user := re.Group("/webSocket")
	{
		user.GET("/login", controllers.WebSocketLogin)
	}
}
