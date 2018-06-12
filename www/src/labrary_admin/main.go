package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "labrary_admin/routers"
)

func main() {
	beego.Run()
}
