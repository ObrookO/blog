package models

import (
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id        int    `json:"id"`
	Aid       int    `json:"aid"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt string
}

func init() {
	orm.RegisterModel(new(Comment))
}

// 新增评论
func AddComment(value *Comment) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(value)
}

// 获取所有评论
func GetComments(where map[string]interface{}) ([]Comment, int64, error) {
	var comments []Comment
	o := orm.NewOrm()
	needle := o.QueryTable("comment")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	rows, err := needle.OrderBy("-id").Limit(20).All(&comments)

	return comments, rows, err
}

// 获取的评论数量
func GetCommentsCount(where map[string]interface{}) (int64, error) {
	o := orm.NewOrm()
	needle := o.QueryTable("comment")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	return needle.Count()
}
