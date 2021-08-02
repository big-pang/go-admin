package services

import (
	"go-admin/initialize/mysql"
	"go-admin/utils"
	"time"
)

// DbVersion struct
type DbVersion struct {
	DbVersion string
}

// DatabaseService struct
type DatabaseService struct {
}

type TableStatus struct {
	Name       string     `gorm:"column:Name"`
	Comment    string     `gorm:"column:Comment"`
	Engine     string     `gorm:"column:Engine"`
	Collation  string     `gorm:"column:Collation"`
	Rows       string     `gorm:"column:Rows"`
	DataLength string     `gorm:"column:Data_length"`
	CreatedAt  time.Time  `gorm:"column:Create_time"`
	UpdatedAt  *time.Time `gorm:"column:Update_time"`
}

// GetMysqlVersion 获取mysql的版本
func (*DatabaseService) GetMysqlVersion() string {
	var dbVersion DbVersion
	err := mysql.DB.Raw("select VERSION() as db_version").First(&dbVersion).Error
	if err != nil {
		return "not found."
	}
	return dbVersion.DbVersion
}

// GetTableStatus 获取所有数据表的状态
func (ds *DatabaseService) GetTableStatus() ([]map[string]string, int) {
	var resultMaps []map[string]string
	var tableStatus []TableStatus
	result := mysql.DB.Raw("SHOW TABLE STATUS").Find(&tableStatus)
	if result.RowsAffected > 0 && result.Error == nil {
		for _, item := range tableStatus {
			var createTimeStr = ""
			var updateTimeStr = ""
			createTimeStr = item.CreatedAt.Format(utils.TimeLayout)
			if item.UpdatedAt != nil {
				updateTimeStr = item.UpdatedAt.Format(utils.TimeLayout)
			}
			resultMaps = append(resultMaps, map[string]string{
				"name":        ds.nil2String(item.Name),
				"comment":     ds.nil2String(item.Comment),
				"engine":      ds.nil2String(item.Engine),
				"collation":   ds.nil2String(item.Collation),
				"rows":        ds.nil2String(item.Rows),
				"data_length": ds.nil2String(item.DataLength),
				"created_at":  ds.nil2String(createTimeStr),
				"updated_at":  ds.nil2String(updateTimeStr),
			})
		}
	}

	return resultMaps, int(result.RowsAffected)
}

// OptimizeTable 优化数据表
func (*DatabaseService) OptimizeTable(tableName string) bool {
	err := mysql.DB.Exec("OPTIMIZE TABLE `" + tableName + "`").Error
	if err == nil {
		return true
	}
	return false
}

// RepairTable 修复数据表
func (*DatabaseService) RepairTable(tableName string) bool {
	err := mysql.DB.Exec("REPAIR TABLE `" + tableName + "`").Error
	if err == nil {
		return true
	}
	return false
}

// GetFullColumnsFromTable 获取数据表的所有字段
func (ds *DatabaseService) GetFullColumnsFromTable(tableName string) []map[string]string {
	var resultMaps []map[string]string
	var fullColumens []struct {
		Field      string
		Type       string
		Collation  string
		Null       string
		Key        string
		Default    string
		Extra      string
		Privileges string
		Comment    string
	}

	result := mysql.DB.Raw("SHOW FULL COLUMNS FROM `" + tableName + "`").Find(&fullColumens)

	if result.RowsAffected > 0 && result.Error == nil {
		for _, item := range fullColumens {
			resultMaps = append(resultMaps, map[string]string{
				"name":       ds.nil2String(item.Field),
				"type":       ds.nil2String(item.Type),
				"collation":  ds.nil2String(item.Collation),
				"null":       ds.nil2String(item.Null),
				"key":        ds.nil2String(item.Key),
				"default":    ds.nil2String(item.Default),
				"extra":      ds.nil2String(item.Extra),
				"privileges": ds.nil2String(item.Privileges),
				"comment":    ds.nil2String(item.Comment),
			})
		}
	}

	return resultMaps
}

// nil2String interface 转换 为string
func (*DatabaseService) nil2String(val interface{}) string {
	if val == nil {
		return ""
	}
	return val.(string)
}
