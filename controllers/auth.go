package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"time"

	"github.com/astaxie/beego/validation"

	"html/template"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

// 登录页面
func (a *AuthController) Login() {
	a.TplName = "auth/login.html"
	// 读取flash信息
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
	// 读取flash信息
	beego.ReadFromRequest(&a.Controller)

	a.Data["xsrfdata"] = template.HTML(a.XSRFFormHTML())
}

// 处理注册
func (a *AuthController) PostRegister() {
	username := a.GetString("username")
	password := a.GetString("password1")
	repassword := a.GetString("password2")

	user := &models.Users{Username: username, Password: password}

	// 	验证用户名密码
	valid := validation.Validation{}
	flash := beego.NewFlash()
	if ok, _ := valid.Valid(user); !ok {
		flash.Error("请输入用户名")
		flash.Store(&a.Controller)
		a.Redirect("/register", 302)
	} else {
		where := make(map[string]interface{})
		where["username"] = username

		// 验证用户名是否存在
		if exist := models.IsUserExists(where); exist {
			flash.Error("用户已存在")
			flash.Store(&a.Controller)
			a.Redirect("/register", 302)
		} else {
			if password != repassword {
				flash.Error("两次密码不一致")
				flash.Store(&a.Controller)
				a.Redirect("/register", 302)
			} else {
				user.Password = utils.Md5Str(password)
				user.CreatedAt = time.Now().Unix()
				if _, err := models.AddUser(user); err != nil {
					fmt.Println(err)
					flash.Error("用户注册失败")
					flash.Store(&a.Controller)
					a.Redirect("/register", 302)
				} else {
					a.Redirect("/login", 302)
				}
			}
		}
	}
}
