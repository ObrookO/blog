package controllers

import (
	"blog/models"
	"blog/utils"
	"html/template"

	"github.com/astaxie/beego"
)

type AuthController struct {
	BaseController
}

// 登录页面
func (c *AuthController) Login() {
	c.Layout = "layouts/master.html"
	c.TplName = "auth/login.html"
	c.LayoutSections = map[string]string{
		"Script": "auth/login_script.html",
	}

	c.Data = map[interface{}]interface{}{
		"isLogin":  IsLogin,
		"user":     UserInfo,
		"xsrfdata": template.HTML(c.XSRFFormHTML()),
	}
}

// 处理登录
func (c *AuthController) DoLogin() {
	c.EnableRender = false

	username := c.GetString("username")
	password := c.GetString("password")

	where := map[string]interface{}{
		"username": username,
		"password": utils.Md5Str(password),
	}

	user, _ := models.GetOneUser(where)
	if user.Id == 0 {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "用户名或密码错误"}
		c.ServeJSON()
		return
	}

	c.SetSession("isLogin", true)
	c.SetSession("user", map[string]interface{}{
		"id":       user.Id,
		"username": username,
	})

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK"}
	c.ServeJSON()
}

// 注册页面
func (c *AuthController) Register() {
	c.Layout = "layouts/master.html"
	c.TplName = "auth/register.html"
	c.LayoutSections = map[string]string{
		"Script": "auth/reg_script.html",
	}

	c.Data = map[interface{}]interface{}{
		"isLogin":  IsLogin,
		"user":     UserInfo,
		"xsrfdata": template.HTML(c.XSRFFormHTML()),
	}
}

// 处理注册
func (c *AuthController) DoRegister() {
	c.EnableRender = false

	username := c.GetString("username")
	password1 := c.GetString("password1")
	password2 := c.GetString("password2")

	// 验证用户名是否唯一
	if models.IsUserExists(map[string]interface{}{"username": username}) {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "用户名已存在"}
		c.ServeJSON()
		return
	}

	// 验证密码长度
	if len(password1) < 8 || len(password1) > 16 {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "密码长度必须在8-16位之间"}
		c.ServeJSON()
		return
	}

	// 验证两次密码是否一致
	if password1 != password2 {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "两次密码不一致"}
		c.ServeJSON()
		return
	}

	// 添加用户
	id, err := models.AddUser(&models.User{Username: username, Password: utils.Md5Str(password1)})
	if err != nil {
		//logs.Error("register error: %v", err)
		c.Data["json"] = &JSONResponse{Code: 400003, Msg: "注册失败"}
		c.ServeJSON()
		return
	}

	// 注册成功后直接登录
	c.SetSession("isLogin", true)
	c.SetSession("user", map[string]interface{}{
		"id":       id,
		"username": username,
	})

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "注册成功"}
	c.ServeJSON()
}

func (c *AuthController) Logout() {
	c.EnableRender = false
	c.DelSession("isLogin")
	c.DelSession("user")
	c.Redirect(beego.URLFor("IndexController.Get"), 302)
}
