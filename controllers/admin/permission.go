package admin

type PermissionController struct {
	BaseController
}

func (c *PermissionController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/permission/index.html"
}
