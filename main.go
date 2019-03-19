package main

import (
	"blog/controllers"
	"blog/models"
	_ "blog/routers"
	"strings"

	"github.com/go-redis/redis"

	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:12345678@tcp(127.0.0.1:3306)/blog?charset=utf8")
	// 注册聊天系统的数据库
	orm.RegisterDataBase("chat", "mysql", "root:12345678@tcp(127.0.0.1:3306)/chat?charset=utf8")

	// 获取所有的栏目
	controllers.Categories = models.GetCategories(make(map[string]interface{}))

	// 获取所有的标签
	controllers.Tags = models.GetTags(make(map[string]interface{}))

	// 获取猜你喜欢的文章
	controllers.Like = models.GetArticlesLimit(make(map[string]interface{}), 0, 5)

	// 文章归档
	controllers.Archive = models.Archive()

	// 连接redis数据库
	controllers.RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 读取所有文章的点赞记录
	records, _ := models.GetAllFavorRecord()
	for _, value := range records {
		key := "favor_" + value.Aid
		if exists := controllers.RedisClient.Exists(key).Val(); exists == 0 {
			controllers.RedisClient.SAdd("favor_"+value.Aid, value.Ip)
		}
	}
}

func main() {
	// 设置前端资源路径
	beego.SetStaticPath("/mycss", "static/css")
	beego.SetStaticPath("/myimg", "static/img")
	beego.SetStaticPath("/myjs", "static/js")
	beego.SetStaticPath("/myi", "static/i")
	beego.SetStaticPath("/myfonts", "static/fonts")

	// 注册自定义函数
	beego.AddFuncMap("formatDate", formatDate)

	// orm.Debug = true
	beego.Run()
}

// 将Y-m-d H:i:s的时间格式化为Y/m/d
func formatDate(t string) string {
	date := t[0:10]
	return strings.Replace(date, "-", "/", -1)
}
