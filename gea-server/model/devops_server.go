package model

import (
	"gin-element-admin/global"
)

// 如果含有time.Time 请自行import time包
type DevopsServer struct {
      global.GeaModel
      Area  string `json:"area" form:"area" gorm:"column:area;comment:区域;type:varchar(40);size:40;"`
      Location  string `json:"location" form:"location" gorm:"column:location;comment:位置;type:varchar(40);size:40;"`
      Category  string `json:"category" form:"category" gorm:"column:category;comment:类别;type:varchar(40);size:40;"`
      Os  int `json:"os" form:"os" gorm:"column:os;comment:操作系统;type:int;size:3;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(30);size:30;"`
      Ip  string `json:"ip" form:"ip" gorm:"column:ip;comment:IP地址;type:varchar(30);size:30;"`
      Port  string `json:"port" form:"port" gorm:"column:port;comment:端口;type:varchar(8);size:8;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:描述备注;type:varchar(200);size:200;"`
}


func (DevopsServer) TableName() string {
  return "devops_server"
}
