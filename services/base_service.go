package services

import (
	"github.com/beego/beego/v2/client/orm"
	"go-admin/utils"
	"go-admin/utils/page_gorm"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

// BaseService struct
type BaseService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
	//禁止删除的数据id,在model中声明就可以了，可不用在此处声明
	//NoDeletionId []int
	//分页
	Pagination page_gorm.Pagination
}

// Paginate 分页处理
func (bs *BaseService) Paginate(seter *gorm.DB, listRows int, parameters url.Values) *gorm.DB {
	var pagination page_gorm.Pagination
	qs := pagination.Paginate(seter, listRows, parameters)
	bs.Pagination = pagination
	return qs
}

// ScopeWhere 查询处理
func (bs *BaseService) ScopeWhere(seter *gorm.DB, parameters url.Values) *gorm.DB {
	//关键词like搜索
	keywords := parameters.Get("_keywords")
	cond := orm.NewCondition()
	if keywords != "" && len(bs.SearchField) > 0 {
		for _, v := range bs.SearchField {
			cond = cond.Or(v+"__icontains", keywords)
		}
	}

	//字段条件查询
	if len(bs.WhereField) > 0 && len(parameters) > 0 {
		for k, v := range parameters {
			if v[0] != "" && utils.InArrayForString(bs.WhereField, k) {
				cond = cond.And(k, v[0])
			}
		}
	}

	//时间范围查询
	if len(bs.TimeField) > 0 && len(parameters) > 0 {
		for key, value := range parameters {
			if value[0] != "" && utils.InArrayForString(bs.TimeField, key) {
				timeRange := strings.Split(value[0], " - ")
				startTimeStr := timeRange[0]
				endTimeStr := timeRange[1]

				cond = cond.And(key+" > ?", startTimeStr).And(key+" < ?", endTimeStr)

				//loc, _ := time.LoadLocation("Local")
				//startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, loc)
				//
				//if err == nil {
				//	unixStartTime := startTime.Unix()
				//	if len(endTimeStr) == 10 {
				//		endTimeStr += "23:59:59"
				//	}
				//
				//	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, loc)
				//	if err == nil {
				//		unixEndTime := endTime.Unix()
				//		cond = cond.And(key+"__gte", unixStartTime).And(key+"__lte", unixEndTime)
				//	}
				//}
			}
		}
	}

	//将条件语句拼装到主语句中
	seter = seter.Where(cond)

	//排序
	order := parameters.Get("_order")
	by := parameters.Get("_by")
	if order == "" {
		order = "id"
	}

	if by == "" {
		by = "-"
	} else {
		if by == "asc" {
			by = ""
		} else {
			by = "-"
		}
	}

	//排序
	seter = seter.Order(by + order)

	return seter
}

// PaginateAndScopeWhere 分页和查询合并，多用于首页列表展示、搜索
func (bs *BaseService) PaginateAndScopeWhere(seter *gorm.DB, listRows int, parameters url.Values) *gorm.DB {
	return bs.Paginate(bs.ScopeWhere(seter, parameters), listRows, parameters)
}
