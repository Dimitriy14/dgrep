package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/oreuta/easytrip/translate"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	toolbox.StatisticsMap.AddStatistics("GET", "/", "&controllers.MainController", time.Duration(13000))

	this.Layout = "main_layout.tpl"
	this.TplName = "index.tpl"

	translate := translate.New()
	lang := this.GetString("lang")
	if lang != "" {
		translate.Lang = lang
		this.Ctx.SetCookie("lang", translate.Lang)
	} else {
		translate.Lang = this.Ctx.GetCookie("lang")
		if translate.Lang == "" {
			translate.Lang = "en-US"
		}
	}
	translate.Path = "conf/locale_" + translate.Lang + ".ini"
	this.Data["i18n"] = translate.Tr

	session := this.GetSession("session")
	if session == nil {
		return
	}
	this.Data["Registred"] = true
	this.Data["Session"] = session.(map[string]interface{})

}
