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
func (index *IndexController) Get() {
	index.Layout = "layouts/master.html"
	index.TplName = "article/list.html"
	index.LayoutSections = make(map[string]string)
	index.LayoutSections["Style"] = "article/list_style.html"
	index.LayoutSections["Script"] = "article/list_script.html"

	// 每页的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(index.Ctx.Input.Query("p"))
	// 文章总数
	total, _ := models.GetTotal(make(map[string]interface{}))
	// 分页器
	p := utils.NewPaginator(index.Ctx.Request, per, total)

	articles := models.GetArticlesLimit(make(map[string]interface{}), (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), index.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := make(map[string]interface{})
		where["aid"] = value.Id
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	index.Data["username"] = index.GetSession("username")
	index.Data["Articles"] = articles
	index.Data["Categories"] = Categories
	index.Data["Tags"] = Tags
	index.Data["Archive"] = Archive
	index.Data["Like"] = Like
	index.Data["Paginator"] = p
}

// 点赞/取消点赞 功能
func (index *IndexController) Favor() {
	ip := index.Ctx.Input.IP()
	aid := index.GetString(":aid")

	// 构建查询条件
	where := make(map[string]interface{})
	where["id"] = aid

	if exist := models.IsArticleExists(where); !exist {
		index.Data["json"] = &JsonResult{Code: 400, Message: "Article Does Not Exist"}
	} else {
		key := "favor_" + aid

		// 如果没点过赞就点赞否则就取消
		// 把文章id，ip记录到redis中
		if isFavored := RedisClient.SIsMember(key, ip).Val(); isFavored {
			RedisClient.SRem(key, ip)
			index.Data["json"] = &JsonResult{Code: 200, Message: "Cancle Success"}
		} else {
			RedisClient.SAdd(key, ip)
			index.Data["json"] = &JsonResult{Code: 200, Message: "Favor Success"}
		}

	}

	index.ServeJSON()
}
