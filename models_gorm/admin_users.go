package models_gorm

import (
	"crypto/sha1"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"go-admin/initialize/mysql"
	"strings"
	"time"
)

// AdminUsers 后台用户
type AdminUsers struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"-"`
	Username  string    `gorm:"column:username;type:varchar(30);not null;default:''" json:"username"`                           // 用户名
	Password  string    `gorm:"column:password;type:varchar(255);not null;default:''" json:"password"`                          // 密码
	Nickname  string    `gorm:"column:nickname;type:varchar(30);not null;default:''" json:"nickname"`                           // 昵称
	Avatar    string    `gorm:"column:avatar;type:varchar(255);not null;default:/static/admin/images/avatar.png" json:"avatar"` // 头像
	Role      string    `gorm:"column:role;type:varchar(200);not null;default:''" json:"role"`                                  // 角色
	Status    int8      `gorm:"column:status;type:tinyint(4);not null;default:1" json:"status"`                                 // 是否启用 0：否 1：是
	DeletedAt time.Time `gorm:"column:deleted_at;type:timestamp" json:"deletedAt"`                                              // 删除时间
}

// TableName get sql table name.获取数据库表名
func (user *AdminUsers) TableName() string {
	return "admin_users"
}

// AdminUsersColumns get sql column name.获取数据库列名
var AdminUsersColumns = struct {
	ID        string
	Username  string
	Password  string
	Nickname  string
	Avatar    string
	Role      string
	Status    string
	DeletedAt string
}{
	ID:        "id",
	Username:  "username",
	Password:  "password",
	Nickname:  "nickname",
	Avatar:    "avatar",
	Role:      "role",
	Status:    "status",
	DeletedAt: "deleted_at",
}

// SearchField 定义模型的可搜索字段
func (*AdminUsers) SearchField() []string {
	return []string{"nickname", "username"}
}

// NoDeletionID 禁止删除的数据id
func (*AdminUsers) NoDeletionID() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*AdminUsers) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminUsers) TimeField() []string {
	return []string{}
}

// GetSignStrByAdminUser 获取加密字符串，用在登录的时候加密处理
func (user *AdminUsers) GetSignStrByAdminUser(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s%s", user.ID, user.Username, ua))))
}

// GetAuthUrl 获取已授权url
func (user *AdminUsers) GetAuthUrl() map[string]interface{} {
	var (
		urlArr ParamsList
	)
	authURL := make(map[string]interface{})

	err := mysql.DB.Model(AdminRoles{}).
		Where("id in ?", strings.Split(user.Role, ",")).Where("status =?", 1).
		Pluck("url", &urlArr).Error
	if err != nil {
		logs.Error(err)
	}

	urlIDStr := ""
	for k, row := range urlArr {
		urlStr, ok := row.(string)
		if ok {
			if k == 0 {
				urlIDStr = urlStr
			} else {
				urlIDStr += "," + urlStr
			}
		}
	}
	urlIDArr := strings.Split(urlIDStr, ",")

	var authURLArr ParamsList

	if len(urlIDStr) > 0 {

		mysql.DB.Model(AdminMenus{}).Where("id in ?", urlIDArr).Pluck("url", &authURLArr)

		for k, row := range authURLArr {
			val, ok := row.(string)
			if ok {
				authURL[val] = k
			}
		}

	}
	return authURL

}

// GetShowMenu 获取当前用户已授权的显示菜单
func (user *AdminUsers) GetShowMenu() map[int]Params {
	//var maps []Params
	var menus []AdminMenus
	returnMaps := make(map[int]Params)

	if user.ID == 1 {
		err := mysql.DB.Model(AdminMenus{}).Select([]string{"id", "parent_id", "name", "url", "icon", "sort_id"}).Where("is_show = ?", 1).Order("sort_id, id").Find(&menus).Error
		if err == nil {
			for _, m := range menus {
				returnMaps[m.ID] = Params{
					"Id":       m.ID,
					"ParentId": m.ParentID,
					"Name":     m.Name,
					"Url":      m.URL,
					"Icon":     m.Icon,
					"SortId":   m.SortID,
				}
			}
			return returnMaps
		}
		return map[int]Params{}
	}

	var list ParamsList
	err := mysql.DB.Model(AdminRoles{}).Where("id in ?", strings.Split(user.Role, ",")).Where("status = ?", 1).Pluck("url", &list).Error

	if err == nil {
		var urlIDArr []string
		for _, m := range list {
			urlIDArr = append(urlIDArr, strings.Split(m.(string), ",")...)
		}
		err := mysql.DB.Model(new(AdminMenus)).Select("id", "parent_id", "name", "url", "icon", "sort_id").Where("id in ?", urlIDArr).Where("is_show = ?", 1).Order("sort_id, id").Find(&menus).Error
		if err == nil {
			for _, m := range menus {
				returnMaps[m.ID] = Params{
					"Id":       m.ID,
					"ParentId": m.ParentID,
					"Name":     m.Name,
					"Url":      m.URL,
					"Icon":     m.Icon,
					"SortId":   m.SortID,
				}
			}
			return returnMaps
		}
		return map[int]Params{}
	}
	return returnMaps

}

// GetRoleText 用户角色名称
func (user *AdminUsers) GetRoleText() map[int]*AdminRoles {
	roleIDArr := strings.Split(user.Role, ",")
	var adminRoles []*AdminRoles
	mysql.DB.Model(AdminRoles{}).Where("id in ?", roleIDArr).Select("id", "name").Find(&adminRoles)

	adminRoleMap := make(map[int]*AdminRoles)
	for _, v := range adminRoles {
		adminRoleMap[v.ID] = v
	}
	return adminRoleMap
}

// GetAdminUser 获取所有用户
func (*AdminUsers) GetAdminUser() []*AdminUsers {
	var adminUsers []*AdminUsers
	mysql.DB.Model(AdminMenus{}).Find(&adminUsers)
	return adminUsers
}
