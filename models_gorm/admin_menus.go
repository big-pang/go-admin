package models_gorm

// AdminMenus 后台菜单
type AdminMenus struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"-"`
	ParentID  int    `gorm:"column:parent_id;type:int(10) unsigned;not null;default:0" json:"parentId"`   // 父级菜单
	Name      string `gorm:"column:name;type:varchar(30);not null;default:''" json:"name"`                // 名称
	URL       string `gorm:"index:index_url;column:url;type:varchar(100);not null;default:''" json:"url"` // url
	RouteName string `gorm:"column:route_name;type:varchar(100);not null;default:''" json:"routeName"`    // 路由名或者路由标识
	Icon      string `gorm:"column:icon;type:varchar(30);not null;default:fa-list" json:"icon"`           // 图标
	IsShow    int8   `gorm:"column:is_show;type:tinyint(4);not null;default:1" json:"isShow"`             // 等级
	SortID    int    `gorm:"column:sort_id;type:int(10) unsigned;not null;default:1000" json:"sortId"`    // 排序
	LogMethod string `gorm:"column:log_method;type:varchar(8);not null;default:不记录" json:"logMethod"`     // 记录日志方法
}

// TableName get sql table name.获取数据库表名
func (m *AdminMenus) TableName() string {
	return "admin_menus"
}

// AdminMenusColumns get sql column name.获取数据库列名
var AdminMenusColumns = struct {
	ID        string
	ParentID  string
	Name      string
	URL       string
	RouteName string
	Icon      string
	IsShow    string
	SortID    string
	LogMethod string
}{
	ID:        "id",
	ParentID:  "parent_id",
	Name:      "name",
	URL:       "url",
	RouteName: "route_name",
	Icon:      "icon",
	IsShow:    "is_show",
	SortID:    "sort_id",
	LogMethod: "log_method",
}

// GetLogMethod 自定义方法
func (*AdminMenus) GetLogMethod() []string {
	return []string{"不记录", "GET", "POST", "PUT", "DELETE"}
}

// SearchField 定义模型的可搜索字段
func (*AdminMenus) SearchField() []string {
	return []string{}
}

// WhereField 定义模型可作为条件的字段
func (*AdminMenus) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminMenus) TimeField() []string {
	return []string{}
}

// NoDeletionId 不允许删除的id
func (*AdminMenus) NoDeletionId() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
}
