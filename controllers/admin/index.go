package admin

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/index/index.html"
}
