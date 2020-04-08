package controllers

import (
	"blog/models"
	"time"

	"strconv"
	"strings"

	"github.com/beego/wetalk/modules/utils"
)

type ArticleController struct {
	BaseController
}

// 根据栏目id获取文章列表
func (this *ArticleController) GetAllArticlesByCate() {
	this.Layout = "layouts/master.html"
	this.TplName = "article/list.html"
	this.LayoutSections = map[string]string{
		"Style":  "article/list_style.html",
		"Script": "article/list_script.html",
	}

	// 栏目id
	cid := this.GetString(":cid")

	// 搜索条件
	where := make(map[string]interface{})

	// 获取查询条件
	where["cate"] = cid
	search := this.GetString("search")
	if len(search) > 0 {
		where["title__icontains"] = search
	}

	// 每页显示的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(this.Ctx.Input.Query("p"))
	// 文章总数
	total, _ := models.GetTotal(where)
	// 分页器
	p := utils.NewPaginator(this.Ctx.Request, per, total)

	articles := models.GetArticlesLimit(where, (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), this.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := make(map[string]interface{})
		where["aid"] = value.Id
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	this.Data["username"] = this.GetSession("username")
	this.Data["Articles"] = articles
	this.Data["Tags"] = Tags
	this.Data["Archive"] = Archive
	this.Data["Like"] = Like
	this.Data["Paginator"] = p
	this.Data["Page"] = page
	this.Data["Search"] = search
}

// 根据日期获取文章列表
func (this *ArticleController) GetArticlesByDate() {
	this.Layout = "layouts/master.html"
	this.TplName = "article/list.html"
	this.LayoutSections = map[string]string{
		"Style":  "article/list_style.html",
		"Script": "article/list_script.html",
	}

	date := this.GetString(":year") + "-" + this.GetString(":month")

	createdAt, err := time.Parse("2006-01", date)

	if err != nil {
		this.Abort("404")
	}

	// 开始日期
	start := createdAt.Format("2006-01-02 15:04:05")
	// 结束日期
	end := createdAt.AddDate(0, 1, 0).Format("2006-01-02 15:04:05")

	// 搜索条件
	where := make(map[string]interface{})
	where["created_at__gt"] = start
	where["created_at__lt"] = end
	search := this.GetString("search")
	if len(search) > 0 {
		where["title__icontains"] = search
	}

	// 每页显示的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(this.Ctx.Input.Query("p"))
	// 文章总数
	total, _ := models.GetTotal(where)
	// 分页器
	p := utils.NewPaginator(this.Ctx.Request, per, total)

	articles := models.GetArticlesLimit(where, (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), this.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := make(map[string]interface{})
		where["aid"] = value.Id
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	this.Data["username"] = this.GetSession("username")
	this.Data["Articles"] = articles
	this.Data["Tags"] = Tags
	this.Data["Archive"] = Archive
	this.Data["Like"] = Like
	this.Data["Paginator"] = p
	this.Data["Page"] = page
	this.Data["Search"] = search
}

// 根据标签获取文章列表
func (this *ArticleController) GetArticlesByTag() {
	this.Layout = "layouts/master.html"
	this.TplName = "article/list.html"
	this.LayoutSections = map[string]string{
		"Style":  "article/list_style.html",
		"Script": "article/list_script.html",
	}

	tid := this.GetString(":tid")

	// 每页显示的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(this.Ctx.Input.Query("p"))
	// 文章总数
	total := len(models.GetArticlesByTag("tags", tid, 0, 0))
	// 分页器
	p := utils.NewPaginator(this.Ctx.Request, per, total)

	articles := models.GetArticlesByTag("tags", tid, (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), this.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := make(map[string]interface{})
		where["aid"] = value.Id
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	this.Data["username"] = this.GetSession("username")
	this.Data["Articles"] = articles
	this.Data["Tags"] = Tags
	this.Data["Archive"] = Archive
	this.Data["Like"] = Like
	this.Data["Paginator"] = p
}

// 文章详情页
func (this *ArticleController) ArticleInfo() {
	this.Layout = "layouts/master.html"
	this.TplName = "article/detail.html"
	this.LayoutSections = map[string]string{
		"Script": "article/detail_script.html",
	}

	// 文章id
	aid, err := strconv.Atoi(this.GetString(":aid"))
	// 文章详情的查询条件
	where := make(map[string]interface{})
	where["id"] = aid
	// 获取文章详情
	info, err := models.GetOneArticle(where)
	if err != nil {
		this.Abort("404")
	}

	// 获取文章的所有评论
	where2 := make(map[string]interface{})
	where2["aid"] = aid

	comments, _, err := models.GetComments(where2)

	// 查询文章的标签
	con := make(map[string]interface{})
	con["id__in"] = strings.Split(info.Tags, ",")
	info.Tags = strings.Join(models.GetTagFields(con, "name"), ", ")

	// 获取上一篇以及下一篇的id
	before, after := models.GetBeforeAndAfter(aid)

	this.Data["username"] = this.GetSession("username")
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["Tags"] = Tags
	this.Data["Comments"] = comments
	this.Data["Archive"] = Archive
	this.Data["Like"] = Like
	this.Data["Info"] = info
	this.Data["Before"] = before
	this.Data["After"] = after
}
