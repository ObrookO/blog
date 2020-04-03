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
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String(
		"db_password")+"@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai")
	// 注册聊天系统的数据库
	// orm.RegisterDataBase("chat", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_password")+"@tcp(127.0.0.1:3306)/chat?charset=utf8")
}

func main() {
	// 设置前端资源路径
	beego.SetStaticPath("/hcss", "static/home/css")
	beego.SetStaticPath("/himg", "static/home/img")
	beego.SetStaticPath("/hjs", "static/home/js")
	beego.SetStaticPath("/hi", "static/home/i")
	beego.SetStaticPath("/hfonts", "static/home/fonts")

	beego.SetStaticPath("/acss", "static/admin/css")
	beego.SetStaticPath("/aimg", "static/admin/img")
	beego.SetStaticPath("/ajs", "static/admin/js")
	beego.SetStaticPath("/ai", "static/admin/i")
	beego.SetStaticPath("/afonts", "static/admin/fonts")

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
