package router

import (
	"gin-element-admin/api/v1"
	"gin-element-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDevopsRouter(Router *gin.RouterGroup) {
	DevopsRouter := Router.Group("devops").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DevopsRouter.POST("createDevopsServer", v1.CreateDevopsServer)             // 新增远程主机
		DevopsRouter.DELETE("deleteDevopsServer", v1.DeleteDevopsServer)           // 删除远程主机
		DevopsRouter.DELETE("deleteDevopsServerByIds", v1.DeleteDevopsServerByIds) // 批量删除远程主机
		DevopsRouter.PUT("updateDevopsServer", v1.UpdateDevopsServer)              // 更新远程主机
		DevopsRouter.GET("findDevopsServer", v1.FindDevopsServer)                  // 根据ID获取远程主机
		DevopsRouter.GET("getDevopsServerList", v1.GetDevopsServerList)            // 获取远程主机列表
	}
}
