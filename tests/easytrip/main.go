package main

import (
	"os"
	"strconv"

	"github.com/oreuta/easytrip/repository"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/oreuta/easytrip/clients"
	_ "github.com/oreuta/easytrip/routers"
)

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = beego.AppConfig.String("HTTPPort")
	}

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(port)
	beego.BConfig.WebConfig.Session.SessionOn = true
	if err != nil {
		panic(err)
	}
	go repository.Update()
	beego.Run()
}
