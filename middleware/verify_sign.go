package middleware

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"go-admin/global/response"
	"go-admin/utils"
	"strings"
)

//不需要验证的url
var verifySignExcept = []string{}

func VerifySign(ctx *context.Context) {
	uri := ctx.Input.URI()
	//需要进行登录验证
	if !utils.InArrayForString(verifySignExcept, strings.ToLower(uri)) {
		//首先获取X-HM-Time头部
		secretToken := "AFghNfxYNBjMR32S"
		getTime := ctx.Request.Header.Get("X-HM-Time")
		getSign := ctx.Request.Header.Get("X-HM-Credential")
		json := ctx.Input.RequestBody
		signStr := uri + "&" + string(json) + "&" + secretToken + "&" + getTime
		sign := utils.GetMd5String(signStr)
		logs.Debug(sign)
		if sign != getSign {
			response.ErrorWithMessage("哦豁,签名错误", ctx)
			return
		}
	}
}
