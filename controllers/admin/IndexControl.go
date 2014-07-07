package admin

import (
	"github.com/astaxie/beego"
)

type IndexControl struct {
	beego.Controller
}

func (u *IndexControl) Index() {
	u.TplNames = "admin/index.html"
	u.Render()
}
