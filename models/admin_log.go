package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// AdminLog struct
type AdminLog struct {
	Id          int       `orm:"column(id);auto;size(11);description(表ID)" json:"id""`
	AdminUserId int       `orm:"column(admin_user_id);size(10);default(0);description(用户id)" json:"admin_user_id"`
	Name        string    `orm:"column(name);size(30);description(操作)" json:"name"`
	Url         string    `orm:"column(url);size(100);description(url)" json:"url""`
	LogMethod   string    `orm:"column(log_method);size(8);default(不记录);description(记录日志方法)" json:"log_method"`
	LogIp       string    `orm:"column(log_ip);size(20);description(操作IP)" json:"log_ip"`
	CreatedAt   time.Time `orm:"column(created_at);type(timestamp);default(CURRENT_TIMESTAMP);description(创建时间)" json:"created_at"`
	UpdatedAt   time.Time `orm:"column(updated_at);type(timestamp);default(CURRENT_TIMESTAMP);description(更新时间)" json:"updated_at"`
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
