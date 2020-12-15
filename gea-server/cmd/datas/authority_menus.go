package datas

import (
	"github.com/gookit/color"
	"gorm.io/gorm"
	"os"
)

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId           uint
}

var AuthorityMenus = []SysAuthorityMenus{
	{"888", 1},
	{"888", 2},
	{"888", 3},
	{"888", 4},
	{"888", 5},
	{"888", 6},
	{"888", 7},
	{"888", 8},
	{"888", 9},
	{"888", 10},
	{"888", 11},
	{"888", 12},
	{"888", 13},
	{"888", 14},
	{"888", 15},
	{"888", 16},
	{"888", 17},
	{"888", 18},
	{"888", 19},
	{"888", 20},
	{"888", 21},
	{"888", 22},
	{"888", 23},
	{"888", 24},
	{"888", 25},
	{"8881", 1},
	{"9528", 1},
}

func InitSysAuthorityMenus(db *gorm.DB) {
	if err := db.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ?", []string{"888", "8881", "9528"}).Find(&[]SysAuthorityMenus{}).RowsAffected == 53 {
			color.Danger.Println("sys_authority_menus表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&AuthorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_authority_menus 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
