package middleware

import (
	"fmt"

	"github.com/beego/beego/v2/server/web/context"
	"go-admin/global"
	"go-admin/global/response"
	"go-admin/models_gorm"
	"go-admin/services"
	"strconv"
	"strings"
)

//不需要验证的url
var authExcept = map[string]interface{}{
	"admin/auth/login":           0,
	"admin/auth/check_login":     1,
	"admin/auth/logout":          2,
	"admin/auth/captcha":         3,
	"admin/editor/server":        4,
	"admin/auth/refresh_captcha": 5,
}

// AuthMiddle 中间件
func AuthMiddle(ctx *context.Context) {

	//登录认证中间件过滤器
	url := strings.TrimLeft(ctx.Input.URL(), "/")
	//需要进行登录验证
	if !isAuthExceptUrl(strings.ToLower(url), authExcept) {
		//验证是否登录
		loginUser, isLogin := isLogin(ctx)
		if !isLogin {
			response.ErrorWithMessageAndUrl("未登录", "/admin/auth/login", ctx)
			return
		}

		//验证，是否有权限访问
		var adminUserService services.AdminUserService
		if loginUser.ID != 1 && !adminUserService.AuthCheck(url, authExcept, loginUser) {
			errorBackURL := global.URL_CURRENT
			if ctx.Request.Method == "GET" {
				errorBackURL = ""
			}
			response.ErrorWithMessageAndUrl("无权限", errorBackURL, ctx)
			return
		}
	}

	checkAuth, _ := strconv.Atoi(ctx.Request.PostForm.Get("check_auth"))

	if checkAuth == 1 {
		response.Success(ctx)
		return
	}

}

//判断是否是不需要验证登录的url,只针对admin模块路由的判断
func isAuthExceptUrl(url string, m map[string]interface{}) bool {
	urlArr := strings.Split(url, "/")
	if len(urlArr) > 3 {
		url = fmt.Sprintf("%s/%s/%s", urlArr[0], urlArr[1], urlArr[2])
	}
	_, ok := m[url]
	if ok {
		return true
	}
	return false
}

//是否登录
func isLogin(ctx *context.Context) (*models_gorm.AdminUsers, bool) {
	loginUser, ok := ctx.Input.Session(global.LOGIN_USER).(models_gorm.AdminUsers)
	if !ok {
		loginUserIDStr := ctx.GetCookie(global.LOGIN_USER_ID)
		loginUserIDSign := ctx.GetCookie(global.LOGIN_USER_ID_SIGN)

		if loginUserIDStr != "" && loginUserIDSign != "" {
			loginUserID, _ := strconv.Atoi(loginUserIDStr)
			var adminUserService services.AdminUserService
			loginUserPointer := adminUserService.GetAdminUserById(loginUserID)

			if loginUserPointer != nil && loginUserPointer.GetSignStrByAdminUser(ctx) == loginUserIDSign {
				ctx.Output.Session(global.LOGIN_USER, *loginUserPointer)
				return loginUserPointer, true
			}
		}
		return nil, false
	}

	return &loginUser, true
}
