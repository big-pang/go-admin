package services

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"go-admin/global"
	"go-admin/initialize/mysql"
	"go-admin/models_gorm"
	"go-admin/utils"
	"go-admin/utils/encrypter"
	"go-admin/utils/page_gorm"

	"net/url"
	"time"
)

// AdminLogService struct
type AdminLogService struct {
	BaseService
}

// CreateAdminLog 创建操作日志
func (*AdminLogService) CreateAdminLog(loginUser *models_gorm.AdminUsers, menu *models_gorm.AdminMenus, url string, ctx *context.Context) {
	var adminLog models_gorm.AdminLogs

	if loginUser == nil {
		adminLog.AdminUserID = 0
	} else {
		adminLog.AdminUserID = loginUser.ID
	}
	adminLog.Name = menu.Name
	adminLog.LogMethod = menu.LogMethod
	adminLog.URL = url
	adminLog.LogIP = ctx.Input.IP()
	adminLog.CreatedAt = time.Now()
	adminLog.UpdatedAt = time.Now()

	//开启事务
	tx := mysql.DB.Begin()

	err := tx.Create(&adminLog).Error
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return
	}
	requestData := ctx.Request.PostForm

	for key, _ := range requestData {
		if utils.InArrayForString([]string{"pwd", "key", "key2"}, key) {
			requestData.Set(key, "******")
		}
		if utils.InArrayForString([]string{"password", "new_password", "renew_password"}, key) {
			requestData.Set(key, utils.GetMd5String(requestData.Get(key)))
		}
	}
	//adminLogData数据表添加数据
	jsonData, _ := json.Marshal(requestData)
	cryptData := encrypter.Encrypt(jsonData, []byte(global.BA_CONFIG.Other.LogAesKey))
	var adminLogData models_gorm.AdminLogData
	adminLogData.AdminLogID = adminLog.ID
	adminLogData.Data = cryptData
	err = tx.Create(&adminLogData).Error
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return
	}
	tx.Commit()
}

// LoginLog 登录日志
func (*AdminLogService) LoginLog(loginUserID int, ctx *context.Context) {
	var adminLog models_gorm.AdminLogs
	adminLog.AdminUserID = loginUserID
	adminLog.Name = "登录"
	adminLog.URL = "admin/auth/login"
	adminLog.LogMethod = "POST"
	adminLog.LogIP = ctx.Input.IP()
	adminLog.CreatedAt = time.Now()
	adminLog.UpdatedAt = time.Now()

	//开启事务
	tx := mysql.DB.Begin()
	err := tx.Create(&adminLog).Error
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return
	}

	requestData := ctx.Request.PostForm

	for key, _ := range requestData {
		if utils.InArrayForString([]string{"pwd", "key", "key2"}, key) {
			requestData.Set(key, "******")
		}
		if utils.InArrayForString([]string{"password", "new_password", "renew_password"}, key) {
			requestData.Set(key, utils.GetMd5String(requestData.Get(key)))
		}
	}
	//adminLogData数据表添加数据
	jsonData, _ := json.Marshal(requestData)
	cryptData := encrypter.Encrypt(jsonData, []byte(global.BA_CONFIG.Other.LogAesKey))

	var adminLogData models_gorm.AdminLogData
	adminLogData.AdminLogID = adminLog.ID
	adminLogData.Data = cryptData
	err = tx.Create(&adminLogData).Error
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return
	}
	tx.Commit()
}

// GetCount 获取admin_log 总数
func (*AdminLogService) GetCount() int {
	var count int64
	err := mysql.DB.Model(models_gorm.AdminLogs{}).Count(&count).Error
	if err != nil {
		return 0
	}
	return int(count)
}

// GetPaginateData 获取所有adminuser
func (als *AdminLogService) GetPaginateData(listRows int, params url.Values) ([]*models_gorm.AdminLogs, page_gorm.Pagination) {
	//搜索、查询字段赋值
	als.SearchField = append(als.SearchField, new(models_gorm.AdminLogs).SearchField()...)
	als.WhereField = append(als.WhereField, new(models_gorm.AdminLogs).WhereField()...)
	als.TimeField = append(als.TimeField, new(models_gorm.AdminLogs).TimeField()...)

	var adminLog []*models_gorm.AdminLogs
	o := mysql.DB.Model(models_gorm.AdminLogs{})
	err := als.PaginateAndScopeWhere(o, listRows, params).Find(&adminLog).Error
	if err != nil {
		return nil, als.Pagination
	}
	return adminLog, als.Pagination
}

func (als *AdminLogService) GetItemAdminLog(id int) *models_gorm.AdminLogs {
	var adminLog models_gorm.AdminLogs
	err := mysql.DB.Model(adminLog).Preload("AdminUser").Joins("Data").First(&adminLog, id).Error
	if err != nil {
		return nil
	}
	logs.Debug(fmt.Sprintf("数据为：%+v\n", adminLog))
	return &adminLog
}
