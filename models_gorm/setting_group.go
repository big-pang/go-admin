package models_gorm

import (
	"gorm.io/gorm"
)

// SettingGroup 设置分组
type SettingGroup struct {
	gorm.Model
	Module         string `gorm:"column:module;type:varchar(30);not null;default:''" json:"module"`                 // 作用模块
	Name           string `gorm:"column:name;type:varchar(50);not null;default:''" json:"name"`                     // 名称
	Description    string `gorm:"column:description;type:varchar(100);not null;default:''" json:"description"`      // 描述
	Code           string `gorm:"column:code;type:varchar(50);not null;default:''" json:"code"`                     // 代码
	SortNumber     uint32 `gorm:"column:sort_number;type:int(10) unsigned;not null;default:1000" json:"sortNumber"` // 排序
	AutoCreateMenu int8   `gorm:"column:auto_create_menu;type:tinyint(4);not null;default:0" json:"autoCreateMenu"` // 自动生成菜单
	AutoCreateFile int8   `gorm:"column:auto_create_file;type:tinyint(4);not null;default:0" json:"autoCreateFile"` // 自动生成配置文件
	Icon           string `gorm:"column:icon;type:varchar(30);not null;default:fa-list" json:"icon"`                // 图标
}

// TableName get sql table name.获取数据库表名
func (m *SettingGroup) TableName() string {
	return "setting_group"
}

// SettingGroupColumns get sql column name.获取数据库列名
var SettingGroupColumns = struct {
	ID             string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
	Module         string
	Name           string
	Description    string
	Code           string
	SortNumber     string
	AutoCreateMenu string
	AutoCreateFile string
	Icon           string
}{
	ID:             "id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	Module:         "module",
	Name:           "name",
	Description:    "description",
	Code:           "code",
	SortNumber:     "sort_number",
	AutoCreateMenu: "auto_create_menu",
	AutoCreateFile: "auto_create_file",
	Icon:           "icon",
}
