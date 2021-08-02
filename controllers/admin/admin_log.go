package admin

import (
	"go-admin/global/response"
	"go-admin/services"
)

// AdminLogController struct.
type AdminLogController struct {
	baseController
}

// Index index.
func (alc *AdminLogController) Index() {
	var (
		adminLogService  services.AdminLogService
		adminUserService services.AdminUserService
	)
	data, pagination := adminLogService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	alc.Data["admin_user_list"] = adminUserService.GetAllAdminUser()

	alc.Data["data"] = data
	alc.Data["paginate"] = pagination
	alc.Layout = "admin/public/base.html"
	alc.TplName = "admin/admin_log/index.html"
}

// View view.
func (alc *AdminLogController) View() {
	var (
		adminLogService services.AdminLogService
	)
	id, _ := alc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", alc.Ctx)
		response.Error(alc.Ctx)
	}

	data := adminLogService.GetItemAdminLog(id)
	if data == nil {
		response.ErrorWithMessage("Not Found Info By Id.", alc.Ctx)
		response.Error(alc.Ctx)
	}

	alc.Data["data"] = data
	alc.Layout = "admin/public/log_data_base.html"
	alc.TplName = "admin/admin_log/view.html"
}
