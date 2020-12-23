package router

import "github.com/gin-gonic/gin"
import "gin-element-admin/api/v1"

func InitWebsocketRouter(Router *gin.RouterGroup) {
	WebsocketRouter := Router.Group("ws")
	{
		WebsocketRouter.GET("webSSH", v1.CreateWebSSH)
	}
}
