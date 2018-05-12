package main

import (
	//"fmt"
	"./controllers"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func main() {
	beego.SessionOn = true
	beego.RegisterController("/", &controllers.IndexController{})
	beego.RegisterController("/regist", &controllers.RegistController{})
	beego.Run()
}
