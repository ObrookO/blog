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
func (a *ArticleController) GetAllArticlesByCate() {
	a.Layout = "layouts/master.html"
	a.TplName = "article/list.html"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Style"] = "article/list_style.html"
	a.LayoutSections["Script"] = "article/list_script.html"

	// 栏目id
	cid := a.GetString(":cid")

	// 搜索条件
	where := make(map[string]interface{})
	where["cate"] = cid

	// 每页显示的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(a.Ctx.Input.Query("p"))
	// 文章总数
	total, _ := models.GetTotal(where)
	// 分页器
	p := utils.NewPaginator(a.Ctx.Request, per, total)

	articles := models.GetArticlesLimit(where, (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), a.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := make(map[string]interface{})
		where["aid"] = value.Id
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	a.Data["username"] = a.GetSession("username")
	a.Data["Categories"] = Categories
	a.Data["Articles"] = articles
	a.Data["Tags"] = Tags
	a.Data["Archive"] = Archive
	a.Data["Like"] = Like
	a.Data["Paginator"] = p
}

// 根据日期获取文章列表
func (a *ArticleController) GetArticlesByDate() {
	a.Layout = "layouts/master.html"
	a.TplName = "article/list.html"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Style"] = "article/list_style.html"
	a.LayoutSections["Script"] = "article/list_script.html"

	date := a.GetString(":year") + "-" + a.GetString(":month")

	createdAt, err := time.Parse("2006-01", date)

	if err != nil {
		a.Abort("404")
	}

	// 开始日期
	start := createdAt.Format("2006-01-02 15:04:05")
	// 结束日期
	end := createdAt.AddDate(0, 1, 0).Format("2006-01-02 15:04:05")

	// 搜索条件
	where := make(map[string]interface{})
	where["created_at__gt"] = start
	where["created_at__lt"] = end

	// 每页显示的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(a.Ctx.Input.Query("p"))
	// 文章总数
	total, _ := models.GetTotal(where)
	// 分页器
	p := utils.NewPaginator(a.Ctx.Request, per, total)

	articles := models.GetArticlesLimit(where, (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), a.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := make(map[string]interface{})
		where["aid"] = value.Id
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	a.Data["username"] = a.GetSession("username")
	a.Data["Categories"] = Categories
	a.Data["Articles"] = articles
	a.Data["Tags"] = Tags
	a.Data["Archive"] = Archive
	a.Data["Like"] = Like
	a.Data["Paginator"] = p

}

// 根据标签获取文章列表
func (a *ArticleController) GetArticlesByTag() {
	a.Layout = "layouts/master.html"
	a.TplName = "article/list.html"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Style"] = "article/list_style.html"
	a.LayoutSections["Script"] = "article/list_script.html"

	tid := a.GetString(":tid")

	// 每页显示的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(a.Ctx.Input.Query("p"))
	// 文章总数
	total := len(models.GetArticlesByTag("tags", tid, 0, 0))
	// 分页器
	p := utils.NewPaginator(a.Ctx.Request, per, total)

	articles := models.GetArticlesByTag("tags", tid, (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), a.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := make(map[string]interface{})
		where["aid"] = value.Id
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	a.Data["username"] = a.GetSession("username")
	a.Data["Categories"] = Categories
	a.Data["Articles"] = articles
	a.Data["Tags"] = Tags
	a.Data["Archive"] = Archive
	a.Data["Like"] = Like
	a.Data["Paginator"] = p

}

// 文章详情页
func (a *ArticleController) ArticleInfo() {
	a.Layout = "layouts/master.html"
	a.TplName = "article/detail.html"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Script"] = "article/detail_script.html"

	// 文章id
	aid, err := strconv.Atoi(a.GetString(":aid"))
	// 文章详情的查询条件
	where := make(map[string]interface{})
	where["id"] = aid
	// 获取文章详情
	info, err := models.GetOneArticle(where)
	if err != nil {
		a.Abort("404")
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

	a.Data["username"] = a.GetSession("username")
	a.Data["xsrf_token"] = a.XSRFToken()
	a.Data["Categories"] = Categories
	a.Data["Tags"] = Tags
	a.Data["Comments"] = comments
	a.Data["Archive"] = Archive
	a.Data["Like"] = Like
	a.Data["Info"] = info
	a.Data["Before"] = before
	a.Data["After"] = after
}
