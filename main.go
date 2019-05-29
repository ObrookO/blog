package main

import (
	_ "blog/routers"
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_password")+"@tcp(127.0.0.1:3306)/blog?charset=utf8")
	// 注册聊天系统的数据库
	// orm.RegisterDataBase("chat", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_password")+"@tcp(127.0.0.1:3306)/chat?charset=utf8")
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
