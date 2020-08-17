package controllers

import (
	"blog/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils"

	"html/template"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/context"
)

type BaseController struct {
	beego.Controller
}

type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	rc     cache.Cache // 缓存实例，优先使用redis，如出现错误，则使用内存
	aesKey string      // aes key

	accountInfo          *models.Account // 用户信息
	articleAmountPerPage = 10            // 每页显示的文章数
	recommendAmount      = 10            // 推荐文章数
)

func (c *BaseController) Prepare() {
	// 实例化redis缓存
	if rc == nil {
		if rcObj, err := getRedisCache(); err != nil {
			logs.Error("实例化redis缓存失败：%v", err)
			rc = cache.NewMemoryCache()
		} else {
			logs.Info("实例化redis缓存成功")
			rc = rcObj
		}
	}

	aesKey = beego.AppConfig.String("aes_key")

	if c.IsAjax() {
		c.EnableRender = false
	} else {
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())       // 全局xsrf
		c.Data["showTB"] = true                                    // 是否显示返回顶部按钮
		c.Data["appTitle"] = beego.AppConfig.String("app_title")   // 项目标题
		c.Data["githubUrl"] = beego.AppConfig.String("github_url") // github地址

		// 所有栏目
		categories, _ := models.GetAllCategories(nil)
		c.Data["categories"] = categories

		// 文章归档
		archive, _ := models.GetArticleArchive()
		c.Data["archive"] = archive

		// 所有标签
		tags, _ := models.GetAllTags(nil)
		c.Data["tags"] = tags

		// 推荐文章
		recommends, _ := models.GetRecommendArticles(recommendAmount)
		c.Data["recommends"] = recommends
	}

	// 记录账号信息以及登录状态
	l := c.GetSession("isLogin")
	a := c.GetSession("account")

	// 判断是否登录
	if l != nil && a != nil {
		accountInfo = a.(*models.Account)

		c.Data["isLogin"] = l.(bool)              // 是否登录标识
		c.Data["username"] = accountInfo.Username // 用户名
	} else {
		c.Data["isLogin"] = false // 是否登录标识
	}

	// 需要验证登录的路由
	needCheckUrl := []string{
		"POST" + beego.URLFor("CommentController.Post"), // 发表评论
	}

	method := c.Ctx.Request.Method
	path := c.Ctx.Request.URL.Path
	if utils.InSlice(method+path, needCheckUrl) {
		if l == nil {
			if c.IsAjax() {
				c.Data["json"] = &JSONResponse{Code: 402, Msg: "用户未登录"}
				c.ServeJSON()
			} else {
				c.Redirect(beego.URLFor("AuthController.Login"), 302)
			}
		}
	}
}

// getClientIp 获取客户端ip
func getClientIp(ctx *context.Context) string {
	return ctx.Input.IP()
}

// getRedisCache 获取redis缓存实例
func getRedisCache() (cache.Cache, error) {
	appName := beego.AppConfig.String("appname")
	address := beego.AppConfig.String("redis_host")

	return cache.NewCache("redis", `{"key":"`+appName+`","conn":"`+address+`"}`)
}

// isRequestValidOfIp 判断同一IP单位时间内的请求数量是否合法
func isClientRequestValid(ctx *context.Context, suffix string, limitDuration time.Duration, limitAmount int) bool {
	ip := getClientIp(ctx)
	key := strings.Replace(ip, ".", "_", -1) + suffix
	limit := rc.Get(key)
	if limit == nil {
		rc.Put(key, 1, limitDuration)
	} else {
		l, _ := strconv.Atoi(string(limit.([]uint8)))
		if l >= limitAmount {
			return false
		}

		rc.Incr(key)
	}

	return true
}
