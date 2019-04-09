package regController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/services/registration"
)

type RegController struct {
	beego.Controller
	Regist registration.RegService
}

func (this *RegController) Get() {
	this.TplName = "registration.tpl"
}

func (this *RegController) Post() {
	this.TplName = "registration.tpl"

	u := models.User{
		Name:     this.GetString("name"),
		Login:    this.GetString("login"),
		Password: this.GetString("password"),
	}

	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		beego.Error("ValidationError: %v", err)
	}
	if !b {
		this.Data["Errors"] = valid.ErrorsMap
		return
	}
	a := this.Regist.CanRegistr(u)
	if a != nil {
		logs.Info(a)
		this.Data["Errors"] = "Error"
		return
	}

	this.Redirect("/login", 303)

}

func New(reg registration.RegService) *RegController {
	return &RegController{Regist: reg}
}
