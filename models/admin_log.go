package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// AdminLog struct
type AdminLog struct {
	Id          int    `orm:"column(id);auto;size(11)" description:"表ID"`
	AdminUserId int    `orm:"column(admin_user_id);size(10);default(0)" description:"用户id"`
	Name        string `orm:"column(name);size(30)" description:"操作"`
	Url         string `orm:"column(url);size(100)" description:"url"`
	LogMethod   string `orm:"column(log_method);size(8);default(不记录)" description:"记录日志方法"`
	LogIp       string `orm:"column(log_ip);size(20)" description:"操作IP"`
	CreatedAt   string `orm:"column(created_at);default(CURRENT_TIMESTAMP)" description:"操作时间"`
	UpdatedAt   string `orm:"column(updated_at);default(CURRENT_TIMESTAMP)" description:"更新时间"`
}

// SearchField 定义模型的可搜索字段
func (*AdminLog) SearchField() []string {
	return []string{"name", "url", "log_ip"}
}

// WhereField 定义模型可作为条件的字段
func (*AdminLog) WhereField() []string {
	return []string{"admin_user_id"}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminLog) TimeField() []string {
	return []string{"created_at"}
}

// NoDeletionId 禁止删除的数据id
func (*AdminLog) NoDeletionId() []int {
	return []int{}
}

// TableName 自定义table 名称
func (*AdminLog) TableName() string {
	return "admin_log"
}

// init 在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminLog))
}
