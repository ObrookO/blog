package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	// 点赞功能
	beego.Router("/favor/:aid([0-9]+)", &controllers.IndexController{}, "get:Favor")
	// 根据栏目Id获取文章列表
	beego.Router("/cate/:cid([0-9]+)", &controllers.ArticleController{}, "get:GetAllArticlesByCate")
	// 根据日期获取文章列表
	beego.Router("/date/:year([0-9]{4})/:month([0-9]{2})", &controllers.ArticleController{}, "get:GetArticlesByDate")
	// 根据标签获取文章列表
	beego.Router("/tag/:tid([0-9]+)", &controllers.ArticleController{}, "get:GetArticlesByTag")
	// 文章详情
	beego.Router("/article/:aid([0-9]+)", &controllers.ArticleController{}, "get:ArticleInfo")
	// 评论功能
	beego.Router("/comment", &controllers.CommentController{}, "post:Comment")

	// 注册、登录、退出
	beego.Router("/login", &controllers.AuthController{}, "get:Login;post:DoLogin")
	beego.Router("/register", &controllers.AuthController{}, "get:Register;post:DoRegister")
	beego.Router("/logout", &controllers.AuthController{}, "get:Logout")
}
