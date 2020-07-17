package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Get")

	// 获取验证码
	beego.Router("/captcha", &controllers.AuthController{}, "get:GetCaptcha")

	// 登录相关路由
	an := beego.NewNamespace("/auth",
		beego.NSRouter("/login", &controllers.AuthController{}, "get:Login;post:DoLogin"),
		beego.NSRouter("/reg", &controllers.AuthController{}, "get:Register;post:DoRegister"),
		beego.NSRouter("/reg/mail", &controllers.AuthController{}, "post:SendRegisterEmail"),
		beego.NSRouter("/logout", &controllers.AuthController{}, "get:Logout"),
		beego.NSRouter("/forget", &controllers.AuthController{}, "get:ForgetPassword"),
		beego.NSRouter("/forget/mail", &controllers.AuthController{}, "post:SendResetEmail"),
		beego.NSRouter("/reset", &controllers.AuthController{}, "post:ResetPassword"),
	)

	// 文章相关路由
	arn := beego.NewNamespace("/articles",
		beego.NSRouter("/", &controllers.ArticleController{}, "get:Get"),
		beego.NSRouter("/:id", &controllers.ArticleController{}, "get:GetArticleDetail"),
		beego.NSRouter("/comments", &controllers.CommentController{}, "post:Post"),
		beego.NSRouter("/favor", &controllers.ArticleController{}, "post:Favor"),
	)

	// 实用工具相关路由
	tn := beego.NewNamespace("/tools",
		beego.NSRouter("/md5", &controllers.ToolController{}, "get:Md5;post:CalculateMd5"),
		beego.NSRouter("/aes", &controllers.ToolController{}, "get:Aes;post:CalculateAes"),
		beego.NSRouter("/base64", &controllers.ToolController{}, "get:Base64"),
		beego.NSRouter("/random", &controllers.ToolController{}, "get:RandomStr"),
		beego.NSRouter("/length", &controllers.ToolController{}, "get:StrLength"),
	)

	// 干货收藏
	ren := beego.NewNamespace("/resource",
		beego.NSRouter("/", &controllers.ResourceController{}, "get:Get"),
	)

	beego.AddNamespace(an, arn, tn, ren)
}
