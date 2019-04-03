package controllers

import (
	"blog/models"
	"strconv"
	"time"
)

type CommentController struct {
	BaseController
}

// 新增评论
func (c *CommentController) Comment() {
	if isLogin := c.GetSession("isLogin"); isLogin == nil {
		c.Data["json"] = &JsonResult{Code: 401, Message: "您还没有登录"}
	} else {
		var comment models.Comment
		aid, _ := strconv.Atoi(c.GetString("aid"))
		comment.Aid = aid
		comment.Username = c.GetSession("username").(string)
		comment.Content = c.GetString("content")
		comment.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

		if _, err := models.AddComment(&comment); err != nil {
			c.Data["json"] = &JsonResult{Code: 402, Message: "评论失败"}
		} else {
			c.Data["json"] = &JsonResult{Code: 200, Message: "OK"}
		}
	}

	c.ServeJSON()
}
