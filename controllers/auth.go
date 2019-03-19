package controllers

import (
	"blog/models"
	"blog/utils"

	"html/template"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

// 登录页面
func (a *AuthController) Login() {
	a.TplName = "auth/login.html"
	// 读取flash中的信息
	beego.ReadFromRequest(&a.Controller)

	a.Data["xsrfdata"] = template.HTML(a.XSRFFormHTML())
}

// 处理登录
func (a *AuthController) PostLogin() {
	username := a.GetString("username")
	password := a.GetString("password")

	where := make(map[string]interface{})
	where["username"] = username
	where["password"] = utils.Md5Str(password)

	if exist := models.IsUserExists(where); !exist {
		flash := beego.NewFlash()
		flash.Error("用户名或密码错误")
		flash.Store(&a.Controller)
		a.Redirect("/login", 302)
	} else {
		a.SetSession("username", username)
		a.SetSession("isLogin", true)
		a.Redirect("/", 302)
	}
}

// 注册页面
func (a *AuthController) Register() {
	a.TplName = "auth/register.html"
	a.Data["xsrfdata"] = template.HTML(a.XSRFFormHTML())
}

// 处理注册
func (a *AuthController) PostRegister() {

}
