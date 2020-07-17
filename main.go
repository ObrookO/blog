package main

import (
	"blog/controllers"
	"blog/models"
	_ "blog/routers"
	"blog/tool"
	"encoding/gob"

	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	gob.Register(&models.Manager{})
	gob.Register(&models.Account{})
}

func main() {
	// 设置前端资源路径
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/i", "static/i")
	beego.SetStaticPath("/fonts", "static/fonts")

	// 注册自定义函数
	beego.AddFuncMap("add", add)
	beego.AddFuncMap("sub", sub)
	beego.AddFuncMap("getTitle", getTitle)

	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}

// 增加1
func add(i, step int) int {
	return i + step
}

// 减少1
func sub(i, step int) int {
	return i - step
}

// 获取文章标题
func getTitle(title string, limit int) string {
	l := tool.GetCharAmount(title)
	if l > limit {
		return beego.Substr(title, 0, limit) + " ..."
	}

	return title
}
