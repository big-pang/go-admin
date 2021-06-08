package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// UserLevel struct
type UserLevel struct {
	Id          int        `orm:"column(id);auto;size(11);description(表ID)" json:"id"`
	Name        string     `orm:"column(name);size(20);description(名称)" json:"name"`
	Description string     `orm:"column(description);size(50);description(简介)" json:"description"`
	Img         string     `orm:"column(img);size(255);description(图片)" json:"img"`
	Status      int8       `orm:"column(status);size(1);default(1);description(是否启用 0：否 1：是)" json:"status"`
	CreatedAt   time.Time  `orm:"column(created_at);type(timestamp);default(CURRENT_TIMESTAMP);description(创建时间)" json:"created_at"`
	UpdatedAt   time.Time  `orm:"column(updated_at);type(timestamp);default(CURRENT_TIMESTAMP);description(更新时间)" json:"updated_at"`
	DeletedAt   *time.Time `orm:"column(deleted_at);type(timestamp);default(NULL);description(删除时间)" json:"deleted_at"`
}

// TableName 自定义table 名称
func (*UserLevel) TableName() string {
	return "user_level"
}

// SearchField 定义模型的可搜索字段
func (*UserLevel) SearchField() []string {
	return []string{"name", "description"}
}

// NoDeletionId 禁止删除的数据id
func (*UserLevel) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*UserLevel) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*UserLevel) TimeField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(UserLevel))
}
