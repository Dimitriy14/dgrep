package logout

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {

	this.DelSession("session")
	this.Redirect("/", 303)
}
