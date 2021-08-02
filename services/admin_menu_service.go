package services

import (
	"github.com/beego/beego/v2/core/logs"
	"go-admin/formvalidate"
	"go-admin/initialize/mysql"
	"go-admin/models_gorm"
)

// AdminMenuService struct
type AdminMenuService struct {
}

// GetAdminMenuByUrl 根据url获取admin_menu数据
func (*AdminMenuService) GetAdminMenuByUrl(url string) *models_gorm.AdminMenus {
	var adminMenu models_gorm.AdminMenus
	result := mysql.DB.Model(models_gorm.AdminMenus{}).
		Where("url = ?", url).
		First(&adminMenu)
	if result.Error == nil {
		return &adminMenu
	}
	return &adminMenu
}

// GetCount 获取admin_menu 总数
func (*AdminMenuService) GetCount() int {
	var count int64
	err := mysql.DB.Model(models_gorm.AdminMenus{}).Count(&count).Error
	if err != nil {
		logs.Error("获取admin_menu总数出错：", err)
	}
	return int(count)
}

// AllMenu 获取所有菜单
func (*AdminMenuService) AllMenu() []*models_gorm.AdminMenus {
	var adminMenus []*models_gorm.AdminMenus
	err := mysql.DB.Model(models_gorm.AdminMenus{}).Order("sort_id,id").Find(&adminMenus).Error
	if err != nil {
		logs.Error("获取所有菜单出错：", err)
	}
	return adminMenus
}

// Menu 除去当前id之外的所有菜单id
func (*AdminMenuService) Menu(currentID int) []models_gorm.Params {
	var adminMenusMap []models_gorm.Params
	var adminMenus []models_gorm.AdminMenus
	err := mysql.DB.Model(models_gorm.AdminMenus{}).Where("id != ?", currentID).Order("sort_id, id").Select("id", "parent_id", "name", "sort_id").Find(&adminMenus).Error
	if err == nil {
		for _, v := range adminMenus {
			adminMenusMap = append(adminMenusMap, models_gorm.Params{
				"Id":       v.ID,
				"ParentId": v.ParentID,
				"Name":     v.Name,
				"SortId":   v.SortID,
			})
		}
	}

	return adminMenusMap
}

// Create 创建菜单
func (*AdminMenuService) Create(form *formvalidate.AdminMenuForm) (id int64, err error) {
	adminMenu := models_gorm.AdminMenus{
		ParentID:  form.ParentId,
		Name:      form.Name,
		URL:       form.Url,
		Icon:      form.Icon,
		IsShow:    form.IsShow,
		SortID:    form.SortId,
		LogMethod: form.LogMethod,
	}
	err = mysql.DB.Model(models_gorm.AdminMenus{}).Create(&adminMenu).Error
	if err != nil {
		return 0, err
	}
	return int64(adminMenu.ID), nil
}

// Update 更新菜单
func (*AdminMenuService) Update(form *formvalidate.AdminMenuForm) int {
	result := mysql.DB.Model(models_gorm.AdminMenus{}).Where("id = ?", form.Id).Updates(map[string]interface{}{
		"parent_id":  form.ParentId,
		"name":       form.Name,
		"url":        form.Url,
		"icon":       form.Icon,
		"is_show":    form.IsShow,
		"sort_id":    form.SortId,
		"log_method": form.LogMethod,
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}

	return 0
}

// IsExistUrl Url验重
func (*AdminMenuService) IsExistUrl(url string, id int) bool {
	var count int64
	if id == 0 {
		mysql.DB.Model(models_gorm.AdminMenus{}).Where("url = ?", url).Count(&count)
	}
	mysql.DB.Model(models_gorm.AdminMenus{}).Where("url = ?", url).Where("id != ?", id).Count(&count)

	return count > 0
}

// IsChildMenu 判断是否有子菜单
func (*AdminMenuService) IsChildMenu(ids []int) bool {
	var count int64
	mysql.DB.Model(models_gorm.AdminMenus{}).Where("parent_id in ?", ids).Count(&count)
	return count > 0
}

// Del 删除菜单
func (*AdminMenuService) Del(ids []int) int {
	tx := mysql.DB.Where("id  in ids", ids).Delete(models_gorm.AdminMenus{})
	if tx.Error == nil {
		return int(tx.RowsAffected)
	}
	return 0
}

// GetAdminMenuById 通过id获取菜单信息
func (*AdminMenuService) GetAdminMenuById(id int) *models_gorm.AdminMenus {
	var adminMenu models_gorm.AdminMenus
	err := mysql.DB.Model(models_gorm.AdminMenus{}).Where("id = ?", id).First(&adminMenu).Error
	if err == nil {
		return &adminMenu
	}
	return nil
}
