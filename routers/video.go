package routers

type Video struct {
	BaseRouter
}

func (u *Video) Get() {
	u.TplNames = "test.html"
	u.Render()
}
