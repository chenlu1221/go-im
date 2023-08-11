package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mt/mappers"
	"mt/utils"
)

func CreateRoute(ctx *gin.Engine) *gin.Engine {
	ctx.Use(utils.Login) //拦截器
	ctx.Use(cors.Default())
	mappers.UserMapper(ctx)
	mappers.ImgMapper(ctx)
	mappers.FriendMapper(ctx)
	mappers.MessageMapper(ctx)
	mappers.FriendRequestMapper(ctx)
	return ctx
}
