package routers

import (
	"github.com/astaxie/beego"
	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/controllers"
	"github.com/oreuta/easytrip/controllers/bank-rating"
	"github.com/oreuta/easytrip/controllers/best-bank"
	"github.com/oreuta/easytrip/controllers/history"
	"github.com/oreuta/easytrip/controllers/login"
	"github.com/oreuta/easytrip/controllers/logout"
	"github.com/oreuta/easytrip/controllers/registration"
	"github.com/oreuta/easytrip/controllers/statistics"
	"github.com/oreuta/easytrip/services/bank-rating"
	"github.com/oreuta/easytrip/services/best-bank"
	"github.com/oreuta/easytrip/services/registration"
)

func init() {
	ratesclient := clients.New()
	regServ := registration.New()
	ratesService := bankRatingService.New(ratesclient)
	ratesController := bankRatingController.New(ratesService)
	bestService := bestBankService.New(ratesclient)
	bestController := bestBankController.New(bestService)
	loginCont := login.New(regServ)
	regCont := regController.New(regServ)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/comparision", ratesController)
	beego.Router("/best", bestController)
	beego.Router("/statistics", &statistics.StatisticController{})
	beego.Router("/signup", regCont)
	beego.Router("/login", loginCont)
	beego.Router("/logout", &logout.LogoutController{})
	beego.Router("/history", &historyController.HistoryController{})
}
