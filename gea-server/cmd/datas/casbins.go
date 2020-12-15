package datas

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

var Carbines = []gormadapter.CasbinRule{
	// 无须鉴权
	{PType: "p", V0: "888", V1: "/base/login", V2: "POST"},

	// API组
	{PType: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},

	// 角色管理
	{PType: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},

	// 菜单管理
	{PType: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},

	// Casbin组
	{PType: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/casbinTest/:pathParam", V2: "GET"},

	// JWT组
	{PType: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},

	// 系统
	{PType: "p", V0: "888", V1: "/system/getSystemConfig", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/setSystemConfig", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/getServerInfo", V2: "POST"},

	// 代码生成器
	{PType: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
	{PType: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
	{PType: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
	{PType: "p", V0: "888", V1: "/autoCode/getColumn", V2: "GET"},

	// 用户管理
	{PType: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/user/getUserinfo", V2: "GET"},

	// 字典管理
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysDictionary/createSysDictionary", V2: "POST"},
	{PType: "p", V0: "888", V1: "/sysDictionary/deleteSysDictionary", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/sysDictionary/updateSysDictionary", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/sysDictionary/findSysDictionary", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysDictionary/getSysDictionaryList", V2: "GET"},

	// 日志管理
	{PType: "p", V0: "888", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

	//	devops
	{PType: "p", V0: "888", V1: "/devops/createDevopsServer", V2: "POST"},
	{PType: "p", V0: "888", V1: "/devops/deleteDevopsServer", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/devops/deleteDevopsServerByIds", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/devops/updateDevopsServer", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/devops/findDevopsServer", V2: "GET"},
	{PType: "p", V0: "888", V1: "/devops/getDevopsServerList", V2: "GET"},
}

func InitCasbinModel(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("p_type = ? AND v0 IN ?", "p", []string{"888", "8881", "9528"}).Find(&[]gormadapter.CasbinRule{}).RowsAffected == 53 {
			color.Danger.Println("casbin_rule表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&Carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> casbin_rule 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
