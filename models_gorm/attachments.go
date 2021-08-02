package models_gorm

import (
	"gorm.io/gorm"
)

// Attachments 附件表
type Attachments struct {
	gorm.Model
	AdminUserID  int    `gorm:"column:admin_user_id;type:int(10) unsigned;not null;default:0" json:"adminUserId"` // 后台用户ID
	UserID       int    `gorm:"column:user_id;type:int(10) unsigned;not null;default:0" json:"userId"`            // 前台用户ID
	OriginalName string `gorm:"column:original_name;type:varchar(200);not null;default:''" json:"originalName"`   // 原文件名
	SaveName     string `gorm:"column:save_name;type:varchar(200);not null;default:''" json:"saveName"`           // 保存文件名
	SavePath     string `gorm:"column:save_path;type:varchar(255);not null;default:''" json:"savePath"`           // 系统完整路径
	URL          string `gorm:"column:url;type:varchar(255);not null;default:''" json:"url"`                      // 网站url路径
	Extension    string `gorm:"column:extension;type:varchar(100);not null;default:''" json:"extension"`          // 后缀
	Mime         string `gorm:"column:mime;type:varchar(100);not null;default:''" json:"mime"`                    // 类型
	Size         int64  `gorm:"column:size;type:bigint(20);not null;default:0" json:"size"`                       // 大小
	Md5          string `gorm:"column:md5;type:varchar(32);not null;default:''" json:"md5"`                       // MD5
	Sha1         string `gorm:"column:sha1;type:varchar(40);not null;default:''" json:"sha1"`                     // SHA1
}

// TableName get sql table name.获取数据库表名
func (m *Attachments) TableName() string {
	return "attachments"
}

// AttachmentsColumns get sql column name.获取数据库列名
var AttachmentsColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	AdminUserID  string
	UserID       string
	OriginalName string
	SaveName     string
	SavePath     string
	URL          string
	Extension    string
	Mime         string
	Size         string
	Md5          string
	Sha1         string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	AdminUserID:  "admin_user_id",
	UserID:       "user_id",
	OriginalName: "original_name",
	SaveName:     "save_name",
	SavePath:     "save_path",
	URL:          "url",
	Extension:    "extension",
	Mime:         "mime",
	Size:         "size",
	Md5:          "md5",
	Sha1:         "sha1",
}
