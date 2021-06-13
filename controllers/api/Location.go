package api

import (
	beego "github.com/beego/beego/v2/adapter"
	"go-admin/utils/ipsearch"
)

type LocationController struct {
	beego.Controller
}

func (lc *LocationController) LocationAndTime() {
	ip := lc.Ctx.Input.IP()
	ip = "210.51.200.123"
	p, _ := ipsearch.New()
	location := p.GetLocation(ip)
	lc.Data["json"] = location
	lc.ServeJSON()
}
