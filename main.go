package main

import (
	"DVM-WebSite/controllers/admin"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &admin.IndexControl{}, "get:Index")
	beego.SessionOn = true
	beego.AutoRender = false
	beego.Run()
}
