package home

import (
	"blog/models"
	"strconv"

	"github.com/beego/wetalk/modules/utils"
)

type IndexController struct {
	BaseController
}

// 首页
func (c *IndexController) Get() {
	c.Layout = "home/layouts/master.html"
	c.TplName = "home/article/list.html"
	c.LayoutSections = map[string]string{
		"Style":  "home/article/list_style.html",
		"Script": "home/article/list_script.html",
	}

	// 每页的数量
	per := 10
	// 当前页
	page, _ := strconv.Atoi(c.Ctx.Input.Query("p"))

	where := map[string]interface{}{}
	// 获取查询条件
	search := c.GetString("search")
	if len(search) > 0 {
		where["title__icontains"] = search
	}

	// 文章总数
	total, _ := models.GetTotal(where)
	// 分页器
	p := utils.NewPaginator(c.Ctx.Request, per, total)

	articles := models.GetArticlesLimit(where, (page-1)*per, per)

	// 获取文章的点赞数以及当前IP的点赞情况
	for key, value := range articles {
		articles[key].IsFavored = RedisClient.SIsMember("favor_"+strconv.Itoa(value.Id), c.Ctx.Input.IP()).Val()
		articles[key].FavorNum = len(RedisClient.SMembers("favor_" + strconv.Itoa(value.Id)).Val())
		where := map[string]interface{}{
			"aid": value.Id,
		}
		articles[key].CommentNum, _ = models.GetCommentsCount(where)
	}

	c.Data = map[interface{}]interface{}{
		"isLogin":   IsLogin,
		"user":      UserInfo,
		"Search":    search,
		"Page":      page,
		"Paginator": p,
		"Articles":  articles,
		"Tags":      Tags,
		"Archive":   Archive,
	}
}

// 点赞/取消点赞 功能
func (c *IndexController) Favor() {
	ip := c.Ctx.Input.IP()
	aid := c.GetString(":aid")

	// 构建查询条件
	where := map[string]interface{}{
		"id": aid,
	}

	if exist := models.IsArticleExists(where); !exist {
		c.Data["json"] = &JSONResponse{Code: 400, Msg: "Article Does Not Exist"}
	} else {
		key := "favor_" + aid

		// 如果没点过赞就点赞否则就取消
		// 把文章id，ip记录到redis中
		if isFavored := RedisClient.SIsMember(key, ip).Val(); isFavored {
			RedisClient.SRem(key, ip)
			c.Data["json"] = &JSONResponse{Code: 200, Msg: "Cancle Success"}
		} else {
			RedisClient.SAdd(key, ip)
			c.Data["json"] = &JSONResponse{Code: 200, Msg: "Favor Success"}
		}
	}

	c.ServeJSON()
}
