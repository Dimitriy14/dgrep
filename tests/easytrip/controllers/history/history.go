package historyController

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/repository"
)

type HistoryController struct {
	beego.Controller
}

func (this *HistoryController) Get() {

	this.TplName = "history.tpl"
	session := this.GetSession("session")
	if session == nil {
		return
	}
	usermap := session.(map[string]interface{})
	var user models.User
	user.Name = usermap["name"].(string)
	user.Login = usermap["login"].(string)
	user.Password = usermap["password"].(string)

	r, err := repository.HistoryView(user)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Req"] = r

}
