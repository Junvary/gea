package v1

import (
	"gin-element-admin/global"
    "gin-element-admin/model"
    "gin-element-admin/model/request"
    "gin-element-admin/model/response"
    "gin-element-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags DevopsServer
// @Summary 创建DevopsServer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevopsServer true "创建DevopsServer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /server/createDevopsServer [post]
func CreateDevopsServer(c *gin.Context) {
	var server model.DevopsServer
	_ = c.ShouldBindJSON(&server)
	if err := service.CreateDevopsServer(server); err != nil {
        global.GeaLog.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DevopsServer
// @Summary 删除DevopsServer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevopsServer true "删除DevopsServer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /server/deleteDevopsServer [delete]
func DeleteDevopsServer(c *gin.Context) {
	var server model.DevopsServer
	_ = c.ShouldBindJSON(&server)
	if err := service.DeleteDevopsServer(server); err != nil {
        global.GeaLog.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DevopsServer
// @Summary 批量删除DevopsServer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DevopsServer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /server/deleteDevopsServerByIds [delete]
func DeleteDevopsServerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDevopsServerByIds(IDS); err != nil {
        global.GeaLog.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DevopsServer
// @Summary 更新DevopsServer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevopsServer true "更新DevopsServer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /server/updateDevopsServer [put]
func UpdateDevopsServer(c *gin.Context) {
	var server model.DevopsServer
	_ = c.ShouldBindJSON(&server)
	if err := service.UpdateDevopsServer(server); err != nil {
        global.GeaLog.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DevopsServer
// @Summary 用id查询DevopsServer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DevopsServer true "用id查询DevopsServer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /server/findDevopsServer [get]
func FindDevopsServer(c *gin.Context) {
	var server model.DevopsServer
	_ = c.ShouldBindQuery(&server)
	if err, reserver := service.GetDevopsServer(server.ID); err != nil {
        global.GeaLog.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reserver": reserver}, c)
	}
}

// @Tags DevopsServer
// @Summary 分页获取DevopsServer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DevopsServerSearch true "分页获取DevopsServer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /server/getDevopsServerList [get]
func GetDevopsServerList(c *gin.Context) {
	var pageInfo request.DevopsServerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDevopsServerInfoList(pageInfo); err != nil {
	    global.GeaLog.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
