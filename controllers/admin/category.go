package admin

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/category/index.html"
}
