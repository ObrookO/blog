package controllers

import (
	"blog/models"
	"blog/tool"
	"strconv"
	"time"

	"github.com/ObrookO/go-utils"
	"github.com/mojocn/base64Captcha"
)

type AuthController struct {
	BaseController
}

var (
	// 图片验证码相关
	captchaLimitAmount   = 10                        // 同一IP单位时间内获取图片验证码的次数
	captchaLimitDuration = time.Minute               // 单位时间
	captchaLimitSuffix   = "_request_captcha_amount" // 保存图片验证码次数的key的后缀

	// 注册账号邮箱验证码相关
	registerCodeSuffix   = "_register_account_code" // 注册账号邮箱验证码后缀
	registerCodeDuration = 2 * time.Minute          // 注册账号邮箱验证码有效期

	// 重置密码邮箱验证码相关
	resetCodeSuffix   = "_reset_password_code" // 重置密码邮箱验证码后缀
	resetCodeDuration = 2 * time.Minute        // 重置密码邮箱验证码有效期

	// 发送邮件限制
	sendEmailLimitAmount    = 10                              // 同一IP单位时间内获取邮箱验证码的次数
	sendEmailLimitDuration  = 2 * time.Hour                   // 单位时间
	registerCodeLimitSuffix = "_register_account_code_amount" // 保存注册账号邮箱验证码次数的key的后缀
	resetCodeLimitSuffix    = "_reset_password_code_amount"   // 保存重置密码邮箱验证码次数的key的后缀

	// 用户名密码验证相关
	accountErrorLimitAmount   = 5                    // 同一IP单位时间内允许登录失败的次数
	accountErrorLimitDuration = 30 * time.Minute     // 单位时间
	accountErrorLimitSuffix   = "_account_error_num" // 保存登录失败次数的key的后缀
)

// GetCaptcha 获取验证码
func (c *AuthController) GetCaptcha() {
	/** 判断同一IP单位时间内获取验证码的次数是否合法 **/
	if !isClientRequestValid(c.Ctx, captchaLimitSuffix, captchaLimitDuration, captchaLimitAmount) {
		c.Data["json"] = &JSONResponse{Code: 500000, Msg: "操作过于频繁，请稍后重试"}
		c.ServeJSON()
		return
	}

	// 生成验证码
	id, bs64, err := utils.GetCaptcha()
	if err != nil {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "获取验证码失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK", Data: map[string]string{"id": id, "captcha": bs64}}
	c.ServeJSON()
}

// Register 注册页面
func (c *AuthController) Register() {
	c.TplName = "auth/reg.html"
}

// SendRegisterEmail 发送注册邮件
func (c *AuthController) SendRegisterEmail() {
	/** 判断同一IP单位时间内获取邮箱验证码的次数是否合法 **/
	if !isClientRequestValid(c.Ctx, registerCodeLimitSuffix, sendEmailLimitDuration, sendEmailLimitAmount) {
		c.Data["json"] = &JSONResponse{Code: 500000, Msg: "操作过于频繁，请稍后再试"}
		c.ServeJSON()
		return
	}

	// 验证邮箱格式
	email := c.GetString("email")
	if !tool.IsValidEmail(email) {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "邮箱格式不正确"}
		c.ServeJSON()
		return
	}

	// 判断邮箱是否被占用
	if models.IsAccountExist(map[string]interface{}{"email": email}) {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "邮箱已被占用"}
		c.ServeJSON()
		return
	}

	// 保存验证码
	code := utils.RandomStr(8)
	rc.Put(email+registerCodeSuffix, code, registerCodeDuration)

	go tool.SendRegisterAccountEmail(code, email)

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK"}
	c.ServeJSON()
}

// SendResetEmail 发送重置密码邮件
func (c *AuthController) SendResetEmail() {
	/** 判断同一IP单位时间内获取邮箱验证码的次数是否合法 **/
	if !isClientRequestValid(c.Ctx, resetCodeSuffix, sendEmailLimitDuration, sendEmailLimitAmount) {
		c.Data["json"] = &JSONResponse{Code: 500000, Msg: "操作过于频繁，请稍后再试"}
		c.ServeJSON()
		return
	}

	// 验证用户是否存在
	username := c.GetString("username")
	account, _ := models.GetOneAccount(map[string]interface{}{"username": username})
	if account.Id == 0 {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "用户不存在"}
		c.ServeJSON()
		return
	}

	// 保存验证码
	code := utils.RandomStr(8)
	rc.Put(account.Email+resetCodeSuffix, code, resetCodeDuration)

	go tool.SendResetPasswordEmail(code, account.Email)

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK"}
	c.ServeJSON()
}

