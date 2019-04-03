package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

type BaseController struct {
	beego.Controller
}

var (
	Categories  []models.Category       // 栏目列表
	Tags        []models.Tag            // 标签列表
	Archive     []models.ArticleArchive // 文章归档
	Like        []models.Article        // 猜你喜欢
	RedisClient *redis.Client
)

func (this *BaseController) Prepare() {
	// 获取所有的栏目
	Categories = models.GetCategories(make(map[string]interface{}))

	// 获取所有的标签
	Tags = models.GetTags(make(map[string]interface{}))

	// 获取猜你喜欢的文章
	Like = models.GetArticlesLimit(make(map[string]interface{}), 0, 5)

	// 文章归档
	Archive = models.Archive()

	// 连接redis数据库
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 读取所有文章的点赞记录
	records, _ := models.GetAllFavorRecord()
	for _, value := range records {
		key := "favor_" + value.Aid
		RedisClient.SAdd(key, value.Ip)
	}
}
