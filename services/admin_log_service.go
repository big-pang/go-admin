package services

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/server/web/context"
	"go-admin/global"
	"go-admin/models"
	"go-admin/utils/encrypter"
	"go-admin/utils/page"

	"github.com/beego/beego/v2/client/orm"
	"net/url"
	"time"
)

// AdminLogService struct
type AdminLogService struct {
	BaseService
}

// CreateAdminLog 创建操作日志
func (*AdminLogService) CreateAdminLog(loginUser *models.AdminUser, menu *models.AdminMenu, url string, ctx *context.Context) {
	var adminLog models.AdminLog

	if loginUser == nil {
		adminLog.AdminUserId = 0
	} else {
		adminLog.AdminUserId = loginUser.Id
	}
	adminLog.Name = menu.Name
	adminLog.LogMethod = menu.LogMethod
	adminLog.Url = url
	adminLog.LogIp = ctx.Input.IP()
	adminLog.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	adminLog.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	o := orm.NewOrm()
	//开启事务
	to, err := o.Begin()

	adminLogID, err := to.Insert(&adminLog)
	if err != nil {
		to.Rollback()
		beego.Error(err)
		return
	}

	//adminLogData数据表添加数据
	jsonData, _ := json.Marshal(ctx.Request.PostForm)
	cryptData := encrypter.Encrypt(jsonData, []byte(global.BA_CONFIG.Other.LogAesKey))
	var adminLogData models.AdminLogData
	adminLogData.AdminLogId = int(adminLogID)
	adminLogData.Data = cryptData
	_, err = to.Insert(&adminLogData)
	if err != nil {
		to.Rollback()
		beego.Error(err)
		return
	}
	to.Commit()
}

// LoginLog 登录日志
func (*AdminLogService) LoginLog(loginUserID int, ctx *context.Context) {
	var adminLog models.AdminLog
	adminLog.AdminUserId = loginUserID
	adminLog.Name = "登录"
	adminLog.Url = "admin/auth/login"
	adminLog.LogMethod = "POST"
	adminLog.LogIp = ctx.Input.IP()
	adminLog.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	adminLog.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	o := orm.NewOrm()

	//开启事务
	to, err := o.Begin()

	adminLogID, err := o.Insert(&adminLog)
	if err != nil {
		to.Rollback()
		beego.Error(err)
		return
	}

	//adminLogData数据表添加数据
	jsonData, _ := json.Marshal(ctx.Request.PostForm)
	cryptData := encrypter.Encrypt(jsonData, []byte(global.BA_CONFIG.Other.LogAesKey))

	var adminLogData models.AdminLogData
	adminLogData.AdminLogId = int(adminLogID)
	adminLogData.Data = cryptData
	_, err = o.Insert(&adminLogData)
	if err != nil {
		to.Rollback()
		beego.Error(err)
		return
	}
	to.Commit()
}

// GetCount 获取admin_log 总数
func (*AdminLogService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminLog)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetPaginateData 获取所有adminuser
func (als *AdminLogService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminLog, page.Pagination) {
	//搜索、查询字段赋值
	als.SearchField = append(als.SearchField, new(models.AdminLog).SearchField()...)
	als.WhereField = append(als.WhereField, new(models.AdminLog).WhereField()...)
	als.TimeField = append(als.TimeField, new(models.AdminLog).TimeField()...)

	var adminLog []*models.AdminLog
	o := orm.NewOrm().QueryTable(new(models.AdminLog))
	_, err := als.PaginateAndScopeWhere(o, listRows, params).All(&adminLog)
	if err != nil {
		return nil, als.Pagination
	}
	return adminLog, als.Pagination
}
