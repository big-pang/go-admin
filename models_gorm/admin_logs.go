package models_gorm

import (
	"time"
)

// AdminLogs 后台操作日志
type AdminLogs struct {
	ID          int          `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"-"`
	AdminUserID int          `gorm:"column:admin_user_id;type:int(10) unsigned;not null;default:0" json:"adminUserId"`     // 用户id
	Name        string       `gorm:"column:name;type:varchar(30);not null;default:''" json:"name"`                         // 操作
	URL         string       `gorm:"column:url;type:varchar(100);not null;default:''" json:"url"`                          // URL
	LogMethod   string       `gorm:"column:log_method;type:varchar(8);not null;default:不记录" json:"logMethod"`              // 记录日志方法
	LogIP       string       `gorm:"column:log_ip;type:varchar(20);not null;default:''" json:"logIp"`                      // 操作IP
	CreatedAt   time.Time    `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createdAt"` // 操作时间
	UpdatedAt   time.Time    `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updatedAt"` // 操作时间
	Data        AdminLogData `gorm:"foreignKey:AdminLogID"`
	AdminUser   AdminUsers   `gorm:"foreignKey:AdminUserID"`
}

// TableName get sql table name.获取数据库表名
func (m *AdminLogs) TableName() string {
	return "admin_logs"
}

// AdminLogsColumns get sql column name.获取数据库列名
var AdminLogsColumns = struct {
	ID          string
	AdminUserID string
	Name        string
	URL         string
	LogMethod   string
	LogIP       string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	AdminUserID: "admin_user_id",
	Name:        "name",
	URL:         "url",
	LogMethod:   "log_method",
	LogIP:       "log_ip",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// SearchField 定义模型的可搜索字段
func (*AdminLogs) SearchField() []string {
	return []string{"name", "url", "log_ip"}
}

// WhereField 定义模型可作为条件的字段
func (*AdminLogs) WhereField() []string {
	return []string{"admin_user_id"}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminLogs) TimeField() []string {
	return []string{"created_at"}
}

// NoDeletionId 禁止删除的数据id
func (*AdminLogs) NoDeletionId() []int {
	return []int{}
}
