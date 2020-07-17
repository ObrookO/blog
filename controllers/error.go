package controllers

type ErrorController struct {
	BaseController
}

// Error401 401页面
func (c *ErrorController) Error401() {
	c.TplName = "error/401.html"
}

// Error404 404页面
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
}
