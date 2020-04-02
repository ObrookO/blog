package home

import (
	"blog/models"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
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
	Tags        []models.Tag            // 标签列表
	Archive     []models.ArticleArchive // 文章归档
	Like        []models.Article        // 猜你喜欢
	UserInfo    map[string]interface{}  // 用户信息
	IsLogin     bool                    // 是否登录标识
	RedisClient *redis.Client
)

func (c *BaseController) Prepare() {
	// 用户信息
	if isLogin := c.GetSession("isLogin"); isLogin != nil {
		IsLogin = isLogin.(bool)
		if info := c.GetSession("user"); info != nil {
			UserInfo = info.(map[string]interface{})
		}
	}

	// 连接redis数据库
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 获取所有的标签
	Tags = models.GetTags(make(map[string]interface{}))

	// 文章归档
	Archive = models.Archive()

	// 获取猜你喜欢的文章
	Like = models.GetArticlesLimit(nil, 0, 5)

	// 读取所有文章的点赞记录
	records, _ := models.GetAllFavorRecord()
	for _, value := range records {
		key := "favor_" + value.Aid
		RedisClient.SAdd(key, value.Ip)
	}
}
