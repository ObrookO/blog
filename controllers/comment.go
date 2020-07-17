package controllers

import (
	"blog/models"
	"time"
)

type CommentController struct {
	BaseController
}

var (
	commentLimitSuffix   = "_comment_amount" // 存放评论数量的key的后缀
	commentLimitAmount   = 10                // 同一IP单位时间内评论数量的限制
	commentLimitDuration = time.Minute       // 单位时间
)

// Post 发表评论
func (c *CommentController) Post() {
	id, _ := c.GetInt("id")
	content := c.GetString("content")

	if !isClientRequestValid(c.Ctx, commentLimitSuffix, commentLimitDuration, commentLimitAmount) {
		c.Data["json"] = &JSONResponse{Code: 500000, Msg: "操作过于频繁，请稍后再试"}
		c.ServeJSON()
		return
	}

	// 判断文章是否存在
	article, _ := models.GetOneArticle(map[string]interface{}{"id": id, "status": 1})
	if article.Id == 0 {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "文章不存在"}
		c.ServeJSON()
		return
	}

	// 判断文章是否允许评论
	if article.AllowComment != 1 {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "此文章不允许评论"}
		c.ServeJSON()
		return
	}

	// 判断用户是否允许评论
	if accountInfo.AllowComment != 1 {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "此账号评论权限已被禁用，请联系管理员"}
		c.ServeJSON()
		return
	}

	comment := models.Comment{
		OriginalContent: content,
		Ip:              getClientIp(c.Ctx),
	}
	comment.Account = accountInfo
	comment.Article = &article

	if _, err := models.AddComment(comment); err != nil {
		c.Data["json"] = &JSONResponse{Code: 400003, Msg: "评论失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK"}
	c.ServeJSON()
	return
}
