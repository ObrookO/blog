package admin

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/role/index.html"
}
