package admin

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/article/index.html"
}

func (c *ArticleController) Draft() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/article/index.html"
}
