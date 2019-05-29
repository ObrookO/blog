package controllers

import (
	"blog/models"
	"strconv"

	"github.com/beego/wetalk/modules/utils"
)

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type IndexController struct {
	BaseController
}

// 首页
func (this *IndexController) Get() {
	this.Layout = "layouts/master.html"
	this.TplName = "article/list.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Style"] = "article/list_style.html"
	this.LayoutSections["Script"] = "article/list_script.html"

	// 每页的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(this.Ctx.Input.Query("p"))

	where := make(map[string]interface{})
	// 获取查询条件
	search := this.GetString("search")
	if len(search) > 0 {
		where["title__icontains"] = search
	}

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
	this.Data["Categories"] = Categories
	this.Data["Tags"] = Tags
	this.Data["NuggetTags"] = NuggetTags
	this.Data["Archive"] = Archive
	this.Data["Like"] = Like
	this.Data["Paginator"] = p
	this.Data["Page"] = page
	this.Data["Search"] = search
}

// 点赞/取消点赞 功能
func (this *IndexController) Favor() {
	ip := this.Ctx.Input.IP()
	aid := this.GetString(":aid")

	// 构建查询条件
	where := make(map[string]interface{})
	where["id"] = aid

	if exist := models.IsArticleExists(where); !exist {
		this.Data["json"] = &JsonResult{Code: 400, Message: "Article Does Not Exist"}
	} else {
		key := "favor_" + aid

		// 如果没点过赞就点赞否则就取消
		// 把文章id，ip记录到redis中
		if isFavored := RedisClient.SIsMember(key, ip).Val(); isFavored {
			RedisClient.SRem(key, ip)
			this.Data["json"] = &JsonResult{Code: 200, Message: "Cancle Success"}
		} else {
			RedisClient.SAdd(key, ip)
			this.Data["json"] = &JsonResult{Code: 200, Message: "Favor Success"}
		}

	}

	this.ServeJSON()
}
