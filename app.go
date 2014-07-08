package main

import (
	"./routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &routers.AdminIndex{}, "get:Index")
	beego.Router("/login", &routers.Signin{})
	beego.Router("/video", &routers.Video{})
	beego.SessionOn = true
	beego.AutoRender = false
	beego.Run()
}
