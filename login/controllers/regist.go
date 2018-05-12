package controllers

import (
	"../models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	this.TplNames = "regist.tpl"
}

func (this *RegistController) Post() {
	var user models.User
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("username")
	user.Pwd = inputs.Get("pwd")
	err := models.SaveUser(user)
	if err == nil {
		this.TplNames = "success.tpl"
	} else {
		fmt.Println(err)
		this.TplNames = "error.tpl"
	}
}
