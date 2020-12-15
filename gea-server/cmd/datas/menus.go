package datas

import (
	"gin-element-admin/global"
	"github.com/gookit/color"
	"os"
	"time"

	"gin-element-admin/model"
	"gorm.io/gorm"
)

var BaseMenus = []model.SysBaseMenu{
	// 仪表盘
	{GeaModel: global.GeaModel{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "dashboard/dashboard.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "gea gea-icon-home"}},
	// devops
	{GeaModel: global.GeaModel{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "devops", Name: "devops", Component: "devops/index.vue", Sort: 2, Meta: model.Meta{Title: "运维管理", Icon: "gea gea-icon-devops"}},
	{GeaModel: global.GeaModel{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "server", Name: "server", Component: "devops/server/server.vue", Sort: 1, Meta: model.Meta{Title: "远程主机管理", Icon: "gea gea-icon-server"}},
	{GeaModel: global.GeaModel{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "2", Path: "serverSSH/:id", Name: "serverSSH", Component: "devops/serverSSH/serverSSH.vue", Sort: 2, Meta: model.Meta{Title: "主机连接", Icon: "gea gea-icon-server"}},

	// 系统权限管理
	{GeaModel: global.GeaModel{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "authority", Name: "authority", Component: "authority/index.vue", Sort: 20, Meta: model.Meta{Title: "系统权限管理", Icon: "gea gea-icon-auth"}},
	{GeaModel: global.GeaModel{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "5", Path: "menus", Name: "menus", Component: "authority/menus/menus.vue", Sort: 1, Meta: model.Meta{Title: "菜单管理", Icon: "gea gea-icon-menu", KeepAlive: true}},
	{GeaModel: global.GeaModel{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "5", Path: "roles", Name: "roles", Component: "authority/roles/roles.vue", Sort: 2, Meta: model.Meta{Title: "角色管理", Icon: "gea gea-icon-role"}},
	{GeaModel: global.GeaModel{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "5", Path: "users", Name: "users", Component: "authority/users/users.vue", Sort: 3, Meta: model.Meta{Title: "用户管理", Icon: "gea gea-icon-team"}},
	{GeaModel: global.GeaModel{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "5", Path: "apis", Name: "apis", Component: "authority/apis/apis.vue", Sort: 4, Meta: model.Meta{Title: "API管理", Icon: "gea gea-icon-api", KeepAlive: true}},
	// 系统配置管理
	{GeaModel: global.GeaModel{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "manage", Name: "manage", Component: "manage/index.vue", Sort: 21, Meta: model.Meta{Title: "系统配置管理", Icon: "gea gea-icon-setting"}},
	{GeaModel: global.GeaModel{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "10", Path: "dict", Name: "dict", Component: "manage/dict/dict.vue", Sort: 1, Meta: model.Meta{Title: "字典管理", Icon: "gea gea-icon-dict"}},
	{GeaModel: global.GeaModel{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "10", Path: "dictDetail/:id", Name: "dictDetail", Component: "manage/dict/dictDetail.vue", Sort: 2, Meta: model.Meta{Title: "字典详情", Icon: "gea gea-icon-dict"}},
	{GeaModel: global.GeaModel{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "10", Path: "config", Name: "config", Component: "manage/config/config.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "gea gea-icon-setting"}},
	// 系统工具
	{GeaModel: global.GeaModel{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "tools", Name: "tools", Component: "tools/index.vue", Sort: 22, Meta: model.Meta{Title: "系统工具", Icon: "gea gea-icon-tool"}},
	{GeaModel: global.GeaModel{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "codeGenerator", Name: "codeGenerator", Component: "tools/codeGenerator/codeGenerator.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "gea gea-icon-code", KeepAlive: true}},
	{GeaModel: global.GeaModel{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "tools/formCreate/formCreate.vue", Sort: 2, Meta: model.Meta{Title: "表单生成器", Icon: "gea gea-icon-form", KeepAlive: true}},
	{GeaModel: global.GeaModel{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "backServer", Name: "backServer", Component: "tools/server/server.vue", Sort: 3, Meta: model.Meta{Title: "后端服务器", Icon: "gea gea-icon-back-server"}},
	// 系统日志管理
	{GeaModel: global.GeaModel{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "logs", Name: "logs", Component: "logs/index.vue", Sort: 23, Meta: model.Meta{Title: "系统日志管理", Icon: "gea gea-icon-log"}},
	{GeaModel: global.GeaModel{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "18", Path: "operationLog", Name: "operationLog", Component: "logs/operation/operation.vue", Sort: 1, Meta: model.Meta{Title: "操作日志", Icon: "gea gea-icon-operation"}},
	// 非目录菜单
	{GeaModel: global.GeaModel{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "userInfo", Name: "userInfo", Component: "user/userInfo.vue", Sort: 24, Meta: model.Meta{Title: "个人信息", Icon: "gea gea-icon-user"}},

	// 系统示例
	{GeaModel: global.GeaModel{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "examples", Name: "example", Component: "examples/index.vue", Sort: 25, Meta: model.Meta{Title: "系统示例", Icon: "gea gea-icon-exam"}},
	{GeaModel: global.GeaModel{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "21", Path: "iconList", Name: "iconList", Component: "examples/iconList/iconList.vue", Sort: 1, Meta: model.Meta{Title: "图标合集", Icon: "gea gea-icon-icon"}},

	// 系统介绍
	{GeaModel: global.GeaModel{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "about/index.vue", Sort: 26, Meta: model.Meta{Title: "系统介绍", Icon: "gea gea-icon-about"}},
	{GeaModel: global.GeaModel{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "23", Path: "we", Name: "we", Component: "about/we/we.vue", Sort: 1, Meta: model.Meta{Title: "关于我们", Icon: "gea gea-icon-we"}},
	{GeaModel: global.GeaModel{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "23", Path: "swagger", Name: "swagger", Component: "about/swagger/swagger.vue", Sort: 2, Meta: model.Meta{Title: "API文档", Icon: "gea gea-icon-api"}},

}

func InitSysBaseMenus(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 27}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("sys_base_menus表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&BaseMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_base_menus 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
