package controllers

import (
	"blog/models"
)

type IndexController struct {
	BaseController
}

// Get 文章列表
func (c *IndexController) Get() {
	c.Layout = "layouts/master.html"
	c.TplName = "index/index.html"
	c.LayoutSections = map[string]string{
		"Style":  "index/index_style.html",
		"Script": "index/index_script.html",
	}

	p, _ := c.GetInt("p", 1)                 // 页码
	offset := (p - 1) * articleAmountPerPage // 偏移量

	// 文章列表
	filter := map[string]interface{}{"status": 1}
	articles, _ := models.GetAllArticles(filter, offset, articleAmountPerPage)

	// 判断是否存在下一页
	if len(articles) > 0 {
		if models.IsArticleExists(map[string]interface{}{"id__lt": articles[len(articles)-1].Id}) {
			c.Data["hasNext"] = true
		} else {
			c.Data["hasNext"] = false
		}
	} else {
		c.Data["hasNext"] = false
	}

	c.Data["articles"] = articles
	c.Data["p"] = p
}
