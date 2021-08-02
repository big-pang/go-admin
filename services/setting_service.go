package services

import (
	"encoding/json"
	"go-admin/global"
	"go-admin/initialize/mysql"
	"go-admin/models_gorm"
)

// SettingService struct
type SettingService struct {
	BaseService
}

// Show 显示设置
func (settingService *SettingService) Show(id int) []*models_gorm.Setting {
	data := settingService.getDataBySettingGroupId(id)

	var settingFormService SettingFormService

	for key, value := range data {
		//contentNew := ""
		//value.Content转为json
		var contents []*models_gorm.Content
		if value.Content == "" {
			continue
		}
		err := json.Unmarshal([]byte(value.Content), &contents)

		if err != nil {
			continue
		}

		var contentNew []*models_gorm.Content
		for _, content := range contents {
			content.Form = settingFormService.GetFieldForm(content.Type, content.Name, content.Field, content.Content, content.Option)
			contentNew = append(contentNew, content)
		}
		data[key].ContentStrut = contentNew
	}

	return data
}

// getDataBySettingGroupId 根据设置分组id获取多个设置信息
func (*SettingService) getDataBySettingGroupId(settingGroupId int) []*models_gorm.Setting {
	var settings []*models_gorm.Setting
	err := mysql.DB.Model(models_gorm.Setting{}).Where("setting_group_id = ?", settingGroupId).Find(&settings)
	if err != nil {
		return nil
	}
	return settings
}

// GetSettingInfoById 根据设置id，获取对应的setting info
func (*SettingService) GetSettingInfoById(id int) *models_gorm.Setting {
	var setting models_gorm.Setting
	err := mysql.DB.Model(setting).First(&setting, id).Error
	if err != nil {
		return nil
	}
	return &setting
}

// UpdateSettingInfoToContent 根据id修改content的内容
func (*SettingService) UpdateSettingInfoToContent(id int, content string) int {
	result := mysql.DB.Model(new(models_gorm.Setting)).Where("id = ?", id).Updates(map[string]interface{}{
		"content": content,
	})
	if result.Error == nil {
		return int(result.RowsAffected)
	}
	return 0
}

// LoadOrUpdateGlobalBaseConfig 加载或者更新全局登录、系统配置信息
func (*SettingService) LoadOrUpdateGlobalBaseConfig(setting *models_gorm.Setting) bool {
	if setting == nil {
		return false
	}

	if setting.Code == "base" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "name":
				global.BA_CONFIG.Base.Name = content.Content
			case "short_name":
				global.BA_CONFIG.Base.ShortName = content.Content
			case "author":
				global.BA_CONFIG.Base.Author = content.Content
			case "version":
				global.BA_CONFIG.Base.Version = content.Content
			case "link":
				global.BA_CONFIG.Base.Link = content.Content
			}
		}
	} else if setting.Code == "login" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "token":
				global.BA_CONFIG.Login.Token = content.Content
			case "captcha":
				global.BA_CONFIG.Login.Captcha = content.Content
			case "background":
				global.BA_CONFIG.Login.Background = content.Content
			}
		}
	} else if setting.Code == "index" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "password_warning":
				global.BA_CONFIG.Base.PasswordWarning = content.Content
			case "show_notice":
				global.BA_CONFIG.Base.ShowNotice = content.Content
			case "notice_content":
				global.BA_CONFIG.Base.NoticeContent = content.Content
			}
		}
	} else {
		return false
	}

	return true
}
