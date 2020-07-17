package controllers

import "blog/models"

type ResourceController struct {
	BaseController
}

// Get 学习资源页面
func (c *ResourceController) Get() {
	c.Layout = "layouts/master.html"
	c.TplName = "resource/list.html"

	list, _ := models.GetAllResource(nil)
	c.Data["list"] = list
}
