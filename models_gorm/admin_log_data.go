package models_gorm

import (
	"github.com/beego/beego/v2/client/orm"
)

// AdminLogData 后台操作日志数据
type AdminLogData struct {
	ID         int    `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"-"`
	AdminLogID int    `gorm:"column:admin_log_id;type:int(10) unsigned;not null;default:0" json:"adminLogId"` // 日志ID
	Data       string `gorm:"column:data;type:text;not null" json:"data"`                                     // 日志内容
}

// TableName get sql table name.获取数据库表名
func (m *AdminLogData) TableName() string {
	return "admin_log_data"
}

// AdminLogDataColumns get sql column name.获取数据库列名
var AdminLogDataColumns = struct {
	ID         string
	AdminLogID string
	Data       string
}{
	ID:         "id",
	AdminLogID: "admin_log_id",
	Data:       "data",
}

// SearchField 定义模型的可搜索字段
func (*AdminLogData) SearchField() []string {
	return []string{}
}

// WhereField 定义模型可作为条件的字段
func (*AdminLogData) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminLogData) TimeField() []string {
	return []string{}
}

// NoDeletionId 禁止删除的数据id
func (*AdminLogData) NoDeletionId() []int {
	return []int{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminLogData))
}
