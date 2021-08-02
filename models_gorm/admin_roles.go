package models_gorm

// AdminRoles 后台角色
type AdminRoles struct {
	ID          int    `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"-"`
	Name        string `gorm:"column:name;type:varchar(50);not null;default:''" json:"name"`                // 名称
	Description string `gorm:"column:description;type:varchar(100);not null;default:''" json:"description"` // 简介
	URL         string `gorm:"column:url;type:varchar(1000);not null;default:''" json:"url"`                // 权限
	Status      int8   `gorm:"column:status;type:tinyint(4);not null;default:1" json:"status"`              // 是否启用
}

// TableName get sql table name.获取数据库表名
func (m *AdminRoles) TableName() string {
	return "admin_roles"
}

// AdminRolesColumns get sql column name.获取数据库列名
var AdminRolesColumns = struct {
	ID          string
	Name        string
	Description string
	URL         string
	Status      string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	URL:         "url",
	Status:      "status",
}

// SearchField 定义模型的可搜索字段
func (*AdminRoles) SearchField() []string {
	return []string{"name", "description"}
}

// NoDeletionId 禁止删除的数据id
func (*AdminRoles) NoDeletionId() []int {
	return []int{1}
}

// WhereField 定义模型可作为条件的字段
func (*AdminRoles) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminRoles) TimeField() []string {
	return []string{}
}
