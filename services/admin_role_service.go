package services

import (
	"github.com/beego/beego/v2/client/orm"
	"go-admin/formvalidate"
	"go-admin/initialize/mysql"
	"go-admin/models_gorm"
	"go-admin/utils/page_gorm"
	"net/url"
	"strings"
)

// AdminRoleService struct
type AdminRoleService struct {
	BaseService
}

// GetCount 获取admin_role 总数
func (*AdminRoleService) GetCount() int {
	var count int64
	err := mysql.DB.Model(models_gorm.AdminRoles{}).Count(&count).Error
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllData 获取所有admin role
func (*AdminRoleService) GetAllData() []*models_gorm.AdminRoles {
	var adminRoles []*models_gorm.AdminRoles
	mysql.DB.Model(models_gorm.AdminRoles{}).Find(&adminRoles)

	return adminRoles
}

// GetPaginateData 分页获取adminrole
func (ars *AdminRoleService) GetPaginateData(listRows int, params url.Values) ([]*models_gorm.AdminRoles, page_gorm.Pagination) {
	//搜索、查询字段赋值
	ars.SearchField = append(ars.SearchField, new(models_gorm.AdminRoles).SearchField()...)

	var adminRole []*models_gorm.AdminRoles
	o := mysql.DB.Model(models_gorm.AdminRoles{})
	err := ars.PaginateAndScopeWhere(o, listRows, params).Find(&adminRole).Error
	if err != nil {
		return nil, ars.Pagination
	}
	return adminRole, ars.Pagination
}

// IsExistName 名称验重
func (*AdminRoleService) IsExistName(name string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models_gorm.AdminRoles)).Filter("name", name).Exist()
	}
	return orm.NewOrm().QueryTable(new(models_gorm.AdminRoles)).Filter("name", name).Exclude("id", id).Exist()
}

// Create 创建角色
func (*AdminRoleService) Create(form *formvalidate.AdminRoleForm) int {
	adminRole := models_gorm.AdminRoles{
		Name:        form.Name,
		Description: form.Description,
		URL:         "1,2,18",
		Status:      form.Status,
	}

	err := mysql.DB.Model(adminRole).Create(&adminRole).Error
	if err != nil {
		return 0
	}
	return int(adminRole.ID)
}

// GetAdminRoleById 通过id获取菜单信息
func (*AdminRoleService) GetAdminRoleById(id int) *models_gorm.AdminRoles {
	var adminRole models_gorm.AdminRoles
	err := mysql.DB.Model(adminRole).First(&adminRole, id).Error

	if err == nil {
		return &adminRole
	}
	return nil
}

// Update 更新角色信息
func (*AdminRoleService) Update(form *formvalidate.AdminRoleForm) int {
	result := mysql.DB.Model(new(models_gorm.AdminRoles)).Where("id = ?", form.Id).Updates(map[string]interface{}{
		"name":        form.Name,
		"description": form.Description,
		"status":      form.Status,
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}

// Del 删除角色
func (*AdminRoleService) Del(ids []int) int {
	result := mysql.DB.Where("id in ?", ids).Delete(models_gorm.AdminRoles{})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}

// Enable 启用角色
func (*AdminRoleService) Enable(ids []int) int {
	result := mysql.DB.Model(new(models_gorm.AdminRoles)).Where("id in ?", ids).Updates(map[string]interface{}{
		"status": 1,
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}

// Disable 禁用角色
func (*AdminRoleService) Disable(ids []int) int {
	result := mysql.DB.Model(new(models_gorm.AdminRoles)).Where("id in ?", ids).Updates(map[string]interface{}{
		"status": 0,
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}

// StoreAccess 授权菜单
func (*AdminRoleService) StoreAccess(id int, url []string) int {
	result := mysql.DB.Model(new(models_gorm.AdminRoles)).Where("id = ?", id).Updates(map[string]interface{}{
		"url": strings.Join(url, ","),
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}
