package admin

type TagController struct {
	BaseController
}

func (c *TagController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/tag/index.html"
}
