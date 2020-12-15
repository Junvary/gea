package service

import (
	"gin-element-admin/global"
	"gin-element-admin/model"
	"gin-element-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDevopsServer
//@description: 创建DevopsServer记录
//@param: server model.DevopsServer
//@return: err error

func CreateDevopsServer(server model.DevopsServer) (err error) {
	err = global.GeaDb.Create(&server).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDevopsServer
//@description: 删除DevopsServer记录
//@param: server model.DevopsServer
//@return: err error

func DeleteDevopsServer(server model.DevopsServer) (err error) {
	err = global.GeaDb.Delete(server).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDevopsServerByIds
//@description: 批量删除DevopsServer记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDevopsServerByIds(ids request.IdsReq) (err error) {
	err = global.GeaDb.Delete(&[]model.DevopsServer{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDevopsServer
//@description: 更新DevopsServer记录
//@param: server *model.DevopsServer
//@return: err error

func UpdateDevopsServer(server model.DevopsServer) (err error) {
	err = global.GeaDb.Save(&server).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDevopsServer
//@description: 根据id获取DevopsServer记录
//@param: id uint
//@return: err error, server model.DevopsServer

func GetDevopsServer(id uint) (err error, server model.DevopsServer) {
	err = global.GeaDb.Where("id = ?", id).First(&server).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDevopsServerInfoList
//@description: 分页获取DevopsServer记录
//@param: info request.DevopsServerSearch
//@return: err error, list interface{}, total int64

func GetDevopsServerInfoList(info request.DevopsServerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GeaDb.Model(&model.DevopsServer{})
    var servers []model.DevopsServer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Area != "" {
        db = db.Where("`area` LIKE ?","%"+ info.Area+"%")
    }
    if info.Location != "" {
        db = db.Where("`location` LIKE ?","%"+ info.Location+"%")
    }
    if info.Category != "" {
        db = db.Where("`category` LIKE ?","%"+ info.Category+"%")
    }
    if info.Os != 0 {
        db = db.Where("`os` = ?",info.Os)
    }
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Ip != "" {
        db = db.Where("`ip` LIKE ?","%"+ info.Ip+"%")
    }
    if info.Port != "" {
        db = db.Where("`port` LIKE ?","%"+ info.Port+"%")
    }
    if info.Remark != "" {
        db = db.Where("`remark` LIKE ?","%"+ info.Remark+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&servers).Error
	return err, servers, total
}