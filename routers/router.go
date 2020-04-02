package routers

import (
	"blog/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &home.IndexController{})
	// 点赞功能
	beego.Router("/favor/:aid([0-9]+)", &home.IndexController{}, "get:Favor")
	// g根据栏目Id获取文章列表
	beego.Router("/cate/:cid([0-9]+)", &home.ArticleController{}, "get:GetAllArticlesByCate")
	// 根据日期获取文章列表
	beego.Router("/date/:year([0-9]{4})/:month([0-9]{2})", &home.ArticleController{}, "get:GetArticlesByDate")
	// 根据标签获取文章列表
	beego.Router("/tag/:tid([0-9]+)", &home.ArticleController{}, "get:GetArticlesByTag")
	// 文章详情
	beego.Router("/article/:aid([0-9]+)", &home.ArticleController{}, "get:ArticleInfo")
	// 评论功能
	beego.Router("/comment", &home.CommentController{}, "post:Comment")

	// 注册、登录、退出
	beego.Router("/login", &home.AuthController{}, "get:Login;post:DoLogin")
	beego.Router("/register", &home.AuthController{}, "get:Register;post:DoRegister")
	beego.Router("/logout", &home.AuthController{}, "get:Logout")
}
