package admin

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/user/index.html"
}
