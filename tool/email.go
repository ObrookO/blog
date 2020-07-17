package tool

import (
	"blog/models"
	"strings"

	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
)

var appTitle = beego.AppConfig.String("app_title")                                         // 博客名称
var appURL = beego.AppConfig.String("app_url")                                             // 博客地址
var emailFooter = "博客地址：<a href=\"" + appURL + "\" target=\"_blank\">" + appTitle + "</a>" // 邮件页脚

// SendRegisterAccountEmail 发送注册账号验证码
func SendRegisterAccountEmail(code, toAddress string) {
	subject := "注册账号"
	contentType := "text/html"
	content := []string{
		"您正在使用此邮箱注册账号，若不是本人操作，请忽略此邮件。",
		"验证码：" + code,
		strings.Repeat("<br/>", 2) + emailFooter,
	}

	sendEmail(models.RegisterAccount, toAddress, subject, contentType, strings.Join(content, "<br>"))
}

// SendResetPasswordEmail 发送重置密码验证码
func SendResetPasswordEmail(code, toAddress string) {
	subject := "重置密码"
	contentType := "text/html"
	content := []string{
		"您正在进行重置密码操作，若不是本人操作，请联系<a href=\"mailto:" + beego.AppConfig.String("manager_email") + "\">管理员</a>。",
		"验证码：" + code,
		strings.Repeat("<br/>", 2) + emailFooter,
	}

	sendEmail(models.ResetAccountPassword, toAddress, subject, contentType, strings.Join(content, "<br>"))
}

// sendEmail 发送邮件
func sendEmail(emailType int, toAddress, subject, contentType, msg string) {
	m := gomail.NewMessage()

	host := beego.AppConfig.String("email_host")
	port, _ := beego.AppConfig.Int("email_port")
	fromAddress := beego.AppConfig.String("email_from_address")
	fromName := beego.AppConfig.String("email_from_name")
	password := beego.AppConfig.String("email_password")

	m.SetHeader("From", fromAddress)
	// 设置发送人别名
	m.SetAddressHeader("From", fromAddress, fromName)
	m.SetHeader("To", toAddress)
	m.SetHeader("Subject", subject)
	m.SetBody(contentType, msg)

	var reason string
	var result = "SUCCESS"

	d := gomail.NewDialer(host, port, fromAddress, password)
	if err := d.DialAndSend(m); err != nil {
		result = "FAIL"
		reason = err.Error()
	}

	models.AddEmailLog(models.EmailLog{
		EmailType: emailType,
		Address:   toAddress,
		Content:   msg,
		Result:    result,
		Reason:    reason,
	})
}
