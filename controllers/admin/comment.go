package admin

type CommentController struct {
	BaseController
}

func (c *CommentController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/comment/index.html"
}

func (c *CommentController) Keyword() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/comment/keyword.html"
}
