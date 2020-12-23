package initialize

import (
	_ "gin-element-admin/docs"
	"gin-element-admin/global"
	"gin-element-admin/middleware"
	"gin-element-admin/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GeaConfig.Local.Path, http.Dir(global.GeaConfig.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GeaLog.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GeaLog.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GeaLog.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		router.InitWebsocketRouter(PublicGroup)           // websocket组
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitApiRouter(PrivateGroup)                 // 注册功能api路由
		router.InitJwtRouter(PrivateGroup)                 // jwt相关路由
		router.InitUserRouter(PrivateGroup)                // 注册用户路由
		router.InitMenuRouter(PrivateGroup)                // 注册menu路由
		router.InitSystemRouter(PrivateGroup)              // system相关路由
		router.InitCasbinRouter(PrivateGroup)              // 权限相关路由
		router.InitAutoCodeRouter(PrivateGroup)            // 创建自动化代码
		router.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
		router.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
		router.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
		router.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
		router.InitDevopsRouter(PrivateGroup)              // 运维组

	}
	global.GeaLog.Info("router register success")
	return Router
}
