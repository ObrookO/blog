package admin

type FavorRecordController struct {
	BaseController
}

func (c *FavorRecordController) Get() {
	c.Layout = "admin/layouts/master.html"
	c.TplName = "admin/favor_record/index.html"
}