// DoRegister 处理注册
func (c *AuthController) DoRegister() {
	username := c.GetString("username")
	email := c.GetString("email")
	code := c.GetString("code")
	password := c.GetString("password")

	usernameLength := tool.GetCharAmount(username)
	if usernameLength == 0 || usernameLength > 50 {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "用户名的长度为0-50"}
		c.ServeJSON()
		return
	}

	if !tool.IsValidEmail(email) {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "邮箱格式不正确"}
		c.ServeJSON()
		return
	}

	key := email + registerCodeSuffix
	if rc.Get(key) == nil {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "验证码已过期"}
		c.ServeJSON()
		return
	}

	if code != string(rc.Get(key).([]uint8)) {
		c.Data["json"] = &JSONResponse{Code: 400003, Msg: "邮箱验证码错误"}
		c.ServeJSON()
		return
	}

	if !tool.IsValidPassword(password) {
		c.Data["json"] = &JSONResponse{Code: 400004, Msg: "密码由8-16位字母数字下划线组成"}
		c.ServeJSON()
		return
	}

	// 判断用户名是否重复
	if models.IsAccountExist(map[string]interface{}{"username": username}) {
		c.Data["json"] = &JSONResponse{Code: 400005, Msg: "用户名已被占用"}
		c.ServeJSON()
		return
	}

	// 判断邮箱是否重复
	if models.IsAccountExist(map[string]interface{}{"email": email}) {
		c.Data["json"] = &JSONResponse{Code: 400006, Msg: "邮箱已被占用"}
		c.ServeJSON()
		return
	}

	encryptedPassword, _ := utils.AesEncrypt(password, aesKey)
	if _, err := models.AddAccount(models.Account{
		Username:     username,
		Email:        email,
		Password:     encryptedPassword,
		AllowComment: 1,
		Status:       1,
	}); err != nil {
		c.Data["json"] = &JSONResponse{Code: 400007, Msg: "注册失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK"}
	c.ServeJSON()
}

// Login 登录页面
func (c *AuthController) Login() {
	c.TplName = "auth/login.html"
}

// DoLogin 处理登录
func (c *AuthController) DoLogin() {
	captchaId := c.GetString("captcha_id")
	captcha := c.GetString("captcha")
	username := c.GetString("username")
	password := c.GetString("password")

	if tool.GetCharAmount(username) == 0 {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "请输入用户名"}
		c.ServeJSON()
		return
	}

	if tool.GetCharAmount(captcha) == 0 {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "请输入验证码"}
		c.ServeJSON()
		return
	}

	// 校验验证码
	ca := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)
	if !ca.Verify(captchaId, captcha, true) {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "验证码错误"}
		c.ServeJSON()
		return
	}

	// 判断当前账号密码错误次数是否超出限制
	key := username + accountErrorLimitSuffix
	limit := rc.Get(key)
	if limit != nil {
		l, _ := strconv.Atoi(string(limit.([]uint8)))
		if l >= accountErrorLimitAmount {
			c.Data["json"] = &JSONResponse{Code: 400003, Msg: "由于密码错误过多，此账号已被锁定，请30分钟后重试"}
			c.ServeJSON()
			return
		}
	}

	// 校验用户名密码
	encryptPass, _ := utils.AesEncrypt(password, aesKey)
	filter := map[string]interface{}{
		"username": username,
		"password": encryptPass,
	}

	account, _ := models.GetOneAccount(filter, "id", "username", "email", "avatar", "allow_comment", "status")
	if account.Id == 0 {
		if limit == nil {
			rc.Put(key, 1, accountErrorLimitDuration)
		} else {
			rc.Incr(key)
		}

		c.Data["json"] = &JSONResponse{Code: 400004, Msg: "用户名或密码错误"}
		c.ServeJSON()
		return
	}

	if account.Status == 0 {
		c.Data["json"] = &JSONResponse{Code: 400005, Msg: "该账号已被禁用，请联系管理员"}
		c.ServeJSON()
		return
	}

	c.SetSession("isLogin", true)
	c.SetSession("account", &account)

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK"}
	c.ServeJSON()
}

// Logout 退出
func (c *AuthController) Logout() {
	c.DelSession("isLogin")
	c.DelSession("account")

	c.Redirect(c.URLFor("IndexController.Get"), 302)
}

// ForgetPassword 忘记密码页面
func (c *AuthController) ForgetPassword() {
	c.TplName = "auth/forget.html"
}

// ResetPassword 重置密码
func (c *AuthController) ResetPassword() {
	username := c.GetString("username")
	code := c.GetString("code")
	password := c.GetString("password")

	account, _ := models.GetOneAccount(map[string]interface{}{"username": username})

	key := account.Email + resetCodeSuffix
	if rc.Get(key) == nil {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "验证码已过期，请重新获取"}
		c.ServeJSON()
		return
	}

	if code != string(rc.Get(key).([]uint8)) {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "邮箱验证码错误"}
		c.ServeJSON()
		return
	}

	if !tool.IsValidPassword(password) {
		c.Data["json"] = &JSONResponse{Code: 400003, Msg: "密码由8-16位字母数字下划线组成"}
		c.ServeJSON()
		return
	}

	encryptedPassword, _ := utils.AesEncrypt(password, aesKey)
	if _, err := models.UpdateAccount(map[string]interface{}{
		"username": username,
	}, map[string]interface{}{
		"password":   encryptedPassword,
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		c.Data["json"] = &JSONResponse{Code: 400004, Msg: "重置密码失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK"}
	c.ServeJSON()

}
