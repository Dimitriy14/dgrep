package statistics

import (
	_ "github.com/oreuta/easytrip/controllers"
	_ "github.com/oreuta/easytrip/controllers/bank-rating"
	_ "github.com/oreuta/easytrip/controllers/best-bank"
	"github.com/oreuta/easytrip/translate"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type StatisticController struct {
	beego.Controller
}

func (this *StatisticController) Get() {

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

	statistic := toolbox.StatisticsMap.GetMap()
	this.Data["Stat"] = statistic["Data"]
	this.Data["Statistic"] = statistic["Fields"]
	this.TplName = "statistic.tpl"
}
