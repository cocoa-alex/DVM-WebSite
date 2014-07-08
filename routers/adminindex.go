package routers

type AdminIndex struct {
	BaseRouter
}

func (u *AdminIndex) Index() {

	u.TplNames = "admin/layout.html"
	u.TplNames = "admin/index.html"
	u.Render()
}
