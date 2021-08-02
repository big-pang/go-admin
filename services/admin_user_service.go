package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"go-admin/formvalidate"
	"go-admin/global"
	"go-admin/initialize/mysql"
	"go-admin/models_gorm"
	"go-admin/utils"
	"go-admin/utils/page_gorm"
	"net/url"
	"strconv"
)

// AdminUserService struct
type AdminUserService struct {
	BaseService
}

// GetAdminUserById 根据id获取一条admin_user数据
func (*AdminUserService) GetAdminUserById(id int) *models_gorm.AdminUsers {
	adminUser := models_gorm.AdminUsers{}
	err := mysql.DB.First(&adminUser, id).Error
	if err != nil {
		logs.Error(err)
	}
	return &adminUser
}

// AuthCheck 权限检测
func (*AdminUserService) AuthCheck(url string, authExcept map[string]interface{}, loginUser *models_gorm.AdminUsers) bool {
	authURL := loginUser.GetAuthUrl()
	fmt.Println("获取有权限的URL：")
	fmt.Println(authURL)
	if utils.KeyInMap(url, authExcept) || utils.KeyInMap(url, authURL) {
		return true
	}
	return false
}

// CheckLogin 用户登录验证
func (*AdminUserService) CheckLogin(loginForm formvalidate.LoginForm, ctx *context.Context) (*models_gorm.AdminUsers, error) {
	var adminUser models_gorm.AdminUsers

	err := mysql.DB.Model(models_gorm.AdminUsers{}).Where("username = ?", loginForm.Username).First(&adminUser).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	decodePasswdStr, err := base64.StdEncoding.DecodeString(adminUser.Password)

	if err != nil || !utils.PasswordVerify(loginForm.Password, string(decodePasswdStr)) {
		return nil, errors.New("密码错误")
	}

	if adminUser.Status != 1 {
		return nil, errors.New("用户被冻结")
	}

	ctx.Output.Session(global.LOGIN_USER, adminUser)
	if loginForm.Remember != "" {
		ctx.SetCookie(global.LOGIN_USER_ID, strconv.Itoa(adminUser.ID), 7200)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, adminUser.GetSignStrByAdminUser(ctx), 7200)
	} else {
		ctx.SetCookie(global.LOGIN_USER_ID, ctx.GetCookie(global.LOGIN_USER_ID), -1)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, ctx.GetCookie(global.LOGIN_USER_ID_SIGN), -1)
	}

	return &adminUser, nil

}

// GetCount 获取admin_user 总数
func (*AdminUserService) GetCount() int {
	var count int64
	err := mysql.DB.Model(models_gorm.AdminUsers{}).Count(&count)
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllAdminUser 获取所有adminuser
func (*AdminUserService) GetAllAdminUser() []*models_gorm.AdminUsers {
	var adminUser []*models_gorm.AdminUsers
	err := mysql.DB.Model(models_gorm.AdminUsers{}).Find(&adminUser).Error
	if err != nil {
		return nil
	}
	return adminUser
}

// UpdateNickName 系统管理-个人资料-修改昵称
func (*AdminUserService) UpdateNickName(id int, nickname string) int {
	result := mysql.DB.Model(models_gorm.AdminUsers{}).Where("id = ?", id).Updates(map[string]interface{}{
		"nickname": nickname,
	})

	if result.Error != nil || result.RowsAffected <= 0 {
		return 0
	}
	return int(result.RowsAffected)
}

// UpdatePassword 修改密码
func (*AdminUserService) UpdatePassword(id int, newPassword string) int {
	newPasswordForHash, err := utils.PasswordHash(newPassword)
	if err != nil {
		return 0
	}
	result := mysql.DB.Model(models_gorm.AdminUsers{}).Where("id = ?", id).Updates(map[string]interface{}{
		"password": base64.StdEncoding.EncodeToString([]byte(newPasswordForHash)),
	})

	if result.Error != nil || result.RowsAffected <= 0 {
		return 0
	}

	return int(result.RowsAffected)
}

// UpdateAvatar 系统管理-个人资料-修改头像
func (*AdminUserService) UpdateAvatar(id int, avatar string) int {
	result := mysql.DB.Model(models_gorm.AdminUsers{}).Where("id = ?", id).Updates(map[string]interface{}{
		"avatar": avatar,
	})
	if result.Error != nil || result.RowsAffected <= 0 {
		return 0
	}
	return int(result.RowsAffected)

}

// GetPaginateData 通过分页获取adminuser
func (aus *AdminUserService) GetPaginateData(listRows int, params url.Values) ([]*models_gorm.AdminUsers, page_gorm.Pagination) {
	//搜索、查询字段赋值
	aus.SearchField = append(aus.SearchField, new(models_gorm.AdminUsers).SearchField()...)

	var adminUser []*models_gorm.AdminUsers
	o := mysql.DB.Model(models_gorm.AdminUsers{})
	err := aus.PaginateAndScopeWhere(o, listRows, params).Find(&adminUser).Error
	if err != nil {
		return nil, aus.Pagination
	}
	return adminUser, aus.Pagination
}

// IsExistName 名称验重
func (*AdminUserService) IsExistName(username string, id int) bool {
	var count int64
	if id == 0 {
		mysql.DB.Model(models_gorm.AdminUsers{}).Where("username", username).Count(&count)
	} else {
		mysql.DB.Model(models_gorm.AdminUsers{}).Where("username", username).Where("id != ?", id).Count(&count)
	}
	return count > 0
}

// Create 新增admin user用户
func (*AdminUserService) Create(form *formvalidate.AdminUserForm) int {
	newPasswordForHash, err := utils.PasswordHash(form.Password)
	if err != nil {
		return 0
	}

	adminUser := models_gorm.AdminUsers{
		Username: form.Username,
		Password: base64.StdEncoding.EncodeToString([]byte(newPasswordForHash)),
		Nickname: form.Nickname,
		Avatar:   form.Avatar,
		Role:     form.Role,
		Status:   int8(form.Status),
	}

	err = mysql.DB.Create(&adminUser).Error

	if err == nil {
		return int(adminUser.ID)
	}
	return 0
}

// Update 更新用户信息
func (*AdminUserService) Update(form *formvalidate.AdminUserForm) int {
	adminUser := models_gorm.AdminUsers{}

	err := mysql.DB.Where("id = ?", form.Id).First(&adminUser).Error
	if err == nil {
		adminUser.Username = form.Username
		adminUser.Nickname = form.Nickname
		adminUser.Role = form.Role
		adminUser.Status = int8(form.Status)
		if adminUser.Password != form.Password {
			newPasswordForHash, err := utils.PasswordHash(form.Password)
			if err == nil {
				adminUser.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))
			}
		}
		result := mysql.DB.Model(models_gorm.AdminUsers{}).Updates(&adminUser)
		if result.Error == nil {
			return int(result.RowsAffected)
		}
		return 0
	}
	return 0
}

// Enable 启用用户
func (*AdminUserService) Enable(ids []int) int {
	result := mysql.DB.Model(models_gorm.AdminUsers{}).Where("id in ?", ids).Updates(map[string]interface{}{
		"status": 1,
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}

// Disable 禁用用户
func (*AdminUserService) Disable(ids []int) int {
	result := mysql.DB.Model(models_gorm.AdminUsers{}).Where("id in ?", ids).Updates(map[string]interface{}{
		"status": 0,
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}

// Del 删除用户
func (*AdminUserService) Del(ids []int) int {
	result := mysql.DB.Where("id in ?", ids).Delete(models_gorm.AdminUsers{})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}
