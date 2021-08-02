package models_gorm

import (
	"gorm.io/gorm"
)

//Content 将Content json字符串转换为结构体，便于html界面中range
type Content struct {
	Name    string
	Field   string
	Type    string
	Content string
	Option  string
	Form    string
}

// Setting 设置
type Setting struct {
	gorm.Model
	SettingGroupID uint32     `gorm:"column:setting_group_id;type:int(10) unsigned;not null;default:0" json:"settingGroupId"` // 所属分组
	Name           string     `gorm:"column:name;type:varchar(50);not null;default:''" json:"name"`                           // 名称
	Description    string     `gorm:"column:description;type:varchar(100);not null;default:''" json:"description"`            // 描述
	Code           string     `gorm:"column:code;type:varchar(50);not null;default:''" json:"code"`                           // 代码
	Content        string     `gorm:"column:content;type:text;not null" json:"content"`                                       // 设置配置及内容
	ContentStrut   []*Content `gorm:"-"`
	SortNumber     uint32     `gorm:"column:sort_number;type:int(10) unsigned;not null;default:1000" json:"sortNumber"` // 排序
}

// TableName get sql table name.获取数据库表名
func (m *Setting) TableName() string {
	return "setting"
}

// SettingColumns get sql column name.获取数据库列名
var SettingColumns = struct {
	ID             string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
	SettingGroupID string
	Name           string
	Description    string
	Code           string
	Content        string
	SortNumber     string
}{
	ID:             "id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	SettingGroupID: "setting_group_id",
	Name:           "name",
	Description:    "description",
	Code:           "code",
	Content:        "content",
	SortNumber:     "sort_number",
}

// SearchField 定义模型的可搜索字段
func (*Setting) SearchField() []string {
	return []string{"name", "description", "code"}
}

// NoDeletionId 禁止删除的数据id
func (*Setting) NoDeletionId() []int {
	return []int{1, 2, 3, 4, 5}
}

// WhereField 定义模型可作为条件的字段
func (*Setting) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*Setting) TimeField() []string {
	return []string{}
}
