package controllers

import (
	"blog/models"

	"strconv"

	"github.com/beego/wetalk/modules/utils"
)

type NuggetController struct {
	BaseController
}

// Nugget文章列表
func (this *NuggetController) ArticleList() {
	this.Layout = "layouts/master.html"
	this.TplName = "nugget/list.html"

	where := make(map[string]interface{})
	// 获取查询条件
	cid := this.GetString(":cid")
	search := this.GetString("search")
	where["category"] = cid
	if len(search) > 0 {
		where["title__icontains"] = search
	}

	// 每页的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(this.Ctx.Input.Query("p"))
	// 总数
	total, _ := models.GetNuggetTotal(where)
	// 分页器
	p := utils.NewPaginator(this.Ctx.Request, per, total)

	articles := models.GetAllArticles(where, (page-1)*per, per)

	this.Data["username"] = this.GetSession("username")
	this.Data["Categories"] = Categories
	this.Data["NuggetTags"] = NuggetTags
	this.Data["Articles"] = articles
	this.Data["Paginator"] = p
	this.Data["Page"] = page
	this.Data["Search"] = search
}

// Nugget文章详情
func (this *NuggetController) ArticleDetail() {
	this.Layout = "layouts/master.html"
	this.TplName = "nugget/detail.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Script"] = "nugget/detail_script.html"

	where := make(map[string]interface{})

	aid, _ := strconv.Atoi(this.GetString(":aid"))
	where["id"] = aid

	article := models.GetArticleInfo(where)
	before, after := models.GetLastAndNext(aid)

	this.Data["Categories"] = Categories
	this.Data["NuggetTags"] = NuggetTags
	this.Data["Info"] = article
	this.Data["Before"] = before
	this.Data["After"] = after
}
