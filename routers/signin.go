package routers

type Signin struct {
	BaseRouter
}

func (u *Signin) Get() {
	u.TplNames = "admin/login.html"
	u.Render()
}
