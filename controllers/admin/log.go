package admin

type LogController struct {
	BaseController
}

func (c *LogController) HomeLog() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/log/home.html"
}

func (c *LogController) AdminLog() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/log/admin.html"
}
