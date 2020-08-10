package controllers

import (
	"blog/models"
	"blog/tool"
)

type ArticleController struct {
	BaseController
}

// Get 文章列表
func (c *ArticleController) Get() {
	c.Layout = "layouts/master.html"
	c.TplName = "article/list.html"
	c.LayoutSections = map[string]string{
		"Style":  "article/list_style.html",
		"Script": "article/list_script.html",
	}

	var articles []*models.Article

	filter := map[string]interface{}{"status": 1} // 查询文章的条件

	cid, _ := c.GetInt("cid")   // 栏目id
	tid, _ := c.GetInt("tid")   // 标签id
	mid, _ := c.GetInt("mid")   // 管理员id
	date := c.GetString("date") // 归档

	search := c.GetString("search") // 搜索关键词
	if tool.GetCharAmount(search) > 0 {
		filter["title__startswith"] = search
	}

	p, _ := c.GetInt("p", 1)                 // 页码
	offset := (p - 1) * articleAmountPerPage // 偏移量

	if cid > 0 {
		filter["category_id"] = cid
		articles, _ = models.GetAllArticlesOfCategory(filter, offset, articleAmountPerPage)
		c.Data["param"] = map[string]interface{}{"cid": cid}
	} else if tid > 0 {
		filter["tag_id"] = tid
		articles, _ = models.GetAllArticlesOfTag(filter, offset, articleAmountPerPage)
		c.Data["param"] = map[string]interface{}{"tid": tid}
	} else if mid > 0 {
		filter["manager_id"] = mid
		articles, _ = models.GetAllArticlesOfManager(filter, offset, articleAmountPerPage)
		c.Data["param"] = map[string]interface{}{"mid": mid}
	} else if date != "" {
		filter["created_at__startswith"] = date
		articles, _ = models.GetAllArticlesOfDate(filter, offset, articleAmountPerPage)
		c.Data["param"] = map[string]interface{}{"date": date}
	} else {
		articles, _ = models.GetAllArticles(filter, offset, articleAmountPerPage)
	}

	// 判断是否有下一页
	if len(articles) < articleAmountPerPage {
		c.Data["hasNext"] = false
	} else {
		c.Data["hasNext"] = true
	}

	c.Data["articles"] = articles
	c.Data["search"] = search
	c.Data["p"] = p
}

// GetArticleDetail 获取文章详情
func (c *ArticleController) GetArticleDetail() {
	c.Layout = "layouts/master.html"
	c.TplName = "article/detail.html"
	c.LayoutSections = map[string]string{
		"Script": "article/detail_script.html",
	}

	id, _ := c.GetInt(":id")
	article, _ := models.GetOneArticle(map[string]interface{}{"id": id, "status": 1})

	if article.Id == 0 {
		c.Abort("404")
	}

	c.Data["article"] = article
	c.Data["keywords"] = article.Keyword
}

// Favor 点赞功能
func (c *ArticleController) Favor() {
	id, _ := c.GetInt("id")
	ip := getClientIp(c.Ctx)

	article, _ := models.GetOneArticle(map[string]interface{}{"id": id, "status": 1})
	if article.Id == 0 {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "文章不存在"}
		c.ServeJSON()
		return
	}

	if models.IsFavorRecordExists(map[string]interface{}{"article_id": id, "ip": ip}) {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "已经点过赞啦"}
		c.ServeJSON()
		return
	}

	if _, err := models.AddFavorRecord(models.FavorRecord{
		Article: &models.Article{Id: id},
		Ip:      ip,
	}); err != nil {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "点赞失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK", Data: len(article.Favors) + 1}
	c.ServeJSON()
	return
}
