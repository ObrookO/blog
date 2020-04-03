package admin

type ManagerController struct {
	BaseController
}

func (c *ManagerController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/manager/index.html"
}
