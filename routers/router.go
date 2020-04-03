package routers

import (
	"blog/controllers/admin"
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

	// 后台路由
	ns := beego.NewNamespace("/manong",
		// 首页
		beego.NSRouter("/", &admin.IndexController{}, "get:Get"),
		// 用户管理
		beego.NSNamespace("/users",
			beego.NSRouter("/", &admin.UserController{}),
		),
		// 栏目管理
		beego.NSNamespace("/categories",
			beego.NSRouter("/", &admin.CategoryController{}),
		),
		// 标签管理
		beego.NSNamespace("/tags",
			beego.NSRouter("/", &admin.TagController{}),
		),
		// 文章管理
		beego.NSNamespace("/articles",
			beego.NSRouter("/", &admin.ArticleController{}),
			beego.NSRouter("/draft", &admin.ArticleController{}, "get:Draft"),
		),
		// 点赞记录
		beego.NSNamespace("/favors",
			beego.NSRouter("/", &admin.FavorRecordController{}),
		),
		// 评论管理
		beego.NSNamespace("/comments",
			beego.NSRouter("/", &admin.CommentController{}),
			beego.NSRouter("/keyword", &admin.CommentController{}, "get:Keyword"),
		),
		// 日志管理
		beego.NSNamespace("/logs",
			beego.NSRouter("/home", &admin.LogController{}, "get:HomeLog"),
			beego.NSRouter("/admin", &admin.LogController{}, "get:AdminLog"),
		),
		// 日志管理
		beego.NSNamespace("/system",
			beego.NSRouter("/managers", &admin.ManagerController{}),
			beego.NSRouter("/roles", &admin.RoleController{}),
			beego.NSRouter("/permissions", &admin.PermissionController{}),
		),
	)

	beego.AddNamespace(ns)
}
